package main

import (
	"fmt"
	"github.com/HankWang95/words-gin/word/service"
	"github.com/HankWang95/words-gin/word/service/repository/mysql"
	"github.com/HankWang95/words-gin/word/transport/restful"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smartwalle/dbs"
)

func main() {
	var db, err = dbs.NewSQL("mysql", "root:whc5608698@tcp(127.0.0.1:3306)/words?charset=utf8&parseTime=true", 10, 5)
	if err != nil {
		fmt.Println(err)
	}
	db.Ping()
	var wRepo = mysql.NewWordRepository(db)
	var wServ = service.NewWordService(wRepo)
	var wHandler = restful.NewWordHandler(wServ)

	var s = gin.Default()
	wHandler.Run(s)
	s.Run(":8888")

}
