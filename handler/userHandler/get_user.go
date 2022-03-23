package userHandler

import (
	"go_portofolio/model"
	"log"
)

func GetUser(id int) (user User, err error) {
	user = User{}
	const select_user = `select id, uuid, name, email, password
	from users where id = ?`
	err = model.Db.QueryRow(select_user, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
