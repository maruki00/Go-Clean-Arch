package ports

import (
	contracts "command-line-arguments/home/user/Desktop/person/internal/Person/Domain/Contracts/ViewModel.go"
	shared_entities "command-line-arguments/home/user/Desktop/person/internal/Shared/Domain/Entities/ResponseModel.go"
)


type PersonOutputPort interface {
	Success(data shared_entities.ResponseModel) contracts.ViewModel
	Error(data shared_entities.ResponseModel) contracts.ViewModel
}
