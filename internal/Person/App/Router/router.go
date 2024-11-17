package router

import (
	usecases "go-clean-arch/internal/Person/App/UseCases"
	entities "go-clean-arch/internal/Person/Domain/Entities"
	repositories "go-clean-arch/internal/Person/Infra/Repositories"
	presenters "go-clean-arch/internal/Person/UserGateway/Adapters/Presenters"
	controllers "go-clean-arch/internal/Person/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) {

	db := make(map[int]entities.PersonEntity, 0)
	repo := repositories.NewPersonRepository(db)
	usecases := usecases.NewPersonUSeCase(repo, presenters.NewJSONAuthPresenter())

	controller := controllers.NewPersonController(usecases)

	router.POST("/api/v1/person/create", controller.Create)
	router.PATCH("/api/v1/person/update", controller.Update)
	router.DELETE("/api/v1/person/delete", controller.Delete)
	router.GET("/api/v1/person/get", controller.Get)
	router.GET("/api/v1/person/search", controller.Search)
}
