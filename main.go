package main

import (
		"gopkg.in/mgo.v2/bson"

	"github.com/programmer-richa/mongodb/models"
)

// Entry point of the program

func main(){
	user := models.User{Name: "Richa", Email: "tc.ldh.richa@gmail.com", Password: "12345", SubscribeToEmail: true}
   // Create bson id / Unique user id
	user.Id = bson.NewObjectId()
	// Call Insert Method to add record to database
	user.Insert()
}