package v1

import (
	"fmt"
	"go-gin-demo/global"
	"go-gin-demo/models"
	"go-gin-demo/request"
	"go-gin-demo/response"
	"go-gin-demo/utils"

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
	user := models.User{
		Username: form.Username,
		Password: form.Password,
		Model:    &models.Model{Status: form.Status},
	}
	err1 := global.DB.Create(&user).Error
	if err1 != nil {
		resp.Error(global.ServerError)
	}
	resp.Success(nil)
}
func (u *UserAPI) FindById(c *gin.Context) {
	resp := response.NewResponse(c)
	id := c.Params.ByName("id")
	fmt.Println("id: " + id)
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp.Error(global.ParamError)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("id is: %v", uid)
	fields := []string{"id", "username", "password", "status"}
	user := &models.User{}
	err1 := global.DB.Select(fields).Where("id=?", uid).First(&user).Error
	if err1 != nil {
		resp.Error(global.ParamError)
		fmt.Println(err1.Error())
		return
	}
	resp.Success(user)

}
