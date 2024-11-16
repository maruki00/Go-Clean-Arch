package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors/wrapper/gin"
)

func NewRouter(router *gin.Engine) *gin.Engine {

	r := gin.Default()
	r.Post()
	router.Post

}
