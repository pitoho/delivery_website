package main

import (
    "database/sql"
    "fmt"
	"errors"
    "net/http"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func registerUser(db *sql.DB, username, lastName, phoneNumber, email, hashedPassword string) error {
    var exists bool
    err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM User_Data WHERE user_name = $1)", username).Scan(&exists)
    if err != nil {
        return err
    }

    if exists {
        return errors.New("имя пользователя уже занято")
    }

    

    _, err = db.Exec("CALL add_user($1, $2, $3, $4, $5)", username, lastName, phoneNumber, email, hashedPassword)
    return err
}


func registrationHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if r.Method != "POST" {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
            return
        }


        username := r.FormValue("username")
        lastName := r.FormValue("lastname")
        phoneNumber := r.FormValue("phonenumber")
        email := r.FormValue("email")

        hashedPassword, err := hashPassword(r.FormValue("password"))
        if err != nil {
            http.Error(w, "Ошибка при хешировании пароля", http.StatusInternalServerError)
            fmt.Println(err)
            return
        }


        err = registerUser(db, username, lastName, phoneNumber, email, hashedPassword)
        if err != nil {
            http.Error(w, "Ошибка при регистрации пользователя", http.StatusInternalServerError)
            fmt.Println(err)
            return
        }

        fmt.Fprintf(w, "Пользователь успешно зарегистрирован")
    }
}