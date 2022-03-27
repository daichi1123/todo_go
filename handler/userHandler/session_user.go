package userHandler

import (
	"go_portofolio/model"
	"log"
)

type user model.User

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	const select_user = `SELECT id, uuid, name, email, password
	from users where email = ?`
	err = model.Db.QueryRow(select_user, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
