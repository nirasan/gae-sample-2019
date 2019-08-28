package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// HTTP リクエストの処理関数の登録
	// `/` へのリクエストで `indexHandler` 関数を実行する
	http.HandleFunc("/", indexHandler)
	// `/public/` へのリクエストで /public ディレクトリ以下の静的ファイルを配信する
	http.Handle("/public/", http.FileServer(http.Dir("./")))
	// Datastore へ読み書きをする関数の登録
	http.HandleFunc("/user/datastore/", datastoreUserHandler)

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

func render(w io.Writer, filename string, data interface{}) error {
	tmpl, err := template.ParseFiles("template/layout.html", "template/"+filename)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, data); err != nil {
		return err
	}

	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := render(w, "index.html", nil)
	if err != nil {
		internalServerErrorHandler(w, err)
	}
}

func datastoreUserHandler(w http.ResponseWriter, r *http.Request) {
	c, err := NewDatastoreClient()
	if err != nil {
		internalServerErrorHandler(w, err)
		return
	}
	switch r.Method {
	case "GET":
		if err := datastoreUserGetHandler(w, c); err != nil {
			internalServerErrorHandler(w, err)
			return
		}
	case "POST":
		if err := datastoreUserPostHandler(w, r, c); err != nil {
			internalServerErrorHandler(w, err)
			return
		}
	}
}

func datastoreUserGetHandler(w http.ResponseWriter, c *datastoreClient) error {
	users, err := c.ListUsers()
	if err != nil {
		return err
	}

	if err := render(w, "user_datastore.html", map[string]interface{}{"Users": users}); err != nil {
		return err
	}

	return nil
}

func datastoreUserPostHandler(w http.ResponseWriter, r *http.Request, c *datastoreClient) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	u := &User{}
	age, err := strconv.ParseInt(r.Form.Get("Age"), 10, 64)
	if err != nil {
		return err
	}
	u.Age = age
	u.Name = r.Form.Get("Name")

	id, err := c.AddUser(u)
	if err != nil {
		return err
	}
	u.ID = id

	if err := render(w, "user_datastore.html", map[string]interface{}{"Created": u}); err != nil {
		return err
	}

	return nil
}

func internalServerErrorHandler(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
