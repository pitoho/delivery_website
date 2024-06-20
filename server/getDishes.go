package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"

    _ "github.com/lib/pq"
)

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
		cookie := &http.Cookie{
            Name:     "dishes",
            Value:    string(jsonDishes),
            Path:     "/",
            HttpOnly: true,
            MaxAge:   60 * 60 * 24, 
            Secure:   false,
        }
		http.SetCookie(w, cookie)
        w.Header().Set("Content-Type", "application/json")

        http.ServeFile(w, r, "../web/dist/index.html") 
    }
}