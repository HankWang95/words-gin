package api

import (
	"fmt"
	"github.com/HankWang95/words-gin/app/manager"
	"github.com/HankWang95/words-gin/app/service"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/conv4go"
	"github.com/HankWang95/words-gin/app/form"
)

type WordsHandler struct {
}

func NewWordsHandler() *WordsHandler {
	return &WordsHandler{}
}

func (this *WordsHandler) RUN(r gin.IRouter) {
	r.GET("/search", this.SearchWord4Uid)
	r.GET("/add", this.AddWord)
	r.GET("/words-list", this.SearchWordsList)
}

func (this *WordsHandler) SearchWord4Uid(c *gin.Context) {
	c.Request.ParseForm()
	//var f = &form.WordForm{}
	//f.Word = c.Request.Form.Get("word")
	//f.Translate = c.Request.Form.Get("translate")
	uid := conv4go.Int64(c.Request.Form.Get("uid"))
	fmt.Println(uid)
	result, status := service.SearchWords4Uid(uid)
	if status == manager.SEARCH_WORD_FAIL {
		c.JSON(404, result)
		return
	}

	c.JSON(200, result)
	return
}

func (this *WordsHandler) SearchWordsList(c *gin.Context) {
	result := service.SearchWordsList()
	var wordList []*form.WordForm
	wordList = result

	//fmt.Println(wordList)

	c.JSON(200, wordList)
	return
}

func (this *WordsHandler) AddWord(c *gin.Context) {

	c.Request.ParseForm()

	var Word = c.Request.Form.Get("word")
	var Translation = c.Request.Form.Get("translation")
	result := service.AddWord(Word, Translation)

	if result == manager.ADD_WORD_SUCCESS {
		c.JSON(200, "添加单词成功")
		return
	} else {
		c.JSON(404, "单词已经存在")
		return
	}

}
