package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github/madi-api/config"

	"github.com/dgrijalva/jwt-go"
)

var configs = config.NewConfig()
var signingKey = []byte(configs.JwtSecret)

func IsAuthorized(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Authorization") == "" {
			responseWithError(w, http.StatusForbidden, "Missing authorization header.")
			return
		}

		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		claims, err := validateJWT(token)
		if err != nil {
			responseWithError(w, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		// Add claims to request context
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func responseWithError(res http.ResponseWriter, statusCode int, message string) {
	responseWithJSON(res, statusCode, map[string]string{"error": message})
}

func responseWithJSON(res http.ResponseWriter, statusCode int, payload interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(payload)
}

func validateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GenerateJWT(payload map[string]interface{}) (string, error) {
	// set expiration time for token
	expireTime := time.Now().Add(time.Hour * 24)
	// create new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": payload,
		"exp":     expireTime.Unix(),
	})
	// sign token with secret
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
