package main

import (
    "encoding/json"
    "fmt"
)

func parseDish(dishStr string) []Dish{
    var dishes []Dish

    err := json.Unmarshal([]byte(dishStr), &dishes)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
    }
	return dishes
}