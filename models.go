package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
)

type Memo struct {
	Title       string        `json:"title"`
	Content     string        `json:"content"`
	ContentHTML template.HTML `json:"-"`
}

var memos []Memo

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

// メモをJSONファイルからロード
func loadMemos() {
	file, err := os.Open("data/memos.json")
	if err != nil {
		log.Println("ファイルの読み込みに失敗しました", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&memos)
	if err != nil {
		log.Println("JSONデコードに失敗しました", err)
	}
}