package main

import (
    "database/sql"
    "math/rand"
    "time"

    _ "github.com/lib/pq"
)


func randomDish(db *sql.DB) (*Dish, error) {
    rows, err := db.Query("SELECT id_dish, dish_name, dish_image_path, price, tags_id FROM Dish_Storage")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    dishes := make([]Dish, 0)
    for rows.Next() {
        var dish Dish
        err := rows.Scan(&dish.ID, &dish.Name, &dish.ImagePath, &dish.Price, &dish.TagsID)
        if err != nil {
            return nil, err
        }
        dishes = append(dishes, dish)
    }

    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(dishes))
    return &dishes[randomIndex], nil
}


