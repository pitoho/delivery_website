package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"net/http"
)
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

		rows, err := db.Query("SELECT * FROM get_user_orders($1)", userId)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var orders []Order
        for rows.Next() {
            var order Order
            err = rows.Scan(&order.IDOrder, &order.OrderTime, &order.TotalPrice, &order.OrderStatus)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            orders = append(orders, order)
        }
        if err := rows.Err(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        ordersJSON, err := json.Marshal(orders)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		// fmt.Println(ordersJSON)

		// w.Header().Set("Set-Cookie", "user=" + user.Username + "," + user.Usersurname + "," + user.Phonenum + "," + user.Email + "; Path=/; HttpOnly: true")
        // w.Header().Set("Set-Cookie", "orders="+ string(ordersJSON) +"; Path=/; HttpOnly: true", )

		http.SetCookie(w, &http.Cookie{
			Name:  "user",
			Value: user.Username + "," + user.Usersurname + "," + user.Phonenum + "," + user.Email,
			Path: "/", 
		})
		encodedOrdersJSON := url.QueryEscape(string(ordersJSON))
		http.SetCookie(w, &http.Cookie{
			Name:  "orders",
			Value: encodedOrdersJSON,
			Path: "/", 
		})
	

        if r.Method == "GET" {
            http.ServeFile(w, r, "../web/dist/index.html")
        }
	}

}
