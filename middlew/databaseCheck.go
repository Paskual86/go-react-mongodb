package middlew

import (
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/db"
)

/*Database Check: Middleware that allow us to know if the database connection is available*/
func DatabaseCheck(next http.HandlerFunc) http.HandlerFunc {
	// Devuelvo una funcion anonima
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Database Connection Lost", 500)
			return
		}
		// Le paso todo lo recibido al siguiente valor de la cadena
		next.ServeHTTP(w, r)
	}
}
