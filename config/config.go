package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port    string
	LogFile string
	Db_info string
}

var Config ConfigList

func GetConfigVal() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:    cfg.Section("web").Key("port").String(),
		LogFile: cfg.Section("web").Key("log_file").String(),
		Db_info: cfg.Section("db").Key("db_info").String(),
	}
	fmt.Print(Config)
}
