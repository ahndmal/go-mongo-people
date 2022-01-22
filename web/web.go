package web

import (
	"fmt"
	"go-mongo-people/storage"
	"net/http"
)

type Service struct {
	repo   storage.MongoRepo
	Router http.Handler
}

func InitServer() {
	port := ":8082"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//todo
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is UP on port ")

	http.ListenAndServe(port, nil)
}
