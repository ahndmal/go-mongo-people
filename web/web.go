package web

import (
	"encoding/json"
	"github.com/AndriiMaliuta/go-mongo-people/storage"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Service struct {
	repo   storage.MongoRepo
	Router http.Handler
}

func New(repo storage.MongoRepo) Service {
	service := Service{
		repo: repo,
	}

	router := httprouter.New()
	router.GET("/persons", service.Index)
	//router.POST("/persons", service.Create)
	//router.DELETE("/persons/:id", service.Delete)
	//router.PUT("/persons/:id", service.Update)

	service.Router = UseMiddlewares(router)

	return service
}

func (s Service) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	persBts, err := json.Marshal(s.repo.AllPersons())
	if err != nil {
		log.Panicln(err)
	}
	resp := []byte(persBts)
	w.Write(resp)
}

func UseMiddlewares(h http.Handler) http.Handler {
	//h = JSONApi(h)
	//h = OktaAuth(h)
	//h = Cors(h)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

//func InitServer() {
//	port := ":8082"
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//
//	})
//
//	fs := http.FileServer(http.Dir("static/"))
//	http.Handle("/static/", http.StripPrefix("/static/", fs))
//
//	fmt.Println("Server is UP on port ")
//
//	http.ListenAndServe(port, nil)
//}
