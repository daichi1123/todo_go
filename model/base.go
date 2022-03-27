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

func DB_init() {
	Db, err = sql.Open("mysql", database.Config.Db_info)
	if err != nil {
		log.Fatalln(err)
	}

	// stringは使えなかった
	const createUsersTable = (`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(100) NOT NULL UNIQUE,
		name VARCHAR(100),
		email VARCHAR(100),
		password VARCHAR(100),
		created_at DATETIME)`)

	_, err = Db.Exec(createUsersTable)
	if err != nil {
		log.Print(err)
	}

	// stringは使えなかった
	const createTodosTable = (`CREATE TABLE IF NOT EXISTS todos(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`)

	_, err = Db.Exec(createTodosTable)
	if err != nil {
		log.Print(err)
	}

	const createSessionTable = (`CREATE TABLE IF NOT EXISTS sessions(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(100) NOT NULL UNIQUE,
		email VARCHAR(100),
		user_id INTEGER,
		created_at DATETIME)`)

	Db.Exec(createSessionTable)
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
