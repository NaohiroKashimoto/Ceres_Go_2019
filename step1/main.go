package main

import (
	"github.com/prometheus/common/log"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello,Ceres!!\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
