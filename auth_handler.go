package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"bytes"
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	supabaseURL := os.Getenv("SUPABASE_PROJECT_URL")
	anonKey := os.Getenv("SUPABASE_ANON_KEY")

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		payload := map[string]string{
			"email":    email,
			"password": password,
		}

		jsonPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", supabaseURL+"/auth/v1/signup", bytes.NewBuffer(jsonPayload))
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("apikey", anonKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Request error", http.StatusInternalServerError)
			return
		}

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Supabase response status:", resp.Status)
		fmt.Println("Supabase response body:", string(body))

		defer resp.Body.Close()

		//jwtを受け取ってcookieに保存する処理を後でここに書く
	} else {
		//getだったらフォーム表示
		w.Header().Set("Content-Type", "text/html")
		tmpl := template.Must(template.ParseFiles("templates/signup.html"))
		tmpl.Execute(w, nil)
	}
}
