package main

import (
	"html/template"
	"net/http"
	"strconv"
)
import "mymodule/internal"
// メモ一覧ページ
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートを読み込む
	tmpl, err := template.ParseFiles("templates/index.html")
	if (err != nil) {
		http.Error(w, "テンプレートの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	// メモデータを渡してレンダリング
	err = tmpl.Execute(w, memos)
	if (err != nil) {
		http.Error(w, "テンプレートの実行に失敗しました", http.StatusInternalServerError)
	}
}

// メモ作成ページ
func createHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodGet) {
		// メモ作成ページを表示
		tmpl, err := template.ParseFiles("templates/create.html")
		if (err != nil) {
			http.Error(w, "テンプレート読み込みエラー", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if (r.Method == http.MethodPost) {
		// フォームデータを取得
		title := r.FormValue("title")
		content := r.FormValue("content")
		if (title == "" || content == "") {
			http.Error(w, "タイトルまたは内容が空です", http.StatusBadRequest)
			return
		}

		// 新しいメモを追加
		newMemo := Memo{Title: title, Content: content}
		memos = append(memos, newMemo)

		// JSONファイルに保存
		saveMemos()

		// メモ一覧ページにリダイレクト
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "無効なリクエスト", http.StatusMethodNotAllowed)
	}
}

// メモ削除
func deleteMemoHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method != http.MethodPost) {
		http.Error(w, "無効なリクエストです", http.StatusMethodNotAllowed)
		return
	}

	// フォームデータを取得
	indexStr := r.FormValue("index")
	index, err := strconv.Atoi(indexStr)
	if (err != nil || index < 0 || index >= len(memos)) {
		http.Error(w, "無効なインデックス", http.StatusBadRequest)
		return
	}

	// メモを削除（スライスから削除）
	memos = append(memos[:index], memos[index+1:]...)

	// JSONファイルに保存
	saveMemos()

	// メモ一覧にリダイレクト
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// メモ詳細ページ
func viewHandler(w http.ResponseWriter, r *http.Request) {
	indexStr := r.URL.Query().Get("index")
	index, err := strconv.Atoi(indexStr)
	if (err != nil || index < 0 || index >= len(memos)) {
		http.Error(w, "無効なインデックス", http.StatusBadRequest)
		return
	}
	// MarkdownをHTMLに変換し、エスケープを防ぐ
	memo := memos[index]
	memo.ContentHTML = template.HTML(markdown.ToHTML(memo.Content))

	// テンプレートを読み込んで表示
	tmpl, err := template.ParseFiles("templates/view.html")
	if (err != nil) {
		http.Error(w, "テンプレート読み込みエラー", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, memo)
}