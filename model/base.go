package model

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go_portofolio/database"
	"log"

	"github.com/google/uuid"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open("mysql", database.Config.Db_info)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}

// uuidの作成
func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return
}

// hash値の作成
func Encypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprint("%x", sha1.Sum([]byte(plaintext)))
	return
}
