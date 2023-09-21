package repository

import (
	"context"
	//"errors"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"log/syslog"
	"sync"
	"taskOne/models"
	//"taskOne/mongodatabase"
)

//var (
//	nextID  int
//	dataMux sync.Mutex
//)

var repository *Repository
var repoOnce sync.Once

type Repository struct {
	database *mongo.Database
}

func NewCategoryRepo(db *mongo.Database) *Repository {

	repoOnce.Do(func() {
		repository = &Repository{
			database: db,
		}
	})
	return repository
	//client := mongodatabase.Init()
	//var Collection *mongo.Collection
	//Collection = client.Database("mydb").Collection("entities")
	//return &Repo{
	//	collection: Collection,
	//}, nil
}

func (r *Repository) Create(ctx context.Context, entity *models.Entity) (err error) {

	_, err = r.database.Collection("entities").InsertOne(ctx, entity)
	if err != nil {
		fmt.Printf("create error %v", err.Error())
		return
	}
	return

	//dataMux.Lock()
	//defer dataMux.Unlock()
	//nextID++
	//entity.ID = nextID
	//fmt.Println(entity)
	//_, err := r.collection.InsertOne(context.Background(), entity)
	//if err != nil {
	//	fmt.Println("Disconnected", err)
	//	return
	//}

}

//
//func Read(id int) (models.Entity, error) {
//	dataMux.Lock()
//	defer dataMux.Unlock()
//	var entitywithid models.Entity
//	filter := bson.M{"id": id}
//	err := mongodatabase.Collection.FindOne(context.Background(), filter).Decode(&entitywithid)
//	if err != nil {
//		return models.Entity{}, errors.New("Entity not found")
//	}
//	return entitywithid, nil
//}
//
//func Update(id int, updatedEntity *models.Entity) error {
//	dataMux.Lock()
//	defer dataMux.Unlock()
//
//	updatedEntity.ID = id
//	filter := bson.M{"id": id}
//	update := bson.M{"$set": updatedEntity}
//	updateResult, err := mongodatabase.Collection.UpdateOne(context.Background(), filter, update)
//	if err != nil {
//		return errors.New("Error")
//	}
//	if updateResult.MatchedCount == 0 {
//		return errors.New("Entitu not found")
//	}
//	return nil
//}
//
//func Delete(id int) error {
//	dataMux.Lock()
//	defer dataMux.Unlock()
//	filter := bson.M{"id": id}
//	deleteResult, err := mongodatabase.Collection.DeleteOne(context.Background(), filter)
//	if err != nil {
//		return errors.New("Error")
//	}
//	if deleteResult.DeletedCount == 0 {
//		return errors.New("Entity not found")
//	}
//	return nil
//}
