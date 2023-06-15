package response

import (
	"fmt"
	"net/http"

	"ginsample/lib/factory"

	"github.com/gin-gonic/gin"
)

type ApiResult struct {
	Rows    interface{} `json:"rows"`
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Details interface{} `json:"details,omitempty"`
}

type ApiError struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"msg"`
	Details interface{} `json:"details,omitempty"`
}

type ArrayApiResult struct {
	Rows    interface{} `json:"rows"`
	Code    int         `json:"code"`
	Total   int64       `json:"total"`
	Message string      `json:"msg"`
	Details interface{} `json:"details,omitempty"`
}

var (
	// System Error
	ApiErrorSystem             = ApiError{Code: 10001, Message: "System Error"}
	ApiErrorServiceUnavailable = ApiError{Code: 10002, Message: "Service unavailable"}
	ApiErrorRemoteService      = ApiError{Code: 10003, Message: "Remote service error"}
	ApiErrorIPLimit            = ApiError{Code: 10004, Message: "IP limit"}
	ApiErrorPermissionDenied   = ApiError{Code: 10005, Message: "Permission denied"}
	ApiErrorIllegalRequest     = ApiError{Code: 10006, Message: "Illegal request"}
	ApiErrorHTTPMethod         = ApiError{Code: 10007, Message: "HTTP method is not suported for this request"}
	ApiErrorParameter          = ApiError{Code: 10008, Message: "Parameter error"}
	ApiErrorMissParameter      = ApiError{Code: 10009, Message: "Miss required parameter"}
	ApiErrorDB                 = ApiError{Code: 10010, Message: "DB error, please contact the administator"}
	ApiErrorTokenInvaild       = ApiError{Code: 10011, Message: "Token invaild"}
	ApiErrorMissToken          = ApiError{Code: 10012, Message: "Miss token"}
	ApiErrorVersion            = ApiError{Code: 10013, Message: "API version %s invalid"}
	ApiErrorNotFound           = ApiError{Code: 10014, Message: "Resource not found"}
	// Business Error
	ApiErrorUserNotExists = ApiError{Code: 20001, Message: "User does not exists"}
	ApiErrorPassword      = ApiError{Code: 20002, Message: "Password error"}
	ApiErrorWechatContext = ApiError{Code: 20006, Message: "Please operate in the qiye wechat"}

	ApiErrorLogin   = ApiError{Code: 30002, Message: "Login failure"}
	ApiErrorDetails = ApiError{Code: 30003, Message: "Details are as follows"}

	ApiErrorFailedVerify = ApiError{Code: 40005, Message: "failed to verify"}
)

func ArrayApiSucc(ctx *gin.Context, status int, total int64, rows interface{}) {
	apiResult := ApiResult{
		Code: status,
		Rows: ArrayApiResult{Total: total, Rows: rows},
	}

	factory.Logger(ctx.Request.Context()).Info().Msg(fmt.Sprintf("%+v", apiResult))

	ctx.JSON(status, apiResult)
}
func ApiSucc(ctx *gin.Context, status int, result interface{}) {
	apiResult := ApiResult{
		Code: status,
		Rows: result,
	}

	factory.Logger(ctx.Request.Context()).Info().Msg(fmt.Sprintf("%+v", apiResult))

	ctx.JSON(status, apiResult)
}

func ReturnApiWarn(ctx *gin.Context, status int, msg string, err error) {

	apiResult := ApiResult{
		Code:    status,
		Message: msg,
		Details: fmt.Sprint(err),
	}

	factory.Logger(ctx.Request.Context()).Warn().Msg(fmt.Sprintf("%+v", apiResult))

	ctx.JSON(status, apiResult)
}

func ApiParameterWarn(ctx *gin.Context, parameters []string) {
	apiResult := ApiResult{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("参数错误： %v", parameters),
		Details: fmt.Sprint(parameters),
	}

	factory.Logger(ctx.Request.Context()).Warn().Msg(fmt.Sprintf("%+v", apiResult))

	ctx.JSON(http.StatusBadRequest, apiResult)
}

func ApiFail(ctx *gin.Context, err error) {
	apiResult := ApiResult{
		Code:    http.StatusInternalServerError,
		Message: "系统出错，请联系管理员查看！",
		Details: fmt.Sprint(err),
	}

	factory.Logger(ctx.Request.Context()).Error().Msg(fmt.Sprintf("%+v", apiResult))

	ctx.JSON(http.StatusInternalServerError, apiResult)
}

func ApiFailWithMsg(ctx *gin.Context, status int, msg string, err error) {
	apiResult := ApiResult{
		Code:    status,
		Message: msg,
		Details: fmt.Sprint(err),
	}

	factory.Logger(ctx.Request.Context()).Error().Msg(fmt.Sprintf("%+v", apiResult))

	ctx.JSON(status, apiResult)
}

func ApiFailWithMultiMsg(ctx *gin.Context, status int, err error, v ...interface{}) {
	apiResult := ApiResult{
		Code:    status,
		Message: fmt.Sprint(v...),
		Details: fmt.Sprint(err),
	}

	factory.Logger(ctx.Request.Context()).Error().Msg(fmt.Sprintf("%+v", apiResult))

	ctx.JSON(status, apiResult)
}
