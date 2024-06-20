package main

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