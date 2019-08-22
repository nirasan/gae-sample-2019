package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// HTTP リクエストの処理関数の登録
	// `/` へのリクエストで `indexHandler` 関数を実行する
	http.HandleFunc("/", indexHandler)

	// HTTP サーバーの待受ポートの設定
	// GAE では環境変数 PORT で待受ポートが指定される
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	// HTTP サーバーの起動
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
