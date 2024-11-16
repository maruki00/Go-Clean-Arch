package presenters

import viewmodels "person/internal/Person/UserGateway/Adapters/ViewModels"

type PersonPresenter struct {
}

func NewJSONAuthPresenter() JsonPersonPresenter {
	return PersonPresenter{}
}

func (obj *PersonPresenter) Success(data any) viewmodels.ViewModel {
	return *viewmodels.NewJsonViewModel(data)
}

func (obj *PersonPresenter) Error(data any) viewmodels.ViewModel {
	return *viewmodels.NewJsonViewModel(data)
}
