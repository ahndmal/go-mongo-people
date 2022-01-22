package main

import (
	"fmt"
	"go-mongo-people/storage"
	"go-mongo-people/web"
	"log"
	"net/http"
	"os"
)

func main() {

	httpPort := os.Getenv("PORT")
	repo := storage.MongoRepo{}

	webService := web.New(repo)

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpPort), webService.Router))

}
