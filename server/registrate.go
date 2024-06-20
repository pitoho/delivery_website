package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "io/ioutil"

    _ "github.com/lib/pq"
)
func registerHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET"{
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

			_, err = db.Exec("CALL add_user($1, $2, $3, $4, $5)", user.Username, user.Usersurname, user.Phonenum, user.Email, user.Password)
			if err != nil {
				http.Error(w, "Ошибка регистрации", http.StatusInternalServerError)
				log.Println("Ошибка регистрации:", err)
				return
			}
	
			response := LoginResponse{Success: true, Message: "Регистрация успешна"}
			json.NewEncoder(w).Encode(response)
		}
    }
}

