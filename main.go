package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/form"
	"github.com/smartwalle/validator"
	"github.com/smartwalle/conv4go"
	"errors"
)

func main() {
	LinkDB()
	r := gin.Default()
	r.GET("/add", func(c *gin.Context) {
		//c.Request.ParseForm()

		var f *WordForm
		if err := form.BindWithRequest(c.Request, &f); err != nil {
			fmt.Println(err)
			return
		}

		if val := validator.LazyValidate(&f); val.OK() == false {
			fmt.Println(val.Error())
			return
		}

		//fmt.Println(f.CleanedData)
		fmt.Println(f.CleanedData)

		c.JSON(200,f.CleanedData)


		//c.Request.ParseForm()
		//
		//var hh = c.Request.Form.Get("xxx")
		//
		//c.JSON(200, gin.H{
		//	"message": hh,
		//})
	})

	r.Run() // listen and serve on 0.0.0.0:8080

	//http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
	//	req.ParseForm()
	//
	//
	//	w.Write([]byte(req.Form.Get("aa")))
	//})
	//
	//http.ListenAndServe(":8080", nil)
}

type WordForm struct {
	CleanedData map[string]interface{}
	Word string `form:"word"`
}

func (this *WordForm) WordValidator(str string) error {
	if str == "" {
		return errors.New("cannot none")
	}
	return nil
}


func (this *WordForm) DefaultId() int {
	return 1
}

func (this *WordForm) CleanedId(p string) int {
	return conv4go.Int(p)
}