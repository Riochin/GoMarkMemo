package markdown

import (
	"github.com/russross/blackfriday/v2"
)

// MarkdownToHTML は、markdown形式の文字列をHTML形式に変換します
func ToHTML(md string) string {
	return string(blackfriday.Run([]byte(md)))
}