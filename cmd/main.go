package main

import (
	router "go-clean-arch/internal/Person/App/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.NewRouter(r)

	r.Run(":3000")
}
