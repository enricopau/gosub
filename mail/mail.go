package mail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

// SendMail sends a mail to the recipients address.
func SendMail(recipientAddr string, generateMsg func() []byte) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	mailServerAddr := os.Getenv("MAILSERVER_ADDR")
	senderAddr := os.Getenv("SENDER_ADDR")
	auth := smtp.CRAMMD5Auth(os.Getenv("MAILBOX_USER"), os.Getenv("MAILBOX_PW"))
	return smtp.SendMail(mailServerAddr, auth, senderAddr, []string{recipientAddr}, generateMsg())
}

func generateMsg(recipient, subject, verificationLink string) []byte {
	msg := []byte(fmt.Sprintf("To: %s\r\n", recipient) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		fmt.Sprintf("Please confirm your email address by clicking the link below.\r\n") +
		fmt.Sprintf("%s\r\n", verificationLink))
	return msg
}

func generateVerificationLink(token string) string {
	return ""
}
