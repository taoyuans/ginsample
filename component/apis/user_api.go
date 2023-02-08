package apis

import (
	"fmt"
	"ginsample/component/models"
	configutil "ginsample/config"
	"ginsample/lib/factory"
	"ginsample/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserApi struct {
}

func (c UserApi) Test(ctx *gin.Context) {
	name := ctx.Query("name")
	response.ApiSucc(ctx, http.StatusOK, "Hello "+name)
}

func (c UserApi) GetUsers(ctx *gin.Context) {
	result, err := models.User{}.GetUsers(ctx.Request.Context())
	if err != nil {
		response.ApiFail(ctx, response.ApiErrorDB, err, nil)
	}

	//use config value
	fmt.Printf("【apis】:connfig value test http_port => %s\n", configutil.ConfigValue.HttpPort)

	//add log
	factory.Logger(ctx.Request.Context()).Info("【apis】log test => log_info")
	factory.Logger(ctx.Request.Context()).WithFields(logrus.Fields{
		"test_field": "heihei",
		"msg":        "no msg",
		"request_id": ctx.Request.Header.Get("X-Request-Id"),
	}).Warn()

	response.ApiSucc(ctx, http.StatusOK, result)
}
