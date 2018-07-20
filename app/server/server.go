package server

import (
	"github.com/HankWang95/words-gin/app/api"
	"github.com/gin-gonic/gin"
)

func RUN() {
	var s = gin.Default()

	api.RUN(s)
	s.Run()

}
