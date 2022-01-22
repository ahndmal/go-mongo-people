package storage

//import (
//	"github.com/globalsign/mgo"
//	"github.com/globalsign/mgo/bson"
//	"log"
//	"os"
//)
//
//// Repository defines the API repository implementation should follow.
//type Repository interface {
//	Find(id int) (*Person, error)
//	FindAll(selector map[string]interface{}) ([]*Person, error)
//	//Delete(Person *Person) error
//	//Update(Person *Person) error
//	//Create(Person ...*Person) error
//	//Count() (int, error)
//}
//
//const collName = "person"
//
//func GetCollectionName() string {
//	return collName
//}
//
//type MongoRepository struct {
//	logger  *log.Logger
//	session *mgo.Session
//}
//
//func (r MongoRepository) Find(id int) (*Person, error) {
//	//todo
//	session := r.session.Copy()
//	defer session.Close()
//	coll := session.DB("").C(collName)
//
//	var person Person
//	err := coll.Find(bson.M{"id": id}).One(&person)
//	if err != nil {
//		r.logger.Printf("error: %v\n", err)
//		return nil, err
//	}
//	return &person, nil
//}
//
//func newMongoSession() (*mgo.Session, error) {
//	mongoURL := os.Getenv("MONGO_URL")
//	if mongoURL == "" {
//		log.Fatal("MONGO_URL not provided")
//	}
//	return mgo.Dial(mongoURL)
//}
//
//func newMongoRepositoryLogger() *log.Logger {
//	return log.New(os.Stdout, "[mongoDB] ", 0)
//}
//
//func NewMongoRepository() MongoRepository {
//	logger := newMongoRepositoryLogger()
//	session, err := newMongoSession()
//	if err != nil {
//		logger.Fatalf("Could not connect to the database: %v\n", err)
//	}
//
//	return MongoRepository{
//		session: session,
//		logger:  logger,
//	}
//}
