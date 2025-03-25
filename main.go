package main

import(
	"fmt"
	"net/http"
	"html/template"

	"log"
	"encoding/json"
	"os"
)

type Memo struct {
	Title string `json:"title"`
}

var memos []Memo

//メモ一覧ページ
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートを読み込む
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "テンプレートの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	// メモデータを渡してレンダリング
	err = tmpl.Execute(w, memos)
	if err != nil {
		http.Error(w, "テンプレートの実行に失敗しました", http.StatusInternalServerError)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// メモ作成ページを表示
		tmpl, err := template.ParseFiles("templates/create.html")
		if err != nil {
			http.Error(w, "テンプレート読み込みエラー", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// フォームデータを取得
		title := r.FormValue("title")
		if title == "" {
			http.Error(w, "タイトルが空です", http.StatusBadRequest)
			return
		}

		// 新しいメモを追加
		newMemo := Memo{Title: title}
		memos = append(memos, newMemo)

		// JSONファイルに保存
		saveMemos()

		// メモ一覧ページにリダイレクト
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "無効なリクエスト", http.StatusMethodNotAllowed)
	}
}

// メモをJSONファイルに保存
func saveMemos() {
	file, err := os.Create("data/memos.json")
	if err != nil {
		log.Println("ファイルの作成に失敗しました", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(memos)
}


func main() {
	// メモ一覧をロード
	memos = []Memo{}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}