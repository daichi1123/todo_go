package userHandler

import (
	"fmt"
	"go_portofolio/model"
	"log"
	"time"
)

type Session model.Session

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	// create_at は、DB設定でparseTimeが必要だった
	const createSession = `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values(?, ?, ?, ?)`

	result, err := model.Db.Exec(createSession, model.CreateUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)

	const select_session = `SELECT id, uuid, email, user_id, created_at
	FROM sessions WHERE user_id = ? and email = ?`

	err = model.Db.QueryRow(select_session, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
