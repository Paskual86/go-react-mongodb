package db

import (
	"context"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UsersExists(email string) (models.User, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer es una instruccion de GO que permite ejecutar una funcion al final de la rutina.
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condicion).Decode(&result)

	ID := result.Id.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
