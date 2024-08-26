package handlers

import (
	"coding_test/models"
	"coding_test/database"
	"github.com/gorilla/mux"
	"database/sql"
	"encoding/json"
    "net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	database.Connect()
    rows, err := database.DB.Query("SELECT username, password, active FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.Username, &user.Password, &user.Active); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	database.Connect()
    params := mux.Vars(r)
    username := params["username"] 

    var user models.User
    query := "SELECT username, password, active FROM users WHERE username = $1"
    err := database.DB.QueryRow(query, username).Scan(&user.Username, &user.Password, &user.Active)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(user)
}