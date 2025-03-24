package main

import(
	"fmt"
	"net/http"
	"html/template"
	// "log"
)

type Memo struct {
	Title string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		memos := []Memo{
			{"Goでwebアプリを作った"},
			{"Markdown対応のwebアプリ"},
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "テンプレートの読み込みに失敗しました", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, memos)
		if err != nil {
			http.Error(w, "テンプレート実行エラー", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}