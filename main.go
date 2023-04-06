package main

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
)

type Email struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func main() {
	// Open and read the JSON file.
	file, err := os.Open("emails.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse the JSON data into an array of email objects.
	var emails []Email
	err = json.NewDecoder(file).Decode(&emails)
	if err != nil {
		panic(err)
	}

	// Set the SMTP server and authentication credentials.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "your_username"
	smtpPassword := "your_password"

	// Loop through the emails and send each one.
	for _, email := range emails {
		// Compose the email message.
		message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", email.From, email.To[0], email.Subject, email.Body)
		if len(email.To) > 1 {
			message = fmt.Sprintf("From: %s\r\nTo: %s\r\nCc: %s\r\nSubject: %s\r\n\r\n%s", email.From, email.To[0], email.To[1:], email.Subject, email.Body)
		}

		// Connect to the SMTP server.
		auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email.From, email.To, []byte(message))
		if err != nil {
			fmt.Printf("Failed to send email: %v\n", err)
			continue
		}

		// Print a confirmation message.
		fmt.Printf("Email sent successfully to %s!\n", email.To[0])
	}
}


