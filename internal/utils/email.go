package utils

import (
    "crypto/tls"
    "fmt"
    "log"
    "os"

    "github.com/go-mail/mail/v2"
)

var (
    smtpHost = os.Getenv("SMTP_HOST")
    smtpPort = 587
    smtpUser = os.Getenv("SMTP_USER")
    smtpPass = os.Getenv("SMTP_PASS")
)

func createMessage(to string, subject string, body string) *mail.Message {
    m := mail.NewMessage()
    m.SetHeader("From", smtpUser)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)
    return m
}

func createDialer() *mail.Dialer {
    d := mail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
    d.TLSConfig = &tls.Config{
        ServerName:         smtpHost,
        InsecureSkipVerify: false,
    }
    return d
}

func SendPaymentEmail(userEmail string, amount float64) error {
    content := fmt.Sprintf(`
        <h1>Платёж проведён успешно</h1>
        <p>Сумма: <strong>%.2f RUB</strong></p>
        <small>Это автоматическое уведомление</small>
    `, amount)

    m := createMessage(userEmail, "Платёж подтверждён", content)
    d := createDialer()

    if err := d.DialAndSend(m); err != nil {
        log.Printf("SMTP error: %v", err)
        return fmt.Errorf("email sending failed")
    }

    log.Printf("Email sent to %s", userEmail)
    return nil
}
