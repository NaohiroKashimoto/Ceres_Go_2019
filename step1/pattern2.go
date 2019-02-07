package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//
	ceresHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello,Ceres!!\n")
	}

	//デフォルト
	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello,World!!\n")
	}

	//ハンドラーの登録
	http.HandleFunc("/ceres", ceresHandler)
	http.HandleFunc("/", defaultHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
