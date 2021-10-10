package controller

import (
	"bytes"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/mailgun"
	"github.com/tidwall/gjson"
)

func NotifyMailgun(c *gin.Context) {

	notifier := notify.New()

	// Provide your mailgun details
	//os.Getenv("")
	mailgunService := mailgun.New(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_KEY"), os.Getenv("EMAIL_SENDER"))
	mailgunService.AddReceivers(os.Getenv("EMAIL_TEST_RECEIVER"))
	notifier.UseServices(mailgunService)

	jsonData, err := c.GetRawData()
	input := bytes.NewBuffer(jsonData).String()

	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"success": "false", "message": err})
		return
	}

	subject := gjson.Get(input, "message.subject")
	body := gjson.Get(input, "message.body")
	log.Info("Subject = " + subject.String())
	log.Info("Body = " + body.String())
	// Send a message
	err = notifier.Send(
		context.Background(),
		subject.String(),
		body.String(),
	)

	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"success": "false", "message": err})
		return
	}

	c.JSON(200, gin.H{"success": "true", "message": "Email has been sent successfully"})
	return
}
