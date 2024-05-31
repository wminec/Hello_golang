package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var startTime = time.Now()

// handleEvent는 Kubernetes 이벤트를 처리하고 출력합니다.
func handleEvent(obj interface{}) {
	event, ok := obj.(*v1.Event)
	if !ok {
		log.Println("Error casting to event")
		return
	}
	if event.CreationTimestamp.Time.After(startTime) {
		fmt.Printf("Event: %s %s %s\n", event.InvolvedObject.Kind, event.InvolvedObject.Name, event.Message)
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
		namespace = "default"
	}

	// Event watcher
	watchlist := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		"events",
		namespace,
		fields.Everything(),
	)

	eventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc:    handleEvent,
		UpdateFunc: func(oldObj, newObj interface{}) { handleEvent(newObj) },
		DeleteFunc: handleEvent,
	}

	_, controller := cache.NewInformer(
		watchlist,
		&v1.Event{},
		0,
		eventHandler,
	)

	stop := make(chan struct{})
	defer close(stop)

	go controller.Run(stop)

	// Handle graceful shutdown
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
}
