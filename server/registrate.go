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
			var user User
			var exists bool
			var response LoginResponse

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
				return
			}
	
			err = json.Unmarshal(body, &user)
			if err != nil {
				http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
				return
			}

			err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM User_Data WHERE user_email = $1)", user.Email).Scan(&exists)
    		if err != nil {
        		response = LoginResponse{Success: false, Message: "Произошла ошибка запроса данных"}
				json.NewEncoder(w).Encode(response)
        		return 
    		}

    		if exists {
        		response = LoginResponse{Success: false, Message: "Пользователь с таким адресом электронной почты уже есть"}
				json.NewEncoder(w).Encode(response)
        		return
    		}

			hashedPassword := MD5(user.Password)

			_, err = db.Exec("CALL add_user($1, $2, $3, $4, $5)", user.Username, user.Usersurname, user.Phonenum, user.Email, hashedPassword)
			if err != nil {
				http.Error(w, "Ошибка регистрации", http.StatusInternalServerError)
				log.Println("Ошибка регистрации:", err)
				return
			}
	
			response = LoginResponse{Success: true, Message: "Регистрация успешна"}
			json.NewEncoder(w).Encode(response)
		}
    }
}

