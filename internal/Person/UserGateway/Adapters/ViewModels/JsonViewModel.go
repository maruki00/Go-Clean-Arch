package viewmodels

import contracts "go-clean-arch/internal/Person/Domain/Contracts"

type ViewModel struct {
	data any
}

func NewJsonViewModel(data any) contracts.ViewModel {
	return &ViewModel{
		data: data,
	}
}

func (v *ViewModel) GetData() any {
	return v.data
}
