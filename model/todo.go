package model

type (
	Todo struct {
		ID   int    `json:"id"`
		Text string `json:"text"`
	}

	TodoUsecase interface {
		AllGet() ([]Todo, error)
		StatusUpdate(id int) error
		Store(todo Todo) error
		Delete(id int) error
	}

	TodoRepository interface {
		AllGet() ([]Todo, error)
		StatusUpdate(id int) error
		Store(todo Todo) error
		Delete(id int) error
	}
)
