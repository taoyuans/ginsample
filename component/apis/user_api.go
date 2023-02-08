package apis

import (
	"fmt"
	"ginsample/component/models"
	configutil "ginsample/config"
	"ginsample/lib/factory"
	"ginsample/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
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

	fmt.Printf("%v", configutil.ConfigValue)
	factory.Logger(ctx.Request.Context()).Warn("log_warn test")

	response.ApiSucc(ctx, http.StatusOK, result)
}
