package controllers

import (
	usecases "go-clean-arch/internal/Person/App/UseCases"
	dtos "go-clean-arch/internal/Person/Domain/DTOS"
	ports "go-clean-arch/internal/Person/Domain/Ports"
	requests "go-clean-arch/internal/Person/UserGateway/Requests"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PersonController struct {
	useCase   usecases.PersonUseCases
	Validate  *validator.Validate
	inputPort ports.PersonInputPort
}

func NewPersonController(
	useCase usecases.PersonUseCases,
	inputPort ports.PersonInputPort,
) *PersonController {
	return &PersonController{
		useCase:   useCase,
		Validate:  validator.New(),
		inputPort: inputPort,
	}
}

func (obj *PersonController) Create(ctx *gin.Context) {

	person := &requests.CreatePersonRequest{}
	if err := ctx.BindJSON(&person); err != nil {
		panic(err)
	}

	if err := obj.Validate.Struct(person); err != nil {
		panic(err)
	}

	result := obj.useCase.CreatePerson(dtos.CreatePersonDTO{
		Name:  person.Name,
		Age:   person.Age,
		Email: person.Email,
	})

	ctx.JSON(200, result.GetResponse())
}
