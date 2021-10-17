package controller

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	kv_key := gjson.Get(request_body, "message.kv_key")
	kv_body := gjson.Get(request_body, "message.kv_body")

	opt, _ := redis.ParseURL(redis_uri)
	client := redis.NewClient(opt)

	// Publish a generated user to the new_users channel
	ctx := context.Background()
	log.Info("kv_key=" + kv_key.String() + ":" + uuid)
	log.Info("kv_value=" + kv_body.String() + ":" + uuid)
	//val := client.Get(ctx, kv_key.String()).Val()

	client.Set(ctx, kv_key.String()+":"+uuid, kv_body.String(), 0)
	c.JSON(200, gin.H{"success": "true", "message": "Successfully Enqueued", "key": kv_key.String() + ":" + uuid})
	return
}
