package server

import (
	"github.com/gin-gonic/gin"
	"github.com/HankWang95/words-gin/app/api"
)

func RUN()  {
	var s = gin.Default()
	api.RUN(s)
	s.Run()

}