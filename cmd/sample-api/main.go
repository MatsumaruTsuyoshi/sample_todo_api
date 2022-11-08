package main

import (
	"net/http"

	"github.com/MatsumaruTsuyoshi/sample_todo_api/controller"
	"github.com/MatsumaruTsuyoshi/sample_todo_api/model/repository"
)

// DIしている
var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	// localhost:8080で起動する
	server := http.Server{
		Addr: ":8080",
	}
	// localhost:8080/todos/ に届いたリクエストをro.HandleTodosRequestで処理する
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe() // ここでサーバー起動
}
