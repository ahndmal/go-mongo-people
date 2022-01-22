package main

import (
	"fmt"
	"go-mongo-people/storage"
)

func main() {

	repo := storage.MongoRepo{}
	//for _, p := range repo.AllPersons() {
	//	fmt.Println(p)
	//}

	for _, p := range repo.FindByAge(5, 36) {
		fmt.Println(p)
	}

}
