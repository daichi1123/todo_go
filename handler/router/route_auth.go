package router

import (
	"go_portofolio/handler/userHandler"
	"go_portofolio/model"
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 第3引数以降は、可変長引数だから複数入れても問題ない
		// signup.htmlファイルの出力だけを行う関数(GETメソッド)
		GenerateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		// ParseForm入力フォームの解析を行うメソッド
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		user := model.User{
			// postで送られてきたフォームの値を取得する name属性の値を引数に書く
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		// 値の取得はできている
		log.Println(user)

		// 値が渡せていない
		if err := userHandler.User.CreateUser(userHandler.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password}); err != nil {
			log.Println()
		}
		http.Redirect(w, r, "/", 302)
	}
}