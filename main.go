package main

import (
	"fmt"
	"go-mongo-people/storage"
)

func main() {

	repo := storage.MongoRepo{}
	for _, p := range repo.AllPersons() {
		fmt.Println(p)
	}

}
