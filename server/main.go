package main

import (
    "fmt"
	"log"
    "net/http"
	"database/sql"
    "crypto/md5"
    "io/ioutil"
    "encoding/json"
    "github.com/dgrijalva/jwt-go"
    _ "github.com/lib/pq" 

)


func main() {
    connStr := "host=localhost port=5432 user=postgres password=Vk691109 dbname=delivery_db sslmode=disable"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        fmt.Println(err)
    }

    defer db.Close()

    if err := db.Ping(); err != nil {
        fmt.Println(err)
        log.Fatalf("Error connecting to database: %v", err)
    }

    fs := http.FileServer(http.Dir("../web/dist"))
	http.Handle("/", fs )

	http.Handle("/private", http.HandlerFunc(private()))
    http.Handle("/dishes", http.HandlerFunc(getDishes(db)))
    http.Handle("/public", http.HandlerFunc(getDishes(db)))
    http.Handle("/login", http.HandlerFunc(loginHandler(db)))
    http.Handle("/registrate", http.HandlerFunc(registerHandler(db)))  
    http.Handle("/order", http.HandlerFunc(order()))
    
    fmt.Println("http://localhost:3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
func getDishes(db *sql.DB) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        rows, err := db.Query("SELECT * FROM get_dish_with_tags()")
        if err != nil {
            http.Error(w, fmt.Sprintf("Ошибка получения данных: %v", err), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var dishes []DishWithTag
        for rows.Next() {
            var dish DishWithTag
            if err := rows.Scan(&dish.ID, &dish.DishName, &dish.ImagePath, &dish.Price, &dish.Tags); err != nil {
                http.Error(w, fmt.Sprintf("Ошибка сканирования данных: %v", err), http.StatusInternalServerError)
                return
            }
            dishes = append(dishes, dish)
        }

        jsonDishes, err := json.Marshal(dishes)
        if err != nil {
            http.Error(w, fmt.Sprintf("Ошибка преобразования в JSON: %v", err), http.StatusInternalServerError)
            return
        }
		w.Write(jsonDishes)
        w.Header().Set("Content-Type", "application/json")
    }
}
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

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

func private()func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" {
			http.ServeFile(w, r, "../web/dist/index.html")
		}
	}
	
}

func order()func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" {
			http.ServeFile(w, r, "../web/dist/index.html")
		}
	}
	
}

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

type User struct {
    Username     string `json:"username"`
    Usersurname  string `json:"usersurname"`
    Phonenum     string `json:"phonenum"`
    Email        string `json:"email"`
    Password     string `json:"password"`
}

type Dish struct {
    ID          int    `json:"id"`
    DishName    string `json:"dish_name"`
    ImagePath   string `json:"dish_image_path"`
    Price       int    `json:"price"`
    TagsID      int    `json:"tags_id"`
}

type DishWithTag struct {
    ID          int    `json:"id"`
    DishName    string `json:"dish_name"`
    ImagePath   string `json:"dish_image_path"`
    Price       int    `json:"price"`
    Tags        string `json:"tags_id"`
}

type LoginResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}
type MyCustomClaims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}
