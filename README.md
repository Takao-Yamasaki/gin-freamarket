# gin-fleamarket
## ginのインストール
```
go get -u github.com/gin-gonic/gin
```
## Goサーバーの起動
```
go run main.go
```
## airの起動
```
air
```
## curlで確認
- Ping
```
curl localhost:8080/ping
{"message":"pong"}
```
- FindAll関数
```
$ curl http://localhost:8080/items
{"data":[{"ID":1,"Name":"商品1","Price":1000,"Description":"説明1","SoldOut":false},{"ID":2,"Name":"商品2","Price":2000,"Description":"説明2","SoldOut":true},{"ID":3,"Name":"商品3","Price":3000,"Description":"説明3","SoldOut":false}]}
```
- FindBy関数
商品が見つかった場合
```
curl http://localhost:8080/items/1
{"data":{"ID":1,"CreatedAt":"2024-02-04T18:22:08.836+09:00","UpdatedAt":"2024-02-04T18:22:08.836+09:00","DeletedAt":null,"Name":"商品1","Price":1000,"Description":"Postの動作確認","SoldOut":false}}
```
商品が見つからない場合
```
curl http://localhost:8080/items/2
{"error":"item not found"}
```
-  Create関数
```
curl -POST -H "Content-Type: application/json" -d "{\"name\": \"商品4\", \"price\": 4000, \"description\": \"Postの動作確認\"}" http://localhost:8080/items
```
- Update関数
```
curl -X PUT -H "Content-Type: application/json" -d "{\"Name\":\"Update Test\"}" http://localhost:8080/items/1
```
- Delete関数
```
curl -X DELETE http://localhost:8080/items/1
```

- Signup関数
```
curl -X POST -H "Content-Type: application/json" -d "{\"email\":\"user1@example.com\", \"password\":\"user1pass\"}" http://localhost:8080/auth/signup
```
- Login関数
```
curl -X POST -H "Content-Type: application/json" -d "{\"email\":\"user1@example.com\", \"password\":\"user1pass\"}" http://localhost:8080/auth/login
```
-  Create関数(認証付き)
```
curl -POST -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzA3MjMzMzkyLCJzdWIiOjJ9.8pLhswOhbKnuym5ASsS6dO0pvAbJfIozrGXYl3vLz0s" -d "{\"name\": \"商品4\", \"price\": 4000, \"description\": \"Postの動作確認\"}" http://localhost:8080/items
```
- FindById関数(認証付き)
```
curl -POST -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzA3NDA2NzIzLCJzdWIiOjJ9.Xqe-EGXji1jSzIJ0S_eEJUyx-xCwb7-Y9lohMW2sRcg" http://localhost:8080/items/1
```

## myphpadmin
```
http://localhost:81
```

## GORMのインストール
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/postgres
```

## TODO
- テストケースの追加
- cookieを使った認証
- memchacheの使用
