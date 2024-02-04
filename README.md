# gin-fleamarket
## ginのインストール
```
go get -u github.com/gin-gonic/gin
```
## Goサーバーの起動
```
go run main.go
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

## myphpadmin
```
http://localhost:81
```

## GORM
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/postgres
```
