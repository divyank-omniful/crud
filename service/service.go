package service

import (
	"taskOne/entity"
	"taskOne/repository"
)

func CreateEntity(entity *entity.Entity) {
	repository.Create(entity)
}

func ReadEntity(id int) (entity.Entity, error) {
	return repository.Read(id)
}

func UpdateEntity(id int, updatedEntity *entity.Entity) error {
	return repository.Update(id, updatedEntity)
}

func DeleteEntity(id int) error {
	return repository.Delete(id)
}
