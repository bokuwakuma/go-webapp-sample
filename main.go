package main

import (
	"net/http"
	"text/template"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	// dir, _ := os.Getwd()
	// http.HandleFunc("/", index)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	// http.Handle("/node_modules/", http.StripPrefix("/node_modules", http.FileServer(http.Dir(dir+"/node_modules"))))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "localhost:8090",
		Handler: mux,
	}

	server.ListenAndServe()

}

// indexのハンドラ関数
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tempates/index.html")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// if err := t.Execute(w, nil); err != nil {
	// 	panic(err.Error())
	// }
	if err == nil {
		t.Execute(w, nil)
	}
}
