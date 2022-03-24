package main

import (
	"fmt"
	"go_portofolio/handler"
	"go_portofolio/model"
	"net/http"

	// データベースに接続するパッケージ

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

// Configで使用したい値を入れる構造体を作成
type ConfigList struct {
	Port    string
	Db_info string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").String(),
		Db_info: cfg.Section("db").Key("db_info").String(),
	}
	model.DB_init()
}

func main() {
	// これでエラー自体は消えた
	// model.User{}は使用できないなぜならCreateUserが参照しているのは、userHandlerで定義したUser構造体だから
	// Create
	// userHandler.User.CreateUser(userHandler.User{})
	// todoHandler.Todo.CreateTodo(todoHandler.Todo{})

	// Get
	// get_user, _ := userHandler.GetUser(2)
	get_todo, _ := todoHandler.GetTodo(2)

	// Update
	// get_user.UpdateUser()
	// get_todo.UpdateTodo()

	// Delete
	// get_user.DeleteUser()
	get_todo.DeleteTodo()

	// サーバ周りの作成
	mux := http.NewServeMux()

	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("テスト用")
	})
	mux.HandleFunc("/index", handler.Index)
	// mux.HandleFunc("/data", handler.DataDisplay)

	http.ListenAndServe(Config.Port, mux)
}
