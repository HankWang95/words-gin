package api

import (
	"github.com/HankWang95/words-gin/app/manager"
	"github.com/HankWang95/words-gin/app/service"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/conv4go"
	"github.com/HankWang95/words-gin/app/form"
	"fmt"
)

type WordsHandler struct {
}

func NewWordsHandler() *WordsHandler {
	return &WordsHandler{}
}

func (this *WordsHandler) RUN(r gin.IRouter) {
	r.GET("/search", this.SearchWord)
	r.POST("/add", this.AddWord)
	r.GET("/all-words-list", this.AllWordsList)
	r.GET("/words-list", this.CreateWordsList)
	r.GET("/redis-test", this.SetRedisWord)
	//r.GET("/", this.SearchWordsList)

}

func (this *WordsHandler) SearchWord(c *gin.Context) {
	c.Request.ParseForm()

	word, uid := c.Request.Form.Get("word"), c.Request.Form.Get("uid")
	fmt.Println(word,uid)
	if uid != "" {
		uid := conv4go.Int64(uid)
		result, status := service.SearchWords4Uid(uid)
		if status == manager.SEARCH_WORD_FAIL {
			c.JSON(404, result)
			return
		}
		c.JSON(200, result)
		return
	}else if word != "" {
		result, status := service.SearchWord4Word(word)
		if status == manager.SEARCH_WORD_FAIL {
			c.JSON(404, result)
			return
		}
		c.JSON(200, result)
		return
	}else {
		c.JSON(404, "请输入查找内容" )
		return
	}


}

func (this *WordsHandler) AllWordsList(c *gin.Context) {
	result := service.AllWordsList()
	var wordList []*form.WordForm
	wordList = result

	//fmt.Println(wordList)

	c.JSON(200, wordList)
	return
}

func (this *WordsHandler) CreateWordsList(c *gin.Context)  {
	result := service.CreateWordsList()
	c.JSON(200,result)
	return
}


func (this *WordsHandler) AddWord(c *gin.Context) {

	c.Request.ParseForm()

	var Word = c.Request.Form.Get("word")
	var Translation = c.Request.Form.Get("translation")
	//data,err := ioutil.ReadAll(c.Request.Body)
	//if err != nil{
	//	return
	//}
	//var m *form.WordForm
	//err = json.Unmarshal(data,&m)
	//if err != nil{
	//	fmt.Println(err)
	//}
	result := service.AddWord(Word, Translation)

	if result == manager.ADD_WORD_SUCCESS {
		c.JSON(200, "添加单词成功")
		return
	} else {
		c.JSON(404, "单词已经存在")
		return
	}





}

func (this *WordsHandler) SetRedisWord(c *gin.Context)  {
	c.Request.ParseForm()

	word := c.Request.Form.Get("word")

	err := service.SetRedisWord(word)
	if err != nil {
		c.JSON(404, "错误")
		return
	}
	c.JSON(200, "set redis word ok!")


}