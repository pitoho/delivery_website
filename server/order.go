package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)


func orderHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {

    return func(w http.ResponseWriter, r *http.Request) {

		var order OrderInfo
		var userId int 
		var existingID int
		var dishes []Dish
		var newOrderID int

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

			for  _, dish := range dishes {
				_, err = db.Exec("CALL add_order_dish($1, $2)", newOrderID, dish.IDDish)
    			if err != nil { 
					panic(err)
    			}else{
					response := LoginResponse{Success: true, Message: "Заказ успешно создан"}
					json.NewEncoder(w).Encode(response)
				}
			}

    
        }  
    }
}