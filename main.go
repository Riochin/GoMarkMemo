package main

import (
	"fmt"
	"log"
	"net/http"
	// "mymodule/internal"
)

func main() {
	// メモ一覧をロード
	loadMemos()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteMemoHandler)
	http.HandleFunc("/view", viewHandler)

	http.HandleFunc("/signup", signupHandler)
	// http.HandleFunc("/login", loginHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
