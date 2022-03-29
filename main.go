package main

import (
	"fmt"
	"go_portofolio/handler"
	"go_portofolio/handler/router"
	"go_portofolio/handler/userHandler"
	"go_portofolio/model"
	"log"
	"net/http"

	// データベースに接続するパッケージ

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

// Configで使用したい値を入れる構造体を作成
type ConfigList struct {
	Port    string
	Db_info string
	LogFile string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").String(),
		LogFile: cfg.Section("web").Key("log_file").String(),
		Db_info: cfg.Section("db").Key("db_info").String(),
	}
	model.DB_init()
}

func main() {
	user, _ := userHandler.GetUserByEmail("sample@gmail.com")
	session, err := user.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}

	valid, _ := session.CheckSession()
	fmt.Println(valid)

	// これでエラー自体は消えた
	// model.User{}は使用できないなぜならCreateUserが参照しているのは、userHandlerで定義したUser構造体だから
	// Create
	// userHandler.User.CreateUser(userHandler.User{})
	// todoHandler.Todo.CreateTodo(todoHandler.Todo{})

	// Get
	// get_user, _ := userHandler.GetUser(2)
	// get_todo, _ := todoHandler.GetTodo(2)
	// todos, _ := todoHandler.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// あるユーザが所持しているtodosを表示する
	// user情報を取得する
	// get_user, _ := userHandler.GetUser(1)

	// Update
	// get_user.UpdateUser()
	// get_todo.UpdateTodo()

	// Delete
	// get_user.DeleteUser()
	// get_todo.DeleteTodo()

	// サーバ周りの作成
	mux := http.NewServeMux()

	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("テスト用")
	})

	mux.HandleFunc("/signup", router.Signup)
	// sessionがあればログイン状態にするようにしなければならない
	mux.HandleFunc("/login", router.Login)
	mux.HandleFunc("/authenticate", router.Authenticate)
	mux.HandleFunc("/todos", router.Index)

	mux.HandleFunc("/index", handler.Index)
	// mux.HandleFunc("/data", handler.DataDisplay)

	http.ListenAndServe(Config.Port, mux)
}
