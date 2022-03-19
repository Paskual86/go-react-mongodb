package db

import (
	"context"
	"log"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(userId string, page int64) ([]*models.TweetResponse, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer es una instruccion de GO que permite ejecutar una funcion al final de la rutina.
	defer cancel()

	db := MongoCN.Database(constants.DATABASE_TWITTER)
	col := db.Collection(constants.COLLECTION_TWEETS)

	var result []*models.TweetResponse

	condition := bson.M{
		"userid": userId,
	}

	options := options.Find()
	options.SetLimit(constants.RECORDS_BY_PAGE)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * constants.RECORDS_BY_PAGE)

	cursor, error := col.Find(ctx, condition, options)

	if error != nil {
		log.Fatal(error.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var record models.TweetResponse
		err := cursor.Decode(&record)
		if err != nil {
			return result, false
		}
		result = append(result, &record)
	}
	return result, true
}
