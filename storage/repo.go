package storage

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
)

var col *qmgo.Collection

type MongoRepo struct {
}

func init() {
	mongoUrl := os.Getenv("MONGO_URL")

	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoUrl})
	db := client.Database("people")
	coll := db.Collection("person") // 'person' colletion

	if err != nil {
		log.Panicln(err)
	}

	col = coll

	// todo - Close the client!
	//defer func() {
	//	if err = client.Close(ctx); err != nil {
	//		log.Panicln(err)
	//	}
	//}()

}

func (m MongoRepo) GetCollection() *qmgo.Collection {
	return col
}

func (m MongoRepo) AllPersons() []Person {

	ctx := context.Background()

	all := make([]Person, 0)
	//err := m.GetCollection() .Find(ctx, bson.M{"age": 36}).Sort("id").Limit(7).All(&batch)
	err := m.GetCollection().Find(ctx, bson.M{}).Sort("id").All(&all)

	if err != nil {
		log.Panicln(err)
	}

	for _, p := range all {
		all = append(all, p)
	}
	return all

}

func (m MongoRepo) FindByAge(limit int, age int) Person {
	ctx := context.Background()
	pers := Person{}
	err := m.GetCollection().Find(ctx, bson.M{"age": age}).One(pers)
	fmt.Println(pers)

	if err != nil {
		log.Panicln(err)
	}
	return pers
}
