package database

import (
	"database/sql"
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

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
}

func Db_init() (db *sql.DB) {
	db, err := sql.Open("mysql", Config.Db_info)
	if err != nil {
		fmt.Print("データベース接続失敗")
		log.Fatal(err)
	} else {
		// 問題なく接続完了
		fmt.Print("データベース接続成功")
	}

	return db
	// err := db.Ping() 意味を調べる
}
