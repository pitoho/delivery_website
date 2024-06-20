package main

import (
    "fmt"
	"log"
    "net/http"
	"database/sql"
        "encoding/json"
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

    // fs := http.FileServer(http.Dir("../web/dist"))
	http.Handle("/", http.HandlerFunc(getDishes(db)))

	// http.Handle("/private", http.HandlerFunc(private()))
    // // http.Handle("/#procrast", http.HandlerFunc(getDishes(db)))
    // http.Handle("/public", http.HandlerFunc(getDishes(db)))
    // http.Handle("/login", http.HandlerFunc(loginHandler(db)))
    // http.Handle("/registrate", http.HandlerFunc(registerHandler(db)))

    
    
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
		cookie := &http.Cookie{
            Name:     "dishes",
            Value:    string(jsonDishes),
            Path:     "/",
            HttpOnly: true,
            MaxAge:   60 * 60 * 24, 
            Secure:   false,
        }
		http.SetCookie(w, cookie)
        w.Header().Set("Content-Type", "text/html")

        http.ServeFile(w, r, "C:/0_DELIVERY_WITH_BACKEND_4/delivery_website/web/index.html") 
    }
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