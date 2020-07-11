package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// ハンドラ
	r.HandleFunc("/", YourHandler)

	// パスパラメータ
	r.HandleFunc("/hello/{name}", VarsHandler)

	// パス変数で正規表現を使用
	r.HandleFunc("/hello/{name}/{age:[0-9]+}", RegexHandler)

	// クエリパラメータ
	r.HandleFunc("/hi/", QueryStringHandler)

	// サービス
	http.ListenAndServe(":8001", r)
}

// YourHandler 通常
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

// VarsHandler 変数
func VarsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%s Loves Gorilla\n", vars["name"])
}

// RegexHandler 正規表現
func RegexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%s is %s years old\n", vars["name"], vars["age"])
}

// QueryStringHandler クエリ
func QueryStringHandler(w http.ResponseWriter, r *http.Request) {
}
