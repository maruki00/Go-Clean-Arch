package viewmodels

type ViewModel struct {
	data any
}

func NewJsonViewModel(data any) *ViewModel {
	return &ViewModel{
		data: data,
	}
}

func (v *ViewModel) GetData() any {
	return v.data
}
