package handlers

import (
	"coding_test/models"
	"coding_test/database"
	"encoding/json"
    "net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	database.Connect()
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    query := `INSERT INTO users (username, password, active) VALUES ($1, $2, $3)`
    _, err := database.DB.Exec(query, user.Username, user.Password, user.Active)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]string{
        "message":  "User created successfully",
        "username": user.Username,
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}
