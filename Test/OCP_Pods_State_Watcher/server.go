package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"syscall"
	"time"

	"gopkg.in/gomail.v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var startTime time.Time

func sendEmail(message string) {
	// 환경 변수에서 이메일 설정을 읽습니다.
	subject := os.Getenv("EMAIL_SUBJECT")
	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	tlsSkipVerify := os.Getenv("TLS_SKIP_VERIFY")

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatalf("Invalid SMTP port: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(smtpServer, port, smtpUser, smtpPassword)
	if tlsSkipVerify == "true" {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return
	}
}

// handlePodEvent는 Pod 이벤트를 처리하고 상태를 출력합니다.
func handlePodEvent(obj interface{}) {
	email := os.Getenv("EMAIL")
	pod, ok := obj.(*v1.Pod)
	if !ok {
		log.Println("Error casting to Pod")
		return
	}
	if pod.CreationTimestamp.Time.After(startTime) && (pod.Status.Phase == v1.PodFailed || pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodRunning) {
		message := fmt.Sprintf("%s : Pod %s in namespace %s has entered %s state", pod.Status.StartTime, pod.Name, pod.Namespace, pod.Status.Phase)
		fmt.Printf("%s\n", message)
		if email == "true" {
			sendEmail(message)
		}
	}
}

func main() {
	// In-cluster configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Error creating in-cluster config: %v\n", err)

		// Local configuration for testing
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			log.Fatalf("Error building kubeconfig: %v\n", err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating clientset: %v\n", err)
	}

	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		log.Fatalf("NAMESPACE environment variable not set")
	}

	debug := os.Getenv("DEBUG")

	labelKeys := os.Getenv("LABEL_KEYS")

	watchlists := make([]cache.ListerWatcher, 0)

	if labelKeys == "" {
		watchlists = append(watchlists, cache.NewListWatchFromClient(
			clientset.CoreV1().RESTClient(),
			"pods",
			namespace,
			fields.Everything(),
		))
	} else {
		keys := strings.Split(labelKeys, ",")

		for _, key := range keys {
			req, err := labels.NewRequirement(key, selection.Exists, nil)
			if err != nil {
				log.Fatalf("Failed to create requirement: %v", err)
			}
			selector := labels.NewSelector().Add(*req)

			fmt.Printf("Selector: %s\n", selector.String())

			watchlists = append(watchlists, cache.NewFilteredListWatchFromClient(
				clientset.CoreV1().RESTClient(),
				"pods",
				namespace,
				func(options *metav1.ListOptions) {
					options.LabelSelector = selector.String()
				},
			))
		}
	}

	eventHandler := cache.ResourceEventHandlerFuncs{
		//AddFunc: func(obj interface{}) {
		//	fmt.Printf("Add: %v\n", obj)
		//	handlePodEvent(obj)
		//},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldPod, ok1 := oldObj.(*v1.Pod)
			newPod, ok2 := newObj.(*v1.Pod)
			if debug == "true" {
				fmt.Printf("Update: %v %v\n", oldPod.Status.Phase, newPod.Status.Phase)
			}
			if ok1 && ok2 && !reflect.DeepEqual(oldPod.Status.Phase, newPod.Status.Phase) {
				handlePodEvent(newObj)
			}
		},
	}

	stop := make(chan struct{})
	defer close(stop)

	for _, watchlist := range watchlists {
		// Create and start informer for each watchlist
		_, controller := cache.NewInformer(
			watchlist,
			&v1.Pod{},
			0,
			eventHandler,
		)

		go func() {
			startTime = time.Now()
			fmt.Printf("Start time: %s\n", startTime)
			controller.Run(stop)
		}()
	}

	// Handle graceful shutdown
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	<-sigterm
	fmt.Println("Received a termination signal, stopping all informers...")
	close(stop) // This will stop all informers
}
