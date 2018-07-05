package app

import (
	"github.com/gin-gonic/gin"
	"github.com/HankWang95/words-gin/app/api"
)

func RUN() error {
	var s = gin.Default()

	err := NewDBLink()
	if err != nil {
		return err
	}
	api.RUN(s)

}