package todoHandler

import (
	"go_portofolio/model"
	"log"
	"time"
)

type Todo model.Todo

func (t Todo) CreateTodo() (err error) {
	// createTodo test
	t.Content = "test todo"
	t.UserId = 1
	t.CreatedAt = time.Now()
	// createTodo test

	insertT := `insert into todos (
		content, user_id, created_at) values (?, ?, ?)`

	_, err = model.Db.Exec(insertT, t.Content, t.UserId, t.CreatedAt)
	if err != nil {
		log.Print(err)
	}
	defer model.Db.Close()

	return
}
