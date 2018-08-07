package restful

import (
	"github.com/HankWang95/words-gin/word"
	"github.com/gin-gonic/gin"
)

type WordHandler struct {
	WordServer word.WordService
}

func NewWordHandler(wServ word.WordService) *WordHandler {
	var h = &WordHandler{}
	h.WordServer = wServ
	return h
}

func (this *WordHandler) Run(r gin.IRouter) {
	r.GET("/search", this.SearchWord)
}

func (this *WordHandler) SearchWord(c *gin.Context) {
	var word, err = this.WordServer.SearchWord(c.Request.FormValue("word"))
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, word)
}
