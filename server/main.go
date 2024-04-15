package main

import (
    "fmt"
	"log"
    "net/http"
)


func main() {
	fs := http.FileServer(http.Dir("../web/dist"))
	http.Handle("/", fs)


	fmt.Println("http://localhost:3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
