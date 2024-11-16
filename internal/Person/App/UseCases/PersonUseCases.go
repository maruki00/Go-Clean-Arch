package usecases

import (
	shared_entities "command-line-arguments/home/user/Desktop/person/internal/Shared/Domain/Entities/ResponseModel.go"
	contracts "person/internal/Person/Domain/Contracts"
	dtos "person/internal/Person/Domain/DTOS"
	ports "person/internal/Person/Domain/Ports"
	models "person/internal/Person/Infra/Models"
)

type PersonUseCases struct {
	repository   contracts.IPersonRepository
	outputPort   ports.PersonOutputPort
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
	})

	if err != nil {
		return obj.outputPort.Error(&shared_entities.ResponseModel{
			Message: err.Error(),
			Status: 400,
			Data: nil,
		})
	}
	
	return obj.outputPort.Success(
}


func (obj *PersonUseCases) DeletePerson(dto dtos.DeletePersonDTO) contracts.ViewModel {
	
	return nil
}


func (obj *PersonUseCases) GetPersonById(dto dtos.DeletePersonDTO) contracts.ViewModel {
	
	return nil
}

func (obj *PersonUseCases) UpdatePerson(dto dtos.DeletePersonDTO) contracts.ViewModel {
	
	return nil
}

func (obj *PersonUseCases) Search(dto dtos.DeletePersonDTO) contracts.ViewModel {
	
	return nil
}
