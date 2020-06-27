package main

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/programmer-richa/mongodb/models"
	"github.com/programmer-richa/mongodb/helpers"
)

// Entry point of the program

func main(){
	found,_:=models.IsExistingUser("tc.ldh.richa@gmail.com")
	if !found {
		user := models.User{Name: "Richa", Email: "tc.ldh.richa@gmail.com", Password: "12345", SubscribeToEmail: true}
		// Create bson id / Unique user id
		user.Id = bson.NewObjectId()
		// Call Insert Method to add record to database
		user.Insert()
		helpers.LogMessage("Record Added successfully")

	}else{
		helpers.LogMessage("Record already exists")
	}

	found,_=models.IsExistingUser("richa@gmail.com")
	if !found {
		user := models.User{Name: "Richa", Email: "richa@gmail.com", Password: "12345", SubscribeToEmail: true}
		// Create bson id / Unique user id
		user.Id = bson.NewObjectId()
		// Call Insert Method to add record to database
		user.Insert()
		helpers.LogMessage("Record Added successfully")
	}else{
		helpers.LogMessage("Record already exists")
	}

}

