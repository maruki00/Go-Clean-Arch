package presenters

import (
	contracts "go-clean-arch/internal/Person/Domain/Contracts"
	viewmodels "go-clean-arch/internal/Person/UserGateway/Adapters/ViewModels"
	viewmodels "person/internal/Person/UserGateway/Adapters/ViewModels"
)

type PersonPresenter struct {
}

func NewJSONAuthPresenter(data any) contracts.ViewModel {
	return viewmodels.NewJsonViewModel(data)
}

func (obj *PersonPresenter) Success(data any) viewmodels.ViewModel {
	return *viewmodels.NewJsonViewModel(data)
}

func (obj *PersonPresenter) Error(data any) viewmodels.ViewModel {
	return *viewmodels.NewJsonViewModel(data)
}
