package userHandler

import (
	"fmt"
	"go_portofolio/model"
)

// pointerレシーバにしている理由としては、特定のアドレスのDBの値にアクセスするため
func (sess *Session) CheckSession() (valid bool, err error) {
	get_session := `SELECT id, uuid, email, user_id, created_at
	FROM sessions WHERE uuid = ?`

	fmt.Println("aaa")
	err = model.Db.QueryRow(get_session, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}

	return
}
