package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	http.Handle("/", fs)

	http.Handle("/private", http.HandlerFunc(getUserForPrivate(db)))
	http.Handle("/success", http.HandlerFunc(private()))
	http.Handle("/dishes", http.HandlerFunc(getDishes(db)))
	http.Handle("/public", http.HandlerFunc(getDishes(db)))
	http.Handle("/login", http.HandlerFunc(loginHandler(db)))
	http.Handle("/registrate", http.HandlerFunc(registerHandler(db)))
	http.Handle("/order", http.HandlerFunc(orderHandler(db)))

	fmt.Println("http://localhost:3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

type User struct {
	Username    string `json:"username"`
	Usersurname string `json:"usersurname"`
	Phonenum    string `json:"phonenum"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type DishWithTag struct {
	ID        int    `json:"id"`
	DishName  string `json:"dish_name"`
	ImagePath string `json:"dish_image_path"`
	Price     int    `json:"price"`
	Tags      string `json:"tags_id"`
}
type OrderInfo struct {
	Dishes     string `json:"orderedFood"`
	Street     string `json:"street"`
	House      int    `json:"house"`
	Corpus     int    `json:"corpus_building"`
	Flat       int    `json:"flat"`
	TotalPrice int    `json:"totalPrice"`
}
type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Dish struct {
	IDDish        int    `json:"id_dish"`
	DishName      string `json:"dish_name"`
	DishImagePath string `json:"dish_image_path"`
	Price         int    `json:"price"`
	TagsID        string `json:"tags_id"`
}

func getUserForPrivate(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		var userId int
		cookie, err := r.Cookie("token")
		if err != nil {
			fmt.Println("Куки не найдена")
		}
		token := cookie.Value
		err = db.QueryRow("SELECT user_id FROM Sessions WHERE token = $1", token).Scan(&userId)
		if err != nil {
			fmt.Println("Ошибка при получении user_id:", err)
		}

		err = db.QueryRow("SELECT user_name, last_name, phone_number, user_email FROM User_Data WHERE id_user = $1", userId).Scan(&user.Username, &user.Usersurname, &user.Phonenum, &user.Email)
		if err != nil {
			fmt.Println("Ошибка при получении user_id:", err)
		}

		w.Header().Set("Set-Cookie", "user="+user.Username+","+user.Usersurname+","+user.Phonenum+","+user.Email+"; Path=/; HttpOnly: true")

		if r.Method == "GET" {
			http.ServeFile(w, r, "../web/dist/index.html")
		}
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

		pass := MD5(password)
		return pass == userdb.Password
	}

	return false
}

func loginHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "../web/dist/index.html")
		} else if r.Method != "POST" {
			http.Error(w, "Метод не поддерживается", http.StatusBadRequest)
			return
		} else {
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
				w.Header().Set("Set-Cookie", "token="+MD5(user.Email)+"; Path=/; HttpOnly: true")

				token := MD5(user.Email)
				_, err = db.Exec("CALL check_token_and_create_session($1)", token)
				if err != nil {
					fmt.Println("Ошибка при вызове процедуры:", err)
				}

				response := LoginResponse{Success: true, Message: "Успешная аутентификация"}
				json.NewEncoder(w).Encode(response)
			} else {
				response := LoginResponse{Success: false, Message: "Неверное имя пользователя или пароль"}
				json.NewEncoder(w).Encode(response)
			}

		}
	}
}

func orderHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		var order OrderInfo
		var userId int
		var existingID int
		var dishes []Dish
		var newOrderID int

		if r.Method == "GET" {
			http.ServeFile(w, r, "../web/dist/index.html")
		} else if r.Method != "POST" {
			http.Error(w, "Метод не поддерживается", http.StatusBadRequest)
			return
		} else {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
				return
			}

			err = json.Unmarshal(body, &order)
			if err != nil {
				http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
				return
			}

			cookie, err := r.Cookie("token")
			if err != nil {
				fmt.Println("Куки не найдена")
			}
			token := cookie.Value

			err = db.QueryRow("SELECT user_id FROM Sessions WHERE token = $1", token).Scan(&userId)
			if err != nil {
				fmt.Println("Ошибка при получении user_id:", err)
			}

			err = db.QueryRow("CALL check_delivery_address_uniqueness($1, $2, $3, $4, $5)", order.Street, order.House, order.Corpus, order.Flat, &existingID).Scan(&existingID)
			if err != nil {
				panic(err)
			}

			err = db.QueryRow("CALL add_order($1, $2, $3, $4, $5)", userId, existingID, order.TotalPrice, "Ожидает оплату", &newOrderID).Scan(&newOrderID)
			if err != nil {
				panic(err)
			}

			dishes = parseDish(order.Dishes)

			for _, dish := range dishes {
				_, err = db.Exec("CALL add_order_dish($1, $2)", newOrderID, dish.IDDish)
				if err != nil {
					panic(err)
				}
			}
			response := LoginResponse{Success: true, Message: "Заказ успешно создан"}
			fmt.Println(response)
			json.NewEncoder(w).Encode(response)
		}
	}
}

func parseDish(dishStr string) []Dish {
	var dishes []Dish

	err := json.Unmarshal([]byte(dishStr), &dishes)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return dishes
}

func private() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "../web/dist/index.html")
		}
	}

}

func registerHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "../web/dist/index.html")
		} else if r.Method != "POST" {
			http.Error(w, "Метод не поддерживается", http.StatusBadRequest)
			return
		} else {
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
