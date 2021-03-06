package helpers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	//err = dbSession.DB(DBName).DropDatabase()
	// Return successful connection
	return dbSession, nil

}

// GetDatabase returns database reference
func GetDatabase(dbSession *mgo.Session) mgo.Database{
	return	mgo.Database{

		Session: dbSession,

		Name:    DBName,

	}

}

func GetCollection(dbSession *mgo.Session,collectionName string) *mgo.Collection{
	return  dbSession.DB(DBName).C(collectionName)
}

// InsertData inserts an entry in the specified collection name using the provided db session
func InsertData(collectionName string, entry interface{}) error {
	// Create DB Session
	dbSession, err := GetSession()
	if err != nil {
		Logger(Panic, err)
		return err
	}
	// Close DB connection after this method is executed.
	defer dbSession.Close()
	//Initialize DB Collection here
	collection := GetCollection(dbSession,collectionName)
	err = collection.Insert(entry)
	if err != nil {

		return err

	}
	return nil
}

// UpdateData updates an entry in the specified collection name using the provided db session
func UpdateData(collectionName string,id ,entry interface{}) error {
	// Create DB Session
	dbSession, err := GetSession()
	if err != nil {
		Logger(Panic, err)
		return err
	}
	// Close DB connection after this method is executed.
	defer dbSession.Close()
	//Initialize DB Collection here
	collection := GetCollection(dbSession,collectionName)
	err = collection.UpdateId(id,entry)
	if err != nil {

		return err

	}
	return nil
}

// IsExistingRecord tests the existance of a record with the provided conditions.
// It returns true if the record is available in the collection.
func IsExistingRecord(collectionName string, condition bson.M ) ( found bool,err error){
	// Create DB Session
	dbSession, err := GetSession()
	if err != nil {
		Logger(Panic, err)
		return found,err
	}
	// Close DB connection after this method is executed.
	defer dbSession.Close()
	//Initialize DB Collection here
	collection := GetCollection(dbSession,collectionName)
	count,err := collection.Find(condition).Count()
	found = count>0
	return  found,err
}

// GetaRecord returns a record that satisfies the provided conditions.
// It returns nil if the record is unavailable in the collection.
func GetaRecord(collectionName string, condition bson.M) ( record interface{},err error){
	// Create DB Session
	dbSession, err := GetSession()
	if err != nil {
		Logger(Panic, err)
		return nil,err
	}
	// Close DB connection after this method is executed.
	defer dbSession.Close()
	//Initialize DB Collection here
	collection := GetCollection(dbSession,collectionName)
	err = collection.Find(condition).One(&record)
	return  record,err
}

// FindByID returns a record that satisfies the provided conditions.
// It returns nil if the record is unavailable in the collection.
func FindByID(collectionName string,id string ) ( record interface{},err error){
	// Create DB Session
	dbSession, err := GetSession()
	if err != nil {
		Logger(Panic, err)
		return nil,err
	}
	// Close DB connection after this method is executed.
	defer dbSession.Close()
	//Initialize DB Collection here
	collection := GetCollection(dbSession,collectionName)
	err = collection.FindId(bson.ObjectIdHex(id)).One(&record)
	return  record,err
}