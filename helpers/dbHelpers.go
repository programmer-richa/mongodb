package helpers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"

)

const DBName = "CRUD"

// The collection name list

// A collection is used to store information in the mongodb.

// A collection is similar to tables,used in relational databases.

const (
	CollectionPrefix ="db."
	UserCollection         = "users"
)

var dbUser string = "root"

var dbPassword string = "local"

// GetSession creates a db session and returns its reference.
func GetSession() (*mgo.Session, error) {

	dbSession, err := mgo.Dial("mongodb://localhost")

	if err != nil {

		Logger(Panic, err)

		return nil, err

	}
	// Return successful connection
	return dbSession, nil

}

// InsertData inserts an entry in the specified collection name using the provided db session

// It returns true if the record is inserted successfully.
func InsertData(collectionName string, entry interface{}) error {
	// Create DB Session
	dbSession, err := GetSession()

	if err != nil {
		Logger(Panic, err)
		return err
	}


	// Close DB connection after this method is executed.

	defer dbSession.Close()


	database := mgo.Database{

		Session: dbSession,

		Name:    DBName,

	}


	//Initialize DB Collection here

	collection := mgo.Collection{

		Database: &database,

		Name:     collectionName,

		FullName: CollectionPrefix+collectionName,

	}

	err = dbSession.DB(DBName).C(collection.Name).Insert(entry)

	if err != nil {

		return err

	}

	return nil

}
