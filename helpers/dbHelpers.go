package helpers

import (
	"gopkg.in/mgo.v2"
	"log"
)

const DBName = "CRUD"

// The collection name list

// A collection is used to store information in the mongodb.

// A collection is similar to tables,used in relational databases.

const (
	UserCollection         = "users"
	UserCollectionFullName = "db.users"
)

var dbUser string = "root"

var dbPassword string = "local"

// GetSession creates a db session and returns its reference.
func GetSession() (*mgo.Session, error) {

	dbSession, err := mgo.Dial("mongodb://localhost")

	if err != nil {

		log.Print(err)

		return nil, err

	}
	// Return successful connection
	return dbSession, nil

}

// InsertData inserts an entry in the specified collection name using the provided db session

// It returns true if the record is inserted successfully.
func InsertData(dbSession *mgo.Session, collection mgo.Collection, entry interface{}) error {

	err := dbSession.DB(DBName).C(collection.Name).Insert(entry)

	if err != nil {

		return err

	}

	return nil

}
