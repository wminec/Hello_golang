package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/gomail.v2"
)

// 템플릿 파일 경로
const templatePath = "templates/form.html"

// 템플릿 로드
var tmpl = template.Must(template.ParseFiles(filepath.Join(templatePath)))

// formHandler 함수는 GET 요청 시 폼을 표시하고, POST 요청 시 이메일을 보냅니다.
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	message := r.FormValue("email")
	sendEmail(message)

	fmt.Fprintf(w, "Email sent successfully!")
}

// sendEmail 함수는 주어진 메시지를 사용하여 이메일을 보냅니다.
func sendEmail(message string) {
	// 환경 변수에서 이메일 설정을 읽습니다.
	subject := os.Getenv("EMAIL_SUBJECT")
	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

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

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
