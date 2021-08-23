package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseContent struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func NewError(code int, msg string) ResponseContent {
	return ResponseContent{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	}
}
func (r *Response) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := ResponseContent{}
	res.Code = 0
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK, res)
}

func (r *Response) ErrorCode(code int, msg string) {
	res := ResponseContent{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK, res)
	r.Ctx.Abort()
}

func (r *Response) Error(res ResponseContent) {
	r.Ctx.JSON(http.StatusOK, res)
	r.Ctx.Abort()

}
