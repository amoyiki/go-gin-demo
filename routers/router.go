package routers

import (
	v1 "go-gin-demo/api/v1"
	"go-gin-demo/response"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	userApi := v1.NewUserAPI()
	r.NoRoute(HandlerNotFound)
	r.NoMethod(HandlerNotFound)
	r.Use(Recover)
	v := r.Group("/v1")
	{
		v.GET("/user/:id", userApi.FindById)
		v.POST("/user", userApi.Add)
	}
	return r
}

func HandlerNotFound(c *gin.Context) {
	response.NewResponse(c).Error(404, "资源未找到")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic %v\n", r)
			debug.PrintStack()
			response.NewResponse(c).Error(500, "服务器内部错误")
		}
	}()
	c.Next()
}
