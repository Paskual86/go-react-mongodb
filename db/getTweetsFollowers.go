package db

import (
	"context"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTweetsFollowers(userId string, page int) ([]models.TweetsFollowerResponse, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	col := CollectionRelation()

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userId": userId}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweets.date": -1}}) // Ordenamiento Descendente.
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)

	var result []models.TweetsFollowerResponse
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}
