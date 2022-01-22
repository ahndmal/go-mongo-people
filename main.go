package main

import (
	"fmt"
	"github.com/AndriiMaliuta/go-mongo-people/storage"
	"github.com/AndriiMaliuta/go-mongo-people/web"
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
