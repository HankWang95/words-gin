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
	r.GET("/search/chinese", this.SearchWordWithChinese)
	r.GET("/search/english", this.SearchWordWithEnglish)
}

func (this *WordHandler) SearchWordWithChinese(c *gin.Context) {
	var word, err = this.WordServer.SearchWordWithChinese(c.Request.FormValue("word"))
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, word)
}

func (this *WordHandler) SearchWordWithEnglish(c *gin.Context) {
	var word, err = this.WordServer.SearchWordWithEnglish(c.Request.FormValue("word"))
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, word)
}
