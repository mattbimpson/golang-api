package main

import (
	"fmt"
	"net/http"
)

func loggerMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received : %v\n", r)
		nextHandler.ServeHTTP(w, r)
		fmt.Println("Request handled successfully")
	})
}
