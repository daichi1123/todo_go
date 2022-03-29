package router

import (
	"net/http"
)

// ログインしていないユーザしか入れない
func top(w http.ResponseWriter, r *http.Request) {
	_, err := Session(w, r)
	if err != nil {
		GenerateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
	// クッキーに保存してクッキーをアクセスしてその値があるかを判定している
}

// ログインしたユーザしか入ることができない
func Index(w http.ResponseWriter, r *http.Request) {
	_, err := Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		GenerateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
