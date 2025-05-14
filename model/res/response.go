package res

import (
	"gin_web_frame/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 0
	ERROR   = 7
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "ok", c)
}

func OkWithMassage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}
func OkWith(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "ok", c)
}

// OkWithList 响应分页操作
func OkWithList[T any](list []T, count int64, c *gin.Context) {
	OkWithData(gin.H{"list": list, "count": count}, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMassage(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

// 用于绑定请求参数错误的响应
func FailWithError(err error, obj any, c *gin.Context) {
	msg := common.GetValidMsg(err, obj)
	FailWithMassage(msg, c)
}

// 根据code 查询出msg
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, b := ErrorMap[code]
	// 若有该错误，则取其内容
	if b {
		Result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	//	若没有该错误
	Result(ERROR, map[string]interface{}{}, "未知错误", c)
}
