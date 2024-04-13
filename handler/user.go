package handler

import (
	"database/sql"
	"demo/model"
	"demo/queries"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		json.NewDecoder(r.Body).Decode(&user)
		defer r.Body.Close()

		err := queries.CreateUser(db, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created successfully"))
	}
}

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		user, err := queries.GetUserByUsername(db, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.Username != "" {
			json.NewEncoder(w).Encode(user)
		} else {
			users, err := queries.GetAllUsers(db)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(users)
		}
	}
}
