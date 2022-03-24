package todoHandler

import (
	"go_portofolio/model"
	"log"
)

func (t *Todo) UpdateTodo() (err error) {
	// update todo test
	t.Content = "Update Content"
	t.UserId = 1
	// update todo test

	const updateT = `UPDATE todos SET content = ?, user_id = ? 
	WHERE id = ?;`

	stmt, err := model.Db.Prepare(updateT)
	_, err = stmt.Exec(t.Content, t.UserId, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	defer model.Db.Close()

	return
}
