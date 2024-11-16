package presenters

import (
	contracts "go-clean-arch/internal/Person/Domain/Contracts"
	shared_entities "go-clean-arch/internal/Shared/Domain/Entities"
	viewmodels "person/internal/Person/UserGateway/Adapters/ViewModels"
)

type PersonPresenter struct {
}

func NewJSONAuthPresenter() contracts.ViewModel {
	return PersonPresenter{}
}

func (obj *PersonPresenter) Success(data any) viewmodels.ViewModel {
	return *viewmodels.NewJsonViewModel(data)
}

func (obj *PersonPresenter) Error(data any) viewmodels.ViewModel {
	return *viewmodels.NewJsonViewModel(data)
}
