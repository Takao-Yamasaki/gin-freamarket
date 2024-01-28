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
-  Create関数
```
curl -POST -H "Content-Type: application/json" -d "{\"name\": \"商品4\", \"price\": 4000, \"description\": \"Postの動作確認\"}" http://localhost:8080/items
```
