package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)


func checkUserCredentials(db *sql.DB, email string, password string) bool {
    rows, err := db.Query("SELECT user_email, user_password FROM User_Data WHERE user_email = $1", email)
    if err != nil {
        if err == sql.ErrNoRows {
            return false
        }
        panic(err)
    }
    defer rows.Close()

    if rows.Next() { 
        var userdb User
        err := rows.Scan(&userdb.Email, &userdb.Password)
        if err != nil {
            log.Fatal(err)
        }
        
        pass:= MD5(password)
        return pass == userdb.Password
    }

    return false 
}

func loginHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            http.ServeFile(w, r, "../web/dist/index.html")
        }else if r.Method != "POST" {
            http.Error(w, "Метод не поддерживается", http.StatusBadRequest)
            return
        }else{
            body, err := ioutil.ReadAll(r.Body)
            if err != nil {
                http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
                return
            }

            var user User 
            err = json.Unmarshal(body, &user)
            if err != nil {
                http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
                return
            }

            if checkUserCredentials(db, user.Email, user.Password) {
                w.Header().Set("Set-Cookie", "token="+ user.Email +"; Path=/; HttpOnly: true") 
                response := LoginResponse{Success: true, Message: "Успешная аутентификация"}
                json.NewEncoder(w).Encode(response)
            } else {
                response := LoginResponse{Success: false, Message: "Неверное имя пользователя или пароль"}
                json.NewEncoder(w).Encode(response)
            }
    
        }  
    }
}
