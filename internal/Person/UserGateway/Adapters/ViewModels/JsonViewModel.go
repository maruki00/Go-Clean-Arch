package viewmodels

import (
	"encoding/json"
	contracts "go-clean-arch/internal/Person/Domain/Contracts"
	shared_entities "go-clean-arch/internal/Shared/Domain/Entities"
)

var _ contracts.ViewModel = (*JsonViewModel)(nil)

type JsonViewModel struct {
	data shared_entities.ResponseModel
}

func NewJsonViewModel(data shared_entities.ResponseModel) contracts.ViewModel {
	return &JsonViewModel{
		data: data,
	}
}

func (v *JsonViewModel) GetResponse() any {
	return v.data
}

func (v *JsonViewModel) ToJson() string {
	d, err := json.Marshal(v.data)
	if err != nil {
		panic(err.Error())
	}
	return string(d)
}
