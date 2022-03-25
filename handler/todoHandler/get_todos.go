package todoHandler

import (
	"fmt"
	"go_portofolio/model"
	"log"
)

// []Todo スライスのtodoという意味
func GetTodos() (t []Todo, err error) {
	const getTodos = `SELECT id, content, user_id FROM todos;`
	rows, err := model.Db.Query(getTodos)
	fmt.Println(rows)
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
			log.Fatal(err)
		}
		t = append(t, todo)
	}
	defer rows.Close()

	return
}
