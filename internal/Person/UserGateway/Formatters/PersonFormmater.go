package formatters

import entities "go-clean-arch/internal/Person/Domain/Entities"

type PersonFormmater struct {
}

func (obj *PersonFormmater) FromOne(person entities.PersonEntity) map[string]any {
	return map[string]any{
		"id":    person.GetId(),
		"name":  person.GetName(),
		"email": person.GetEmail(),
		"age":   person.GetAge(),
	}
}

func (obj *PersonFormmater) FromMany(people []entities.PersonEntity) []map[string]any {
	ret := make([]map[string]any, 0)
	for _, p := range people {
		ret = append(ret, obj.FromOne(p))
	}

	return ret
}
