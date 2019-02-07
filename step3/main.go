package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"upper.io/db.v3/mysql"
)

//ユーザーの構造体定義
type User struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`
}

var settings = mysql.ConnectionURL{
	Host:     "127.0.0.1:3306",
	Database: "ceres",
	User:     "root",
	Password: "ceres",
}

func main() {
	//ルーター初期化
	r := chi.NewRouter()

	sess, err := mysql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		u := User{}
		query := sess.SelectFrom("users").Where("id = ?", chi.URLParam(r, "id"))
		query.One(&u)

		responseBody, err := json.Marshal(&u)
		if err != nil {
			//サーバーエラー
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(responseBody)

	})
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		//POSTされたJSONをParseする。
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		query := sess.InsertInto("users").Columns("name").Values(u.Name)
		res, err := query.Exec()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println(res)
		w.WriteHeader(http.StatusCreated)
	})

	//Putメソッドに対応したハンドラーの定義
	r.Put("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		//idのチェック
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

	})

	//DELETEメソッドに対応したハンドラーの定義
	r.Delete("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		query := sess.DeleteFrom("user").Where("id = ?", id)
		res, err := query.Exec()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println(res)
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", r)
}
