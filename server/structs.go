package main

type User struct {
    Username string
    Password string
}

type Dish struct {
    ID          int
    Name        string
    ImagePath   string
    Price       int
    TagsID      int
}

type LoginResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}