package main

import (
    "github.com/dgrijalva/jwt-go"
)
type User struct {
    Username     string `json:"username"`
    Usersurname  string `json:"usersurname"`
    Phonenum     string `json:"phonenum"`
    Email        string `json:"email"`
    Password     string `json:"password"`
}

type DishWithTag struct {
    ID          int    `json:"id"`
    DishName    string `json:"dish_name"`
    ImagePath   string `json:"dish_image_path"`
    Price       int    `json:"price"`
    Tags        string `json:"tags_id"`
}
type OrderInfo struct {
    Dishes      string  `json:"orderedFood"`
    Street      string  `json:"street"`
    House       int     `json:"house"`
    Corpus      int     `json:"corpus_building"`
    Flat        int     `json:"flat"`
    TotalPrice  int     `json:"totalPrice"`
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
    IDDish          int    `json:"id_dish"`
    DishName        string `json:"dish_name"`
    DishImagePath   string `json:"dish_image_path"`
    Price           int    `json:"price"`
    TagsID          string `json:"tags_id"`
}