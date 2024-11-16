package ports

import (
	contracts "go-clean-arch/internal/Person/Domain/Contracts"
	dtos "go-clean-arch/internal/Person/Domain/DTOS"
)

type PersonInputPort interface {
	CreatePerson(dto dtos.CreatePersonDTO) contracts.ViewModel
	DeletePerson(dto dtos.DeletePersonDTO) contracts.ViewModel
	UpdatePerson(dto dtos.UpdatePersonDTO) contracts.ViewModel
	GetPersonById(dto dtos.GetPersonDTO) contracts.ViewModel
	Search(dto dtos.SearchPersonDTO) contracts.ViewModel
}
