package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=mother545 dbname=delivery_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		log.Fatalf("Error connecting to database: %v", err)
	}

	fs := http.FileServer(http.Dir("../web/dist"))
	http.Handle("/", fs)

	http.Handle("/private", http.HandlerFunc(getUserForPrivate(db)))
	http.Handle("/success", http.HandlerFunc(success(db)))
	http.Handle("/dishes", http.HandlerFunc(getDishes(db)))
	http.Handle("/public", http.HandlerFunc(getDishes(db)))
	http.Handle("/login", http.HandlerFunc(loginHandler(db)))
	http.Handle("/registrate", http.HandlerFunc(registerHandler(db)))
	http.Handle("/order", http.HandlerFunc(orderHandler(db)))

	fmt.Println("http://localhost:3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
