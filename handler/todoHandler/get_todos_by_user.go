package todoHandler

import (
	"fmt"
	"go_portofolio/model"
	"log"
)

type User model.User

func (u *User) GetTodosByUser() (t []Todo, err error) {
	const getTodoByUser = `SELECT id, content, user_id FROM todos
	WHERE user_id = ?`

	rows, err := model.Db.Query(getTodoByUser, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserId)
		if err != nil {
			log.Fatalln(err)
		}
		t = append(t, todo)
	}
	defer rows.Close()
	fmt.Println(t)

	return
}
