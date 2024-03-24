package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"restful-api/config"
)

func SendEmail(cfg *config.Config, emails []string, subject string, body string) error {

	msg := []byte(
		fmt.Sprintf(
			"From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=utf-8;\r\n\r\n%s",
			cfg.Smtp.User,
			emails,
			subject,
			body,
		),
	)

	auth := smtp.PlainAuth("", cfg.Smtp.User, cfg.Smtp.Pass, cfg.Smtp.Host)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", cfg.Smtp.Host, cfg.Smtp.Port),
		auth,
		cfg.Smtp.User,
		emails,
		msg,
	)

	if err != nil {
		return err
	}

	fmt.Println("[+] Email send successfully")

	return nil
}

func renderTemplate(templatePath string, data interface{}) (string, error) {
	templateBytes, err := os.ReadFile(templatePath)

	if err != nil {
		return "", err
	}

	templateContent := string(templateBytes)

	tmpl, err := template.New("email").Parse(templateContent)

	if err != nil {
		return "", err
	}

	var html bytes.Buffer
	err = tmpl.Execute(&html, data)

	if err != nil {
		return "", err
	}

	return html.String(), nil
}

func SendTemplatedEmail(
	cfg *config.Config,
	emails []string,
	subject string,
	body string,
	templatePath string,
	data interface{},
) error {

	html, err := renderTemplate(cfg.App.Root+"/templates/email/"+templatePath, data)

	if err != nil {
		return err
	}

	err = SendEmail(cfg, emails, subject, html)

	if err != nil {
		return err
	}

	return nil
}
