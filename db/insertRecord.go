package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Insert Record: Inserta un registro en la base de datos*/
func InsertRecord(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer es una instruccion de GO que permite ejecutar una funcion al final de la rutina.
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	// Como obtener el ID
	// TODO: Es muy raro la sentencia de abajo. Es extraño como se llama.
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	fmt.Println(ObjID)

	return ObjID.String(), true, nil
}
