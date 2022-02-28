package routers

import (
	"errors"
	"strings"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var AuthorizatedEmail string
var AuthorizatedUserId string

/*Process Token*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	miClave := []byte(constants.SECRET_KEY)
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Token format invalid")
	}

	token = strings.TrimSpace(splitToken[1])
	tkt, err := jwt.ParseWithClaims(token, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, found, _ := db.UsersExists(claims.Email)

		if found {
			AuthorizatedEmail = claims.Email
			AuthorizatedUserId = claims.ID.Hex()
		}
		return claims, found, AuthorizatedUserId, nil
	}

	if !tkt.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}

	return claims, false, string(""), err
}
