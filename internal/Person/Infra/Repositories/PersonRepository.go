package repositories

import (
	"errors"
	contracts "person/internal/Person/Domain/Contracts"
	entities "person/internal/Person/Domain/Entities"
)

type PersonRepository struct {
	db map[int]entities.PersonEntity
}

func NewPersonRepository(d map[int]entities.PersonEntity) contracts.IPersonRepository {
	return &PersonRepository{
		db: d,
	}
}

func (obj *PersonRepository) Create(person entities.PersonEntity) (entities.PersonEntity, error) {
	_, ok := obj.db[person.GetId()]
	if ok {
		return nil, errors.New("record already exists")
	}
	obj.db[person.GetId()] = person
	return person, nil
}

func (obj *PersonRepository) Update(person entities.PersonEntity) (entities.PersonEntity, error) {
	_, ok := obj.db[person.GetId()]
	if !ok {
		return nil, errors.New("record doesn't exists")
	}
	obj.db[person.GetId()] = person
	return obj.db[person.GetId()], nil
}
func (obj *PersonRepository) Delete(id int) (entities.PersonEntity, error) {
	per, ok := obj.db[id]
	if !ok {
		return nil, errors.New("record doesn't exists")
	}
	return per, nil
}

func (obj *PersonRepository) GetById(id int) (entities.PersonEntity, error) {
	per, ok := obj.db[id]
	if !ok {
		return nil, errors.New("record doesn't exists")
	}
	return per, nil
}

func (obj *PersonRepository) Search(query string) ([]entities.PersonEntity, error) {

	var persons []entities.PersonEntity

	for _, per := range obj.db {
		if per.GetEmail() == query || per.GetName() == query {
			persons = append(persons, per)
		}
	}

	return persons, nil
}
