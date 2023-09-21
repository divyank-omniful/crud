package service

import (
	"context"
	"fmt"
	"sync"
	"taskOne/models"
	"taskOne/repository"
)

var service Service
var serviceOnce sync.Once

type Service struct {
	repo repository.Repository
}

func NewService(repository repository.Repository) (service Service) {

	serviceOnce.Do(func() {
		service = Service{
			repo: repository,
		}
	})
	return service
	//categoryRepo, err := repository.NewCategoryRepo()
	//if err != nil {
	//	return Service{}, err
	//}
	//
	//return Service{
	//	repo: *categoryRepo,
	//}, nil
}

func (sv *Service) CreateEntity(ctx context.Context, entity *models.Entity) (entity2 *models.Entity, err error) {
	//sv.repo.Create(entity)
	err = sv.repo.Create(ctx, entity)
	if err != nil {
		fmt.Printf("create error %v", err.Error())
		return
	}
	entity2 = entity
	return
}

//
//func ReadEntity(id int) (models.Entity, error) {
//	return repository.Read(id)
//}
//
//func UpdateEntity(id int, updatedEntity *models.Entity) error {
//	return repository.Update(id, updatedEntity)
//}
//
//func DeleteEntity(id int) error {
//	return repository.Delete(id)
//}
