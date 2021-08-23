package v1

import (
	"fmt"
	global2 "go-gin-demo/internal/global"
	models2 "go-gin-demo/internal/models"
	"go-gin-demo/internal/request"
	"go-gin-demo/internal/response"
	"go-gin-demo/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
}

func NewUserAPI() UserAPI {
	return UserAPI{}
}
func (u *UserAPI) Add(c *gin.Context) {
	resp := response.NewResponse(c)
	form := &request.AddForm{}
	vaild, err := utils.BindAndValid(c, form)
	fmt.Println(err)
	if !vaild {
		resp.ErrorCode(400, err.Error())
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%v", form)
	user := models2.User{
		Username: form.Username,
		Password: form.Password,
		Model:    &models2.Model{Status: form.Status},
	}
	err1 := global2.DB.Create(&user).Error
	if err1 != nil {
		resp.Error(global2.ServerError)
	}
	resp.Success(nil)
}
func (u *UserAPI) FindById(c *gin.Context) {
	resp := response.NewResponse(c)
	id := c.Params.ByName("id")
	fmt.Println("id: " + id)
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp.Error(global2.ParamError)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("id is: %v", uid)
	fields := []string{"id", "username", "password", "status"}
	user := &models2.User{}
	err1 := global2.DB.Select(fields).Where("id=?", uid).First(&user).Error
	if err1 != nil {
		resp.Error(global2.ParamError)
		fmt.Println(err1.Error())
		return
	}
	resp.Success(user)

}
