package db

import (
	"github.com/Paskual86/go-react-mongodb.git/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func RelationCollection() *mongo.Collection {
	return MongoDatabaseConnection.Collection(constants.COLLECTION_RELATIONS)
}
