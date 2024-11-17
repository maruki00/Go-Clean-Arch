package usecases

import (
	contracts "go-clean-arch/internal/Person/Domain/Contracts"
	dtos "go-clean-arch/internal/Person/Domain/DTOS"
	ports "go-clean-arch/internal/Person/Domain/Ports"
	valueobject "go-clean-arch/internal/Person/Domain/ValueObjects"
	models "go-clean-arch/internal/Person/Infra/Models"
	formatters "go-clean-arch/internal/Person/UserGateway/Formatters"
	shared_entities "go-clean-arch/internal/Shared/Domain/Entities"
)

type PersonUseCases struct {
	repository contracts.IPersonRepository
	outputPort ports.PersonOutputPort
	formatter  *formatters.PersonFormmater
}

func NewPersonUSeCase(
	repo contracts.IPersonRepository,
	outputport ports.PersonOutputPort,
) *PersonUseCases {
	return &PersonUseCases{
		repository: repo,
		outputPort: outputport,
		formatter:  &formatters.PersonFormmater{},
	}
}

func (obj *PersonUseCases) CreatePerson(dto dtos.CreatePersonDTO) contracts.ViewModel {
	entity, err := obj.repository.Create(&models.Person{
		Name:  dto.Name,
		Email: valueobject.NewEmail(dto.Email),
		Age:   valueobject.NewAge(dto.Age),
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
		Data:    obj.formatter.FromOne(entity),
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
		Data:    obj.formatter.FromOne(result),
	})
}

func (obj *PersonUseCases) UpdatePerson(dto dtos.UpdatePersonDTO) contracts.ViewModel {

	entity, err := obj.repository.Update(&models.Person{
		Id:    0,
		Name:  dto.Name,
		Email: valueobject.NewEmail(dto.Email),
		Age:   valueobject.NewAge(dto.Age),
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
		Data:    obj.formatter.FromOne(entity),
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
		Data:    obj.formatter.FromOne(entity),
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
		Data:    obj.formatter.FromMany(entity),
	})
}
