package models

import "gopkg.in/mgo.v2/bson"

// Represents a user, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	FirstName        string        `bson:"first_name" json:"first_name"`
	LastName        string        `bson:"last_name" json:"last_name"`
	Job        string        `bson:"job" json:"job"`
	Gender        string        `bson:"gender" json:"gender"`
	ProfileImage  string        `bson:"profile_image" json:"profile_image"`
	LinkedIn        string        `bson:"linkedin" json:"linkedin"`
	Twitter        string        `bson:"twitter" json:"twitter"`
	Skype        string        `bson:"skype" json:"skype"`
	EmailID		 string        `bson:"emailId" json:"emailId"`
}
