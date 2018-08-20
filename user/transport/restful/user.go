package restful

import (
	"github.com/HankWang95/words-gin/user"
	"github.com/HankWang95/words-gin/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserServer user.UserServeice
}

func NewUserHandler(uServ user.UserServeice) *UserHandler {
	return &UserHandler{UserServer: uServ}
}

func (this *UserHandler) Run(r gin.IRouter) {
	r.POST("/login", this.Login)
	r.GET("/logout", this.Logout)
	r.POST("/sign-up", this.SignUp)
}

func (this *UserHandler) Login(c *gin.Context) {
	var data *user.User
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, "格式不正确")
		return
	}
	status := this.UserServer.Login(data.Name, data.Password)

	if status != service.USER_STATUS_OK {
		c.JSON(http.StatusBadRequest, "用户不可用")
		return
	}
	c.JSON(http.StatusOK, "登陆成功")
	return

}

func (this *UserHandler) Logout(c *gin.Context)  {
	this.UserServer.Logout()
	return
}

func (this *UserHandler) SignUp(c *gin.Context) {
	var data *user.User
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, "格式不正确")
		return
	}

	status, err := this.UserServer.SignUp(data.Name, data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	if status != service.USER_STATUS_OK {
		c.JSON(http.StatusBadRequest, "注册失败")
		return
	}
	c.JSON(http.StatusOK, "注册成功")
}