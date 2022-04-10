package db

import (
	"context"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/models"
)

func DeleteRelation(value models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database(constants.DATABASE_TWITTER)
	col := db.Collection(constants.COLLECTION_RELATIONS)

	_, err := col.DeleteOne(ctx, value)

	return err == nil, err
}
