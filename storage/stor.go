package storage

import "log"

// Repository defines the API repository implementation should follow.
type Repository interface {
	Find(id string) (*Person, error)
	FindAll(selector map[string]interface{}) ([]*Person, error)
	Delete(Person *Person) error
	Update(Person *Person) error
	Create(Person ...*Person) error
	Count() (int, error)
}

const collection = "person"

func GetCollectionName() string {
	return collection
}

type MongoRepository struct {
	logger  *log.Logger
	session *mgo.Session
}
