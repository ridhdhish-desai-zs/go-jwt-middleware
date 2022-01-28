package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ridhdhish-zopsmart/go-jwt-middleware/models"
)

// Just for fun
func SetHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		h.ServeHTTP(res, req)
	})
}

// To Validate the auth token
func CheckAuthToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")

		token := strings.Split(req.Header.Get("x-authorization-token"), " ")

		if len(token) <= 1 || token[1] == "" {
			_, _ = res.Write([]byte(`{"error": "Invalid Token. Could not parse"}`))
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("mysecret"), nil
		})

		if err != nil {
			_, _ = res.Write([]byte(`{"error": "Invalid Token. Could not parse"}`))
			return
		}

		email := fmt.Sprint(claims["email"])
		password := fmt.Sprint(claims["password"])

		data := models.User{
			Email:    email,
			Password: password,
		}

		ctx := context.WithValue(req.Context(), "data", data)
		h.ServeHTTP(res, req.WithContext(ctx))

	})
}
