package main

import (
	"authgo/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/create", handlers.CreateUser)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
