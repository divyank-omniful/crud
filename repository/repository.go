package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"taskOne/entity"
	"taskOne/mongodatabase"
)

var (
	data    []entity.Entity
	nextID  int
	dataMux sync.Mutex
)

func Create(entity *entity.Entity) {
	dataMux.Lock()
	defer dataMux.Unlock()
	nextID++
	entity.ID = nextID
	data = append(data, *entity)
	fmt.Println(entity)
	_, err := mongodatabase.Collection.InsertOne(context.Background(), entity)
	if err != nil {
		fmt.Println("Disconnected", err)
		return
	}
}

func Read(id int) (entity.Entity, error) {
	dataMux.Lock()
	defer dataMux.Unlock()
	var entitywithid entity.Entity
	filter := bson.M{"id": id}
	err := mongodatabase.Collection.FindOne(context.Background(), filter).Decode(&entitywithid)
	if err != nil {
		return entity.Entity{}, errors.New("Entity not found")
	}
	return entitywithid, nil
}

func Update(id int, updatedEntity *entity.Entity) error {
	dataMux.Lock()
	defer dataMux.Unlock()

	updatedEntity.ID = id
	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedEntity}
	updateResult, err := mongodatabase.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("Error")
	}
	if updateResult.MatchedCount == 0 {
		return errors.New("Entitu not found")
	}
	return nil
}

func Delete(id int) error {
	dataMux.Lock()
	defer dataMux.Unlock()
	filter := bson.M{"id": id}
	deleteResult, err := mongodatabase.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return errors.New("Error")
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("Entity not found")
	}
	return nil
}
