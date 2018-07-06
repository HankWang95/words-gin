package api

import (
	"github.com/gin-gonic/gin"
	"github.com/HankWang95/words-gin/app/service"
)

type HotelHandler struct {

}

func NewHotelHandler() *HotelHandler {
	return &HotelHandler{}
}

func (this *HotelHandler) RUN(r gin.IRouter)  {
	r.GET("/search-hotel-user", this.SearchHotelUser)
}

func (this *HotelHandler) SearchHotelUser(c *gin.Context)  {
	res := service.SearchHotelUser()
	if res == nil {
		c.JSON(404, "error")
	}
	c.JSON(200,res)

}