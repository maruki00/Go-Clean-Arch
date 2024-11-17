package controllers

import (
	dtos "go-clean-arch/internal/Person/Domain/DTOS"
	ports "go-clean-arch/internal/Person/Domain/Ports"
	requests "go-clean-arch/internal/Person/UserGateway/Requests"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PersonController struct {
	Validate  *validator.Validate
	inputPort ports.PersonInputPort
}

func NewPersonController(
	inputPort ports.PersonInputPort,
) *PersonController {
	return &PersonController{
		Validate:  validator.New(),
		inputPort: inputPort,
	}
}

func (obj *PersonController) Create(ctx *gin.Context) {

	person := &requests.CreatePersonRequest{}
	if err := ctx.BindJSON(&person); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	if err := obj.Validate.Struct(person); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	result := obj.inputPort.CreatePerson(dtos.CreatePersonDTO{
		Name:  person.Name,
		Age:   person.Age,
		Email: person.Email,
	})

	ctx.JSON(200, result.GetResponse())
}

func (obj *PersonController) Update(ctx *gin.Context) {

	person := &requests.UpdatePersonRequest{}
	if err := ctx.BindJSON(&person); err != nil {
		panic(err)
	}

	if err := obj.Validate.Struct(person); err != nil {
		panic(err)
	}

	result := obj.inputPort.UpdatePerson(dtos.UpdatePersonDTO{
		Id:    person.Id,
		Name:  person.Name,
		Age:   person.Age,
		Email: person.Email,
	})

	ctx.JSON(200, result.GetResponse())
}

func (obj *PersonController) Delete(ctx *gin.Context) {

	person := &requests.DeletePersonRequest{}
	if err := ctx.BindJSON(&person); err != nil {
		panic(err)
	}

	if err := obj.Validate.Struct(person); err != nil {
		panic(err)
	}

	result := obj.inputPort.DeletePerson(dtos.DeletePersonDTO{
		Id: person.Id,
	})

	ctx.JSON(200, result.GetResponse())
}

func (obj *PersonController) Get(ctx *gin.Context) {

	id := ctx.Query("id")

	intId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	result := obj.inputPort.GetPersonById(dtos.GetPersonDTO{
		Id: int(intId),
	})

	ctx.JSON(200, result.GetResponse())
}

func (obj *PersonController) Search(ctx *gin.Context) {

	query := ctx.Query("query")

	result := obj.inputPort.Search(dtos.SearchPersonDTO{
		Query: query,
	})

	ctx.JSON(200, result.GetResponse())
}
