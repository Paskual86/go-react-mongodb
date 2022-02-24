package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Paskual86/go-react-mongodb.git/models"
)

/*GenerateJWT*/
func GenerateJWT(t models.User) (string, error) {

	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.Lastname,
		"birthdate": t.Birthdate,
		"location":  t.Location,
		"site":      t.Site,
		"_id":       t.Id.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
