package myapp

import (
	"github.com/HankWang95/words-gin/user/service"
	"github.com/HankWang95/words-gin/user/service/repository/mysql"
	"github.com/HankWang95/words-gin/user/transport/restful"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/dbs"
)

func Run(db dbs.DB) {
	var uRepo = mysql.NewUserRepository(db)
	var uServ = service.NewUserService(uRepo)
	var uHandler = restful.NewUserHandler(uServ)

	var s = gin.Default()
	uHandler.Run(s)
	s.Run(":7777")
}
