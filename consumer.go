package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test email\r\n\r\n%s\r\n", recipient.Email, "Best country India!")

		// msg := []byte(formattedMsg)

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker %d: error parsing template for %s \n", id, recipient.Email)
			continue
		}

		fmt.Printf("Worker %d: sending email to %s \n", id, recipient.Email)

		error := smtp.SendMail(smtpHost+":"+smtpPort, nil, "vxnsh1@ongithub.com", []string{recipient.Email}, []byte(msg))

		if error  != nil {
			log.Fatal(error)
		}

		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: sent email to %s \n", id, recipient.Email)
	}
}
