package apis

import (
	"net/http"
	"strconv"

	"ginsample/component/models"
	"ginsample/lib/errs"
	"ginsample/lib/response"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (c UserApi) GetByUserId(ctx *gin.Context) {
	userIdStr := ctx.Query("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		response.ApiParameterWarn(ctx, []string{"userId"})
	}
	result, err := models.User{}.GetByUserId(ctx.Request.Context(), userId)
	if err != nil {
		response.ApiFail(ctx, errs.Trace(err))
		return
	}

	response.ApiSucc(ctx, http.StatusOK, result)
}

// 用账号密码登陆
func (c UserApi) LoginByUserName(ctx *gin.Context) {
	var loginUser struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&loginUser); err != nil {
		response.ApiParameterWarn(ctx, []string{"userName", "password"})
		return
	}

	if len(loginUser.UserName) <= 0 || len(loginUser.Password) <= 0 {
		response.ApiParameterWarn(ctx, []string{"userName", "password"})
		return
	}

	tokens, err := models.Login{}.LoginByUserName(ctx.Request.Context(), loginUser.UserName, loginUser.Password)
	if err != nil {
		// ctx.SetCookie("user", "", 0, "/", cconfig.ConfigValue.Domain, true, true)
		response.ApiFail(ctx, errs.Trace(err))
		return
	}
	// ctx.SetCookie("user", fmt.Sprint(tokens), 0, "/", cconfig.ConfigValue.Domain, true, true)
	response.ApiSucc(ctx, http.StatusOK, tokens)
}
