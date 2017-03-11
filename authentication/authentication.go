package authentication

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alexendoo/sync/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/urfave/negroni"
)

func Authenticate(secret []byte) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token, err := request.ParseFromRequest(
			r,
			request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return secret, nil
			},
		)

		if token.Valid {
			next(w, r)
		} else {
			http.Error(w, "Authorization failed", http.StatusUnauthorized)
			log.Println(err)
		}
	}
}

func Register(secret []byte, store model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
