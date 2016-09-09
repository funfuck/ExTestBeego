package models

import (
	"gopkg.in/mgo.v2/bson"
	"testBeego/databases/mongo"
	"log"
)

type Person struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	FirstName string
	LastName  string
	Email     string
	Password  string 
}

func AddPerson(m *Person) (id int64, err error) {

	return
}

func GetPersonById(query bson.M) *Person {
	session := mongo.Conn()
	defer session.Close()

	var results = Person{}

	c := session.DB("test").C("person")
	err := c.Find(query).One(&results)
	if err != nil {
		log.Println(err)
	}

	return &results
}

func GetAllPerson(query bson.M, sortBy string) []Person {
	session := mongo.Conn()
	defer session.Close()

	var results []Person

	c := session.DB("test").C("person")
	err := c.Find(nil).Sort(sortBy).All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

func UpdatePersonById(m *Person) (err error) {
	return
}

func DeletePerson(id bson.ObjectId) (err error) {
	return
}
