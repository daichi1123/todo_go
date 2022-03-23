package userHandler

import (
	"go_portofolio/model"
	"log"
	"time"
)

// 外部パッケージの構造体をレシーバとして使用するとエラーになる
// ローカルパッケージ以外でのレシーバ定義はエラーになる
// その場合の対処方法としては、ローカルに定義し直すといい
type User model.User

// user型のメソッド 引数渡さず 返り値エラー
// ポインタレシーバを値レシーバに変更したらエラーが消えた
func (u User) CreateUser() (err error) {
	// CreateUserのTEST用
	u.Name = "test1"
	u.Email = "test@example.com"
	u.Password = "test"
	// CreateUserのTEST用

	insert := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = model.Db.Exec(
		insert,
		model.CreateUUID(),
		u.Name,
		u.Email,
		model.Encypt(u.Password),
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	// Closeをしないととめどなく作り続ける
	model.Db.Close()

	return
}
