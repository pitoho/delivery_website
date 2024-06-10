package main

import (
    "fmt"
	"log"
    "net/http"
	"database/sql"
    _ "github.com/lib/pq" 
)


func main() {
	fs := http.FileServer(http.Dir("../web/dist"))
	http.Handle("/", fs)


	fmt.Println("http://localhost:3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

	connStr := "user=postgres dbname=delivery_db password=babymonster sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        panic(err)
    }

    http.HandleFunc("/login", loginHandler(db))
	http.HandleFunc("/register", registrationHandler(db))

    fmt.Println("Сервер запущен на http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
