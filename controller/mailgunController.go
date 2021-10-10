package controller

import (
	"bytes"
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	mailgun "github.com/mailgun/mailgun-go/v4"
	"github.com/tidwall/gjson"
)

func NotifyMailgun(c *gin.Context) {

	jsonData, err := c.GetRawData()
	input := bytes.NewBuffer(jsonData).String()
	subject := gjson.Get(input, "message.subject")
	body := gjson.Get(input, "message.body")
	log.Info("MAILGUN_DOMAIN ============================ " + os.Getenv("MAILGUN_DOMAIN"))
	log.Info("Subject = " + subject.String())
	log.Info("Body = " + body.String())

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_KEY"))

	sender := os.Getenv("EMAIL_SENDER")

	recipient := os.Getenv("EMAIL_TEST_RECEIVER")

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject.String(), "", recipient)
	message.SetTemplate("welcome_email")
	err = message.AddTemplateVariable("passwordResetLink", "some link to your site unique to your user")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"success": "false", "message": err})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"success": "false", "message": err})
		return
	}
	log.Info(id)
	c.JSON(200, gin.H{"success": "true", "message": "Email has been sent successfully", "m": resp})
	return

}
