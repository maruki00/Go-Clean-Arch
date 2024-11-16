package usecases

import (
	contracts "go-clean-arch/internal/Person/Domain/Contracts"
	dtos "go-clean-arch/internal/Person/Domain/DTOS"
	ports "go-clean-arch/internal/Person/Domain/Ports"
	models "go-clean-arch/internal/Person/Infra/Models"
	shared_entities "go-clean-arch/internal/Shared/Domain/Entities"
)

type PersonUseCases struct {
	repository contracts.IPersonRepository
	outputPort ports.PersonOutputPort
}

func NewPersonUSeCase(
	repo contracts.IPersonRepository,
	outputport ports.PersonOutputPort,
) *PersonUseCases {
	return &PersonUseCases{
		repository: repo,
		outputPort: outputport,
	}
}

func (obj *PersonUseCases) CreatePerson(dto dtos.CreatePersonDTO) contracts.ViewModel {
	entity, err := obj.repository.Create(&models.Person{
		Id:   0,
		Name: dto.Name,
		Age:  dto.Age,
	})

	if err != nil {
		return obj.outputPort.Error(shared_entities.ResponseModel{
			Message: err.Error(),
			Status:  400,
			Data:    nil,
		})
	}

	return obj.outputPort.Success(shared_entities.ResponseModel{
		Message: "Success",
		Status:  200,
		Data:    entity,
	})
}

func (obj *PersonUseCases) DeletePerson(dto dtos.DeletePersonDTO) contracts.ViewModel {

	result, err := obj.repository.Delete(dto.Id)

	if err != nil {
		return obj.outputPort.Error(shared_entities.ResponseModel{
			Message: err.Error(),
			Status:  400,
			Data:    nil,
		})
	}

	return obj.outputPort.Success(shared_entities.ResponseModel{
		Message: "Success",
		Status:  200,
		Data:    result,
	})
}

func (obj *PersonUseCases) UpdatePerson(dto dtos.UpdatePersonDTO) contracts.ViewModel {

	entity, err := obj.repository.Update(&models.Person{
		Id:   dto.Id,
		Name: dto.Name,
		Age:  dto.Age,
	})

	if err != nil {
		return obj.outputPort.Error(shared_entities.ResponseModel{
			Message: err.Error(),
			Status:  400,
			Data:    nil,
		})
	}

	return obj.outputPort.Success(shared_entities.ResponseModel{
		Message: "Success",
		Status:  200,
		Data:    entity,
	})
}

func (obj *PersonUseCases) GetPersonById(dto dtos.GetPersonDTO) contracts.ViewModel {
	entity, err := obj.repository.GetById(dto.Id)

	if err != nil {
		return obj.outputPort.Error(shared_entities.ResponseModel{
			Message: err.Error(),
			Status:  400,
			Data:    nil,
		})
	}

	return obj.outputPort.Success(shared_entities.ResponseModel{
		Message: "Success",
		Status:  200,
		Data:    entity,
	})
}

func (obj *PersonUseCases) Search(dto dtos.SearchPersonDTO) contracts.ViewModel {
	entity, err := obj.repository.Search(dto.Query)

	if err != nil {
		return obj.outputPort.Error(shared_entities.ResponseModel{
			Message: err.Error(),
			Status:  400,
			Data:    nil,
		})
	}

	return obj.outputPort.Success(shared_entities.ResponseModel{
		Message: "Success",
		Status:  200,
		Data:    entity,
	})
}
