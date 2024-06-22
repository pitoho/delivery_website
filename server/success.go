package main

import (
	"net/http"
)
func success()func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" {
			http.ServeFile(w, r, "../web/dist/index.html")
		}
	}
	
}
