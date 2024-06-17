package main

import (
    "database/sql"
    "net/http"
    "encoding/json"
    _ "github.com/lib/pq" 
)


func checkUserCredentials(db *sql.DB, username, password string) bool {

    var dbUsername, dbPassword string

    query := "SELECT user_name, user_password FROM User_Data WHERE user_name = $1"
    err := db.QueryRow(query, username).Scan(&dbUsername, &dbPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            return false
        }
        panic(err)
    }

    return password == dbPassword
}

func loginHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {

    // fs := http.FileServer(http.Dir("../web/dist"))
	// http.Handle("/", fs)

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
            return
        }

        decoder := json.NewDecoder(r.Body)
        var credentials struct {
            Username string `json:"user_name"`
            Password string `json:"user_password"`
        }
        err := decoder.Decode(&credentials)
        if err != nil {
            http.Error(w, "Неверный формат данных", http.StatusBadRequest)
            return
        }

        if checkUserCredentials(db, credentials.Username, credentials.Password) {
            response := LoginResponse{Success: true, Message: "Успешная аутентификация"}
            json.NewEncoder(w).Encode(response)
            http.Redirect(w, r, "/private", http.StatusFound)
            return
        } else {
            response := LoginResponse{Success: false, Message: "Неверное имя пользователя или пароль"}
            json.NewEncoder(w).Encode(response)
        }
    }
}
