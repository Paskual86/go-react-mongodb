package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*El ID tiene que declararse en Mayuscula. Es algo de la convension. Si lo cambias, por Id o iD te lo crea pero no lo podes recurperar ni acceder al id interno.*/
/*User Struct*/
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson: "name" json:"name,omitempty"`
	Lastname  string             `bson: "lastname" json:"lastname,omitempty"`
	Birthdate time.Time          `bson: "birthdate" json:"birthdate,omitempty"`
	Email     string             `bson: "email" json:"email"`
	Password  string             `bson: "password" json:"password, omitempty"`
	Avatar    string             `bson: "avatar" json:"avatar, omitempty"`
	Banner    string             `bson: "banner" json:"banner, omitempty"`
	Location  string             `bson: "location" json:"location, omitempty"`
	Site      string             `bson: "site" json:"site, omitempty"`
	Biography string             `bson: "biography" json:"biography, omitempty"`
}
