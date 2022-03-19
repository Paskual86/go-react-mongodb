package db

import (
	"context"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(tweetId string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer es una instruccion de GO que permite ejecutar una funcion al final de la rutina.
	defer cancel()

	db := MongoCN.Database(constants.DATABASE_TWITTER)
	col := db.Collection(constants.COLLECTION_TWEETS)

	objId, _ := primitive.ObjectIDFromHex(tweetId)

	condition := bson.M{
		"_id":    objId,
		"userid": userId,
	}

	_, err := col.DeleteOne(ctx, condition)

	return err
}
