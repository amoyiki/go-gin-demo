package routers

import (
	v1 "go-gin-demo/api/v1"
	"go-gin-demo/global"
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
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/user/:id", userApi.FindById)
		apiV1.POST("/user", userApi.Add)
	}
	return r
}

func HandlerNotFound(c *gin.Context) {
	response.NewResponse(c).Error(global.NotFound)
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic %v\n", r)
			debug.PrintStack()
			response.NewResponse(c).Error(global.ServerError)
		}
	}()
	c.Next()
}
