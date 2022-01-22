package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go-mongo-people/storage"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
)

func main() {

	mongoUrl := os.Getenv("MONGO_URL")

	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoUrl})
	db := client.Database("people")
	coll := db.Collection("person") // 'person' colletion
	client.Ping(10)

	fmt.Println(client)
	fmt.Println(coll)
	fmt.Println(mongoUrl)

	defer func() {
		if err = client.Close(ctx); err != nil {
			log.Panicln(err)
		}
	}()

	batch := []storage.Person{}
	coll.Find(ctx, bson.M{"age": 36}).Sort("age").Limit(7).All(&batch)

	fmt.Println(batch)

	pers := storage.Person{}
	err = coll.Find(ctx, bson.M{"name": "123"}).One(pers)
	fmt.Println(pers)

	// ============================== CORE Client
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))

	if err != nil {
		log.Panicln(err)
	}

}
