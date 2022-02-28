package db

import (
	"context"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(tweet models.Tweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer es una instruccion de GO que permite ejecutar una funcion al final de la rutina.
	defer cancel()

	db := MongoCN.Database(constants.DATABASE_TWITTER)
	col := db.Collection(constants.COLLECTION_TWEETS)

	record := bson.M{
		"userid":  tweet.UserId,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := col.InsertOne(ctx, record)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), true, nil
}
