package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/senthilsweb/notifier/pkg/utils"
	"github.com/tidwall/gjson"

	"github.com/go-rod/rod"
)

func Export2PDF(c *gin.Context) {
	request_body := utils.GetStringFromGinRequestBody(c)
	webpage := gjson.Get(request_body, "message.webpage")
	filename := gjson.Get(request_body, "message.filename")
	page := rod.New().MustConnect().MustPage(webpage.String())
	page.MustWaitLoad().MustPDF(filename.String() + ".pdf")
	c.JSON(200, gin.H{"success": "true", "message": "PDF Export was successful"})
	return
}

func Export2PNG(c *gin.Context) {
	request_body := utils.GetStringFromGinRequestBody(c)
	webpage := gjson.Get(request_body, "message.webpage")
	filename := gjson.Get(request_body, "message.filename")
	page := rod.New().MustConnect().MustPage(webpage.String())
	page.MustWaitLoad().MustScreenshot(filename.String() + ".png")
	c.JSON(200, gin.H{"success": "true", "message": "PNG Export was successful"})
	return
}
