package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(userId string, page int64, search string, searchType string) ([]*models.User, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := CollectionUser()

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	// La linea "$regex": `(?i)` + search, es una expresion regular. El codigo `(?i)` permite que ignore mayusculas y minusculas,
	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false, err
	}

	var include bool

	// I think that this part could be change to use defer sentce at moment finish the loop.
	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false, err
		}

		var r models.Relation
		r.UserId = userId
		r.UserRelationId = s.ID.Hex()
		include = false

		found, err := QueryRelation(r)

		// Some validations
		include = (searchType == constants.SEARCH_TYPE_NEW && !found) || (searchType == constants.SEARCH_TYPE_FOLLOW && found)

		if r.UserRelationId == userId {
			include = false
		}

		if include {
			s.Password = ""
			s.Biography = ""
			s.Site = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}

		err = cur.Err()
		if err != nil {
			fmt.Println(err.Error())
			return results, false, err
		}
	}

	cur.Close(ctx)
	return results, true, nil
}
