package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/programmer-richa/mongodb/helpers"
)

//User holds basic information regarding site user
//MongoDB represents JSON documents in binary-encoded format called BSON behind the scenes.
//BSON extends the JSON model to provide additional data types
//and to be efficient for encoding and decoding within different languages
type User struct {
	UserName         string        `json:"username" bson:"username"`
	FirstName        string        `json:"firstname" bson:"firstname"`
	LastName         string        `json:"lastname" bson:"lastname"`
	Name             string        `json:"name" bson:"name"`
	Email            string        `json:"email" bson:"email"`
	Password         string        `json:"password" bson:"password"`
	SubscribeToEmail bool          `json:"subscribe" bson:"subscribe"`
	Id               bson.ObjectId `json:"id" bson:"_id"` // Data type  bson.ObjectId is for mongodb bson format
}

// Insert adds user information to database
func(u User) Insert() error{
	//store user to mongodb
	err := helpers.InsertData(helpers.UserCollection,&u)

	// Return if error

	if err != nil {
		helpers.Logger(helpers.Error,err)
		return err
	}
	return nil
}

func (u User) Update() error {
	err := helpers.UpdateData(helpers.UserCollection,u.Id,&u)
	// Return if error
	if err != nil {
		helpers.Logger(helpers.Error,err)
		return err
	}
	return nil
}

func  FindUser(id string) (User, error) {
	u:=User{}
	data,err:=helpers.FindByID(helpers.UserCollection,id)
	// convert bson.M to User
	bsonBytes, _ := bson.Marshal(data)
	bson.Unmarshal(bsonBytes, &u)
	return  u,err
}


// IsExistingUser examines if an user account is already registered
// with the given email id.
func IsExistingUser(email string) (found bool,err error) {
	found,err=helpers.IsExistingRecord(helpers.UserCollection,bson.M{"email":email})
	return found,err
}

// GetUserByEmail returns a user account information if is already registered
// with the given email id.
func GetUserByEmail(email string) ( User, error) {
	u:=User{}
	data,err:=helpers.GetaRecord(helpers.UserCollection,bson.M{"email":email})
	// convert bson.M to User
	bsonBytes, _ := bson.Marshal(data)
	bson.Unmarshal(bsonBytes, &u)
	return u,err
}



