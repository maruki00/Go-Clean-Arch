package internal_usergateway_controllers

import (
	ports "go-clean-arch/internal/Person/Domain/Ports"
	internal_person_app_usecases "person/internal/Person/App/UseCases"
	internal_person_domain_dtos "person/internal/Person/Domain/DTOS"
	requests "person/internal/Person/UserGateway/Requests"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PersonController struct {
	useCase  internal_person_app_usecases.PersonUseCases
	Validate *validator.Validate
	inputPort ports.PersonInputPort
}

func NewPersonController(
	useCase internal_person_app_usecases.PersonUseCases,
	inputPort ports.PersonInputPort,
) *PersonController {
	return &PersonController{
		useCase:  useCase,
		Validate: validator.New(),
		inputPort: inputPort,
	}
}

func (obj *PersonController) Create() *gin.HandlerFunc {
	return func(ctx *gin.Context) {
		person := &requests.CreatePersonRequest{}
		if err := ctx.BindJson(&person); err != nil {
			panic(err)
		}

		if err := obj.Validate.Struct(person); err != nil {
			panic(err)
		}

		result, err := obj.useCase.CreatePerson(internal_person_domain_dtos.CreatePersonDTO{
			Name:  person.Name,
			Age:   person.Age,
			Email: person.Email,
		})

		if err != nil {
			panic(err)
		}

		return ctx.JSON{
			200,
			result.GetResponse()
		} 
	}
}
