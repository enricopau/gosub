package mail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(mail string, generateMsg func() []byte) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	mailServerAddr := os.Getenv("MAILSERVER_ADDR")
	senderAddr := os.Getenv("SENDER_ADDR")
	auth := smtp.CRAMMD5Auth(os.Getenv("MAILBOX_USER"), os.Getenv("MAILBOX_PW"))
	return smtp.SendMail(mailServerAddr, auth, senderAddr, []string{mail}, generateMsg())
}

func generateMsg(recipient, subject, verificationLink string) []byte {
	msg := []byte(fmt.Sprintf("To: %s\r\n", recipient) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		fmt.Sprintf("Please confirm your email address by clicking the link below.\r\n") +
		fmt.Sprintf("%s\r\n", verificationLink))
	return msg
}
