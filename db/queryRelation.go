package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

func QueryRelation(value models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	col := CollectionRelation()

	condition := bson.M{
		"userid":         value.UserId,
		"userrelationid": value.UserRelationId,
	}

	var result models.Relation

	fmt.Println(result)

	err := col.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
