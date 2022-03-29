package router

import (
	"fmt"
	"go_portofolio/handler/userHandler"
	"go_portofolio/model"
	"html/template"
	"net/http"
)

type session userHandler.Session

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// この関数の解読はすごく勉強になる返り値に注目する
// func Session(w http.ResponseWriter, r *http.Request) (sess userHandler.Session, err error) {
// 	// ここがnilになる
// 	fmt.Println(sess)
// 	cookie, err := r.Cookie("_cookie")
// 	if err != nil {
// 		sess = userHandler.Session{UUID: cookie.Value}
// 		if ok, _ := sess.CheckSession(); !ok {
// 			err = fmt.Errorf("Invalid session")
// 		}
// 	}
// 	return

func Session(w http.ResponseWriter, r *http.Request) (sess model.Session, err error) {
	// ここがnilになる
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		sess = model.Session{UUID: cookie.Value}
		fmt.Println(sess)
		var handler_session = userHandler.Session{UUID: cookie.Value}
		if ok, _ := handler_session.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}
