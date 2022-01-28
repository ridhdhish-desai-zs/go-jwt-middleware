package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ridhdhish-zopsmart/go-jwt-middleware/http/users"
	authMiddleware "github.com/ridhdhish-zopsmart/go-jwt-middleware/middlewares/auth"
)

func main() {
	fmt.Println("Intro to middlewares with JWT")

	router := mux.NewRouter()

	// API Routes
	router.Path("/api/users").Methods("POST").HandlerFunc(users.CreateToken)
	router.Path("/api/users").Methods("GET").Handler(
		// TODO: Add middlewares here
		addMiddlewares(http.HandlerFunc(users.ValidateUser), authMiddleware.SetHeader),
	)

	fmt.Println("Listening to port 5000")
	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println("Server Error: ", err)
	}
}

// Registering all the middlewares specified for every URL path
func addMiddlewares(h http.Handler, middlewares ...func(handler http.Handler) http.Handler) http.Handler {
	fmt.Println("Inside add middleware")
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
