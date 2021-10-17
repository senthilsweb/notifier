package controller

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/senthilsweb/notifier/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func Enqueue(c *gin.Context) {

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	request_body := utils.GetStringFromGinRequestBody(c)
	//redis_host := utils.GetValElseSetEnvFallback(request_body, "MAILGUN_DOMAIN")
	redis_uri := utils.GetValElseSetEnvFallback(request_body, "REDIS_URI")

	identity := gjson.Get(request_body, "message.identity")
	kv_key := gjson.Get(request_body, "message.kv_key").String()

	kv_body := gjson.Get(request_body, "message.kv_body").String()

	opt, _ := redis.ParseURL(redis_uri)
	client := redis.NewClient(opt)

	if identity.Bool() {
		kv_key = kv_key + ":" + uuid
	}

	// Publish a generated user to the new_users channel
	ctx := context.Background()
	log.Info("kv_key=" + kv_key)
	log.Info("kv_value=" + kv_body)

	client.Set(ctx, kv_key, kv_body, 0)
	c.JSON(200, gin.H{"success": "true", "message": "Successfully Enqueued", "key": kv_key + ":" + uuid})
	return
}

func Dequeue(c *gin.Context) {
	log.Info("Inside Dequeue")
	request_body := utils.GetStringFromGinRequestBody(c)
	kv_key := c.Param("key")
	redis_uri := utils.GetValElseSetEnvFallback(request_body, "REDIS_URI")

	opt, _ := redis.ParseURL(redis_uri)
	client := redis.NewClient(opt)

	ctx := context.Background()

	val := client.Get(ctx, kv_key).Val()
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(val), &jsonMap)
	c.JSON(200, gin.H{"success": "true", "message": "Successfully Dequeued", "data": jsonMap})
	return
}

func Swissknife(c *gin.Context) {
	request_body := utils.GetStringFromGinRequestBody(c)
	number := gjson.Get(request_body, "number").String()
	form_name := gjson.Get(request_body, "form_name").String()
	redis_uri := utils.GetValElseSetEnvFallback(request_body, "REDIS_URI")

	kv_key := ""
	if len(number) > 0 && len(form_name) > 0 {
		kv_key = form_name + ":" + number
	}
	// Publish a generated user to the new_users channel
	ctx := context.Background()
	log.Info("kv_key=" + kv_key)

	opt, _ := redis.ParseURL(redis_uri)
	client := redis.NewClient(opt)
	client.Set(ctx, form_name+":"+number, request_body, 0)
	client.Set(ctx, "last_message", request_body, 0)
	c.JSON(200, gin.H{"success": "true", "message": "Webhook payload was successfully Enqueued", "key": form_name + ":" + number})
	return
}

func GetEnvironment(c *gin.Context) {
	key := c.Param("key")
	val := os.Getenv(key)

	c.JSON(200, gin.H{"success": "true", "message": "Environment variable get attempt was successful", "key": key, "value": val})
	return
}
