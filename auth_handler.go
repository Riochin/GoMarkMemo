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

func loginHandler(w http.ResponseWriter, r *http.Request) {
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

        // 🕳️ ① Supabaseのログイン用URLを作成せよ
        req, err := http.NewRequest("POST", supabaseURL + "/auth/v1/token?grant_type=password",
            bytes.NewBuffer(jsonPayload))

        if err != nil {
            http.Error(w, "Request creation failed", http.StatusInternalServerError)
            return
        }

        // 🕳️ ② 必要なヘッダーをセットせよ
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("apikey", anonKey)

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            http.Error(w, "Request failed", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // 🕳️ ③ レスポンスを読み取り、ログに出力せよ（access_tokenなどが入ってる）
        body, _ := io.ReadAll(resp.Body)
        fmt.Println("Supabase login response:", string(body))

        // ここにaccess_tokenをCookieに保存する処理を追加予定（まだやらなくてOK）

    } else {
		w.Header().Set("Content-Type", "text/html")
        // 🕳️ ④ login.html テンプレートを読み込んで表示せよ
        tmpl := template.Must(template.ParseFiles("templates/login.html"))
        tmpl.Execute(w, nil)
    }
}
