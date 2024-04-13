package main

import (
	"demo/dbConnection"
	"demo/handler"
	"demo/router"
	"net/http"
)

func main() {
	router := router.NewRouter()
	db := dbConnection.NewDB()

	defer db.Close()

	router.HandleFunc("/user", handler.CreateUserHandler(db)).Methods("POST")
	router.HandleFunc("/user/{username}", handler.GetUserHandler(db)).Methods("GET")

	http.ListenAndServe(":8000", router)
}
