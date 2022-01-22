package main

import (
	"go-mongo-people/storage"
	"go-mongo-people/web"
	"log"
	"net/http"
	"os"
)

func main() {

	httpPort := os.Getenv("PORT")
	repo := storage.MongoRepo{}
	//for _, p := range repo.AllPersons() {
	//	fmt.Println(p)
	//}

	//for _, p := range repo.FindByAge(5, 36) {
	//	fmt.Println(p)
	//}
	//
	webService := web.New(repo)

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(":8088", webService.Router))

}
