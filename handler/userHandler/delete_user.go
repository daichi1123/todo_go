package userHandler

import (
	"go_portofolio/model"
	"log"
)

func (u *User) DeleteUser() (err error) {
	// GetUserで手に入れたUser情報で削除していく

	const deleteCmd = `DELETE FROM users WHERE id = ?;`
	stmt, _ := model.Db.Prepare(deleteCmd)
	_, err = stmt.Exec(u.ID)
	if err != nil {
		log.Fatal(err)
	}
	model.Db.Close()

	return
}
