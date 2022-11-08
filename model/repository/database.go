package repository

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

// "github.com/MatsumaruTsuyoshi/sample_todo_api/model/repository"がimportされたタイミングで動作する
// main.goのメイン処理より先に実行される
func init() {
	var err error
	// DSNは[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]の仕様
	// ここで設定しているusername、password、address、dbnameは動作確認用DBのdocker-compose.ymlとDockerfileの値を設定している
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "todo-app", "todo-password", "sample-api-db:3306", "todo")
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}
