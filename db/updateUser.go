package db

import (
	"context"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(u models.User, id string) (bool, error) {

	_, err := FindProfile(id)

	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer es una instruccion de GO que permite ejecutar una funcion al final de la rutina.
	defer cancel()

	db := MongoCN.Database(constants.DATABASE_TWITTER)
	col := db.Collection(constants.COLLECTION_USERS)

	record := make(map[string]interface{})

	if len(u.Name) > 0 {
		record["name"] = u.Name
	}

	if len(u.Lastname) > 0 {
		record["lastname"] = u.Lastname
	}

	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}

	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}

	if len(u.Location) > 0 {
		record["location"] = u.Location
	}

	if len(u.Biography) > 0 {
		record["biography"] = u.Biography
	}

	record["birthdate"] = u.Birthdate

	updtString := bson.M{
		"$set": record,
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, errUpd := col.UpdateOne(ctx, filter, updtString)

	if errUpd != nil {
		return false, errUpd
	}

	return true, nil
}
