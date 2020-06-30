package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hdlFunc)
	port := getPORT()
	// fmt.Println("serve on", port)
	fmt.Printf("serve on %s\n", port)
	http.ListenAndServe(getPORT(), nil)
}

func hdlFunc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	switch path {
	case "/contact":
		contact(w, r)
	case "/":
		root(w, r)
	default:
		p404(w, r)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Contact</h1>  
		<p><a href="mailto:christophe.dhayer@gmail.com">Mail to me</a></p>
	`))
}

func root(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>hello you how are you</h1>")
}

func p404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1>error 404 page not found</h1>`)
}

func getPORT() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	port = ":" + port
	return port
}
