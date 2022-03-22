package handler

import (
	"go_portofolio/database"
	"html/template"
	l "log"
	"net/http"
)

type Todo struct {
	Id   int
	Task string
}

// errorが出た
// _, err := db.Exec("CREATE TABLE go_api_database.todos (id int, task string);")

// 関数は、capital letter is uppercase letter
func Index(w http.ResponseWriter, r *http.Request) {
	db := database.Db_init()
	// 下記にQuery文を記述する
	selDB, err := db.Query("SELECT * FROM todos;")
	if err != nil {
		l.Println(err, "errorです")
	}

	todos := Todo{
		Id:   1,
		Task: "テストタスクです",
	}
	// res := []Todo{}
	// fmt.Println(todos)
	// fmt.Println("wの内容", w)

	for selDB.Next() {
		var id int
		var task string
		err = selDB.Scan(&id, &task)
		if err != nil {
			panic(err.Error())
		}
		// todos.Id = id
		// todos.Task = task
		// res = append(res, todos)
	}

	t, err := template.ParseFiles("src/template/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, todos); err != nil {
		panic(err.Error())
	}
}

// func CreateTodos(ctx context.Context, task string) (Todo error) {
// const (
// 	insert  = `INSERT INTO todos(subject) VALUES(?)`
// 	confirm = `SELECT task FROM todos WHERE id = ?`
// )

// db := database.Db_init()

// stmt, err := db.PrepareContext(ctx, insert)
// if err != nil {
// 	return err
// }
// defer stmt.Close()

// result, err := stmt.Exec(task)
// if err != nil {
// 	return err
// }

// id, err := result.LastInsertId()
// if err != nil {
// 	return err
// }

// err = db.QueryRowContext(ctx, confirm, id).Scan(
// 	&task)
// if err != nil {
// 	l.Print(err)
// }

// return Todo, err
// }
