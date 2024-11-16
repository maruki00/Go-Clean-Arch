package contracts

import entities "person/internal/Person/Domain/Entities"

type IPersonRepository interface {
	// Create(person entities.PersonEntity) (entities.PersonEntity, error)
	// Update(person entities.PersonEntity) (entities.PersonEntity, error)
	// Delete(id int) (entities.PersonEntity, error)
	// GetById(id int) (entities.PersonEntity, error)
	// Search(query string) ([]entities.PersonEntity, error)

	Create(person entities.PersonEntity) (entities.PersonEntity, error)
	Update(person entities.PersonEntity) (entities.PersonEntity, error)
	Delete(id int) (entities.PersonEntity, error)
	GetById(id int) (entities.PersonEntity, error)
	Search(query string) ([]entities.PersonEntity, error)
}
