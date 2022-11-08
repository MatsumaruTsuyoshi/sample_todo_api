# sample_todo_api

こちらの記事を参考にしてます。

https://dev.classmethod.jp/articles/go-sample-rest-api/


## 動作確認

```
make serve
```

### todo取得
```
curl -i localhost:8080/todos/
```

### todo追加
```
curl -i -X POST -H "Content-Type: application/json" -d '{"title":"test", "content":"テストです。"}' localhost:8080/todos/
```

### todo更新
```
curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"test", "content":"変更テスト"}' localhost:8080/todos/4
```

### todo更新
```
curl -i -X DELETE localhost:8080/todos/4
```