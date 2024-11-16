package router

import (
	usecases "go-clean-arch/internal/Person/App/UseCases"
	entities "go-clean-arch/internal/Person/Domain/Entities"
	repositories "go-clean-arch/internal/Person/Infra/Repositories"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors/wrapper/gin"
)

func NewRouter(router *gin.Engine) *gin.Engine {

	db := make(map[int]entities.PersonEntity, 0)
	repo := repositories.NewPersonRepository(db)
	usecase := usecases.NewPersonUSeCase(repo, )

	r := gin.Default()
	r.POST("/api/v1/person/create")
	r.PATCH("/api/v1/person/update")
	r.DELETE("/api/v1/person/delete")
	r.POST("/api/v1/person/get")
	r.POST("/api/v1/person/search")
	router.Post

}
