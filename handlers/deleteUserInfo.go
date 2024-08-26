package handlers

import (
    "coding_test/database"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
   
    params := mux.Vars(r)
    username := params["username"] 
    query := "DELETE FROM users WHERE username = $1"
    result, err := database.DB.Exec(query, username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if rowsAffected == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
}
