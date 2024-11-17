package presenters

import (
	contracts "go-clean-arch/internal/Person/Domain/Contracts"
	viewmodels "go-clean-arch/internal/Person/UserGateway/Adapters/ViewModels"
	shared_entities "go-clean-arch/internal/Shared/Domain/Entities"
)

// var _ contracts.ViewModel = (*)(nil)

type PersonPresenter struct {
}

func NewJSONAuthPresenter() *PersonPresenter {
	return &PersonPresenter{}
}

func (obj *PersonPresenter) Success(data shared_entities.ResponseModel) contracts.ViewModel {
	return viewmodels.NewJsonViewModel(data)
}

func (obj *PersonPresenter) Error(data shared_entities.ResponseModel) contracts.ViewModel {
	return viewmodels.NewJsonViewModel(data)
}
