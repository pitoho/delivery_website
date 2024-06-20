package main

import (
    "fmt"
	"log"
    "net/http"
	"database/sql"
    _ "github.com/lib/pq" 

)


func main() {
    connStr := "user=postgres password=mother545 dbname=delivery_db sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        fmt.Println(err)
    }

    defer db.Close()

    if err := db.Ping(); err != nil {
        fmt.Println(err)
    }

    // fs := http.FileServer(http.Dir("../web/dist"))
	http.Handle("/", http.HandlerFunc(getDishes(db)))

	http.Handle("/private", http.HandlerFunc(private()))
    // http.Handle("/#procrast", http.HandlerFunc(getDishes(db)))
    http.Handle("/public", http.HandlerFunc(getDishes(db)))
    http.Handle("/login", http.HandlerFunc(loginHandler(db)))
    http.Handle("/registrate", http.HandlerFunc(registerHandler(db)))
    
    fmt.Println("http://localhost:3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
