package apis

import (
	"ginsample/component/models"
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
	result, err := models.User{}.GetApps(ctx.Request.Context())
	if err != nil {
		response.ApiFail(ctx, response.ApiErrorDB, err, nil)
	}

	response.ApiSucc(ctx, http.StatusOK, result)
}
