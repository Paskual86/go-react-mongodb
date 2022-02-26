package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfile(id string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer es una instruccion de GO que permite ejecutar una funcion al final de la rutina.
	defer cancel()

	db := MongoCN.Database(constants.DATABASE_TWITTER)
	col := db.Collection(constants.COLLECTION_USERS)

	var profile models.User
	// Convertimos el id recibido como parametro, a un ObjectId enn mongo
	objID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	// Clean the password
	profile.Password = ""

	if err != nil {
		fmt.Println("Record not found" + err.Error())
		return profile, err
	}
	return profile, nil
}
