package emailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendEmail(to, subject, body string) error {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	var msg strings.Builder

	msg.WriteString(fmt.Sprintf("From: %s\r\n", from))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", to))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	msg.WriteString(fmt.Sprintf("%s\r\n", body))
	msg.WriteString("\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", smtpHost, smtpPort), tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to the SMTP server: %v", err)
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %v", err)
	}
	defer client.Quit()

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to send email data: %v", err)
	}
	defer w.Close()

	_, err = w.Write([]byte(msg.String()))
	if err != nil {
		return fmt.Errorf("failed to write email data: %v", err)
	}

	return nil
}
