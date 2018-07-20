package api

import "github.com/gin-gonic/gin"

func RUN(r gin.IRouter) {
	r = r.Group("/api")
	NewWordsHandler().RUN(r)
}
