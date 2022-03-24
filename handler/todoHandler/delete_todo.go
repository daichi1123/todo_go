package todoHandler

import (
	"go_portofolio/model"
	"log"
)

func (t *Todo) DeleteTodo() (err error) {
	const deleteT = `DELETE FROM todos WHERE id = ?`
	stmt, _ := model.Db.Prepare(deleteT)
	_, err = stmt.Exec(t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	model.Db.Close()

	return
}
