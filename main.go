package main

import (
	"github.com/HankWang95/words-gin/app/server"
)

func main()  {
	server.RUN()
}



//var db = app.NewDBLink()
//
//func main() {
//	r := gin.Default()
//
//
//	r.GET("/add", Add)
//
//	// listen and serve on 0.0.0.0:8080
//
//	//http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
//	//	req.ParseForm()
//	//
//	//
//	//	w.Write([]byte(req.Form.Get("aa")))
//	//})
//	//
//	//http.ListenAndServe(":8080", nil)
//
//	r.GET("/search", func(c *gin.Context) {
//		c.Request.ParseForm()
//		var f *WordForm
//		if err := form.BindWithRequest(c.Request, &f); err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		//id := f.CleanedId(c.Request.Form.Get("id"))
//
//		fmt.Println(f.Id)
//
//		//return
//
//		word, err := app.SearchWord(db, f.Id)
//
//
//
//		if err != nil {
//			fmt.Println("查询失败")
//		}
//		c.JSON(200, gin.H{
//			"word_is": word,
//		})
//
//	})
//
//	r.Run()
//}
//
//type WordForm struct {
//	CleanedData map[string]interface{}
//	Word        string `form:"word"`
//	Id          int    `form:"id"`
//}
//
//func (this *WordForm) WordValidator(str string) error {
//	if str == "" {
//		return errors.New("cannot none")
//	}
//	return nil
//}
//
//func (this *WordForm) DefaultId() int {
//	return 1
//}

//func (this *WordForm) CleanedId(p string) int {
//	return conv4go.Int(p)
//}

//
//func Add(c *gin.Context) {
//
//	c.Request.ParseForm()
//
//	var f *WordForm
//	if err := form.BindWithRequest(c.Request, &f); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if val := validator.LazyValidate(&f); val.OK() == false {
//		fmt.Println(val.Error())
//		return
//	}
//	var word = c.Request.Form.Get("word")
//
//	err := app.AddWord(db, word)
//	if err != nil {
//		fmt.Println("加入单词失败")
//	}
//
//	c.JSON(200, f.CleanedData)

	//c.Request.ParseForm()
	//
	//var hh = c.Request.Form.Get("xxx")
	//
	//c.JSON(200, gin.H{
	//	"message": hh,
	//})
//}