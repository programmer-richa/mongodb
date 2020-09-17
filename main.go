package main

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/programmer-richa/mongodb/models"
	"github.com/programmer-richa/mongodb/helpers"
)

// Entry point of the program

func main(){
	found,_:=models.IsExistingUser("abc@gmail.com")
	if !found {
		user := models.User{Name: "Richa", Email: "abc@gmail.com", Password: "12345", SubscribeToEmail: true}
		// Create bson id / Unique user id
		user.Id = bson.NewObjectId()
		// Call Insert Method to add record to database
		user.Insert()
		helpers.LogMessage("Record Added successfully")

	}else{
		helpers.LogMessage("Record already exists")
	}

	found,_=models.IsExistingUser("newabc@gmail.com")
	if !found {
		user := models.User{Name: "Richa", Email: "newabc@gmail.com", Password: "12345", SubscribeToEmail: true}
		// Create bson id / Unique user id
		user.Id = bson.NewObjectId()
		// Call Insert Method to add record to database
		user.Insert()
		helpers.LogMessage("Record Added successfully")
	}else{
		helpers.LogMessage("Record already exists")
	}
	// Fetch single user record
	user,_ := models.GetUserByEmail("newabc@gmail.com")
	helpers.LogMessage(user.Name,user.Id.Hex())



// Find user by id
id:="5ef7017ccd982c1dbc12c669"
record,_:=models.FindUser(id)
	record.Name="Riya"
	// Update record
	record.Update()

}

