package global

import (
	"go-gin-demo/internal/response"
)

var (
	OK          = response.NewError(0, "OK")
	ParamError  = response.NewError(400, "参数不合法")
	NotFound    = response.NewError(404, "资源未找到")
	ServerError = response.NewError(500, "服务器错误")

	CreateError = response.NewError(10001, "创建失败")
	UpdateError = response.NewError(10002, "创建失败")
	DeleteError = response.NewError(10003, "创建错误")
)
