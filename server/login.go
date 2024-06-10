package main

import (
    "database/sql"
    "fmt"
    "net/http"

    _ "github.com/lib/pq" 
)


func checkUserCredentials(db *sql.DB, username, password string) bool {
    var dbUsername, dbPassword string


    query := "SELECT username, password FROM users WHERE username = $1"
    err := db.QueryRow(query, username).Scan(&dbUsername, &dbPassword)
    if err != nil {
        if err == sql.ErrNoRows {

            return false
        }

        panic(err)
    }


    return password == dbPassword
}

func loginHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if r.Method != "POST" {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
            return
        }

        username := r.FormValue("username")
        password := r.FormValue("password")

        if checkUserCredentials(db, username, password) {
            fmt.Fprintf(w, "Вход выполнен успешно")
        } else {
            http.Error(w, "Неверное имя пользователя или пароль", http.StatusUnauthorized)
        }
    }
}
