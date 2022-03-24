package todoHandler

import (
	"go_portofolio/model"
	"log"
)

func GetTodo(id int) (t Todo, err error) {
	// 恐らく構造体を初期化している
	t = Todo{}

	const getT = `select id, content, user_id 
	from todos where id = ?`

	err = model.Db.QueryRow(getT, id).Scan(
		&t.ID,
		&t.Content,
		&t.UserId,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
