package core

import (
	"fmt"
	"github/madi-api/app/model"
	"github/madi-api/config"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var configs = config.NewConfig()
var signingKey = []byte(configs.JwtSecret)

func IsAuthorized(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil || len(r.Header["Authorization"]) == 0 {
			model.NewResponse(http.StatusForbidden, "Missing authorization header.", nil)
			return
		}

		token := r.Header.Get("Authorization")
		sToken := strings.Split(token, "Bearer ")
		token = sToken[1]

		accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error.")
			}
			return signingKey, nil
		})

		if err != nil {
			if err == jwt.ErrTokenExpired {
				model.NewResponse(http.StatusBadRequest, "Invalid token expired.", w)
				return
			} else {
				model.NewResponse(http.StatusBadRequest, "Invalid authorization token.", w)
				return
			}
		}

		if accessToken.Valid {
			next.ServeHTTP(w, r)
		}

	})
}

func GenerateJWT(payload interface{}) (string, error) {

	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Payload":   payload,
		"ExpiresAt": now.Add(1 * time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
