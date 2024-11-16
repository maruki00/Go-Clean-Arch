package viewmodels

import contracts "go-clean-arch/internal/Person/Domain/Contracts"

type ViewModel struct {
	data any
}

func NewJsonViewModel(data ResponseModel) contracts.ViewModel {
	return &ViewModel{
		data: data,
	}
}

func (v *ViewModel) GetResponse() any {
	return v.data
}
