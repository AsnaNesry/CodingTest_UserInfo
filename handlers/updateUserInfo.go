package handlers

import (
    "coding_test/database"
    "coding_test/models"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    username := params["username"]

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    query := "UPDATE users SET password = $1, active = $2 WHERE username = $3"
    _, err := database.DB.Exec(query, user.Password, user.Active, username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    user.Username = username 
	response := map[string]string{
        "message":  "User updated successfully",
        "username": user.Username,
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}
