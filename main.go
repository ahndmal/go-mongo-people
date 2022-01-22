package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

func main() {

	var err error
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Panicln("MONGO_URL not provided")
	}
	fmt.Println(mongoURL)
	if err != nil {
		log.Panicln(err)
	}
	//session, err2 := mgo.Dial(mongoURL)
	//fmt.Println(session)
	//
	//if err2 != nil {
	//	log.Panicln(err2)
	//}
	//
	//var pers storage.Person
	//c := session.DB("people").C("person")
	//err3 := c.Find("").One(&pers)
	//
	//if err3 != nil {
	//	log.Fatal(err3)
	//}
	//fmt.Println(pers)

	//info := &mgo.CollectionInfo{}
	//err = session.DB("").C("kudos").Create(info)

	//count, err := session.DB("people").C("person").Count()
	////person, err := serv.Service{}.GetPersonById(2606)
	//if err != nil {
	//	log.Panicln("ERROR !!!")
	//	log.Panicln(err)
	//}
	//log.Println(count)
	//defer session.Close()

	// ==============================
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))

	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Panicln(err)
		}
	}()

	// Ping the primary
	if err2 := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Panicln(err2)
	}
	fmt.Println("Successfully connected and pinged.")
}
