package userHandler

import (
	"go_portofolio/model"
	"log"
)

func (u *User) UpdateUser() (err error) {
	// UPDATEのテスト用
	u.Name = "DAICHI"
	u.Email = "example@example.com"
	// UPDATEのテスト用

	const updateCmd = `UPDATE users SET name = ?, email = ? WHERE id = ?;`

	stmt, err := model.Db.Prepare(updateCmd)
	_, err = stmt.Exec(u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	defer model.Db.Close()

	return
}
