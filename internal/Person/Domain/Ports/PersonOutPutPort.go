package ports

import (
	contracts "go-clean-arch/internal/Person/Domain/Contracts"
	shared_entities "go-clean-arch/internal/Shared/Domain/Entities"
)

type PersonOutputPort interface {
	Success(data shared_entities.ResponseModel) contracts.ViewModel
	Error(data shared_entities.ResponseModel) contracts.ViewModel
}
