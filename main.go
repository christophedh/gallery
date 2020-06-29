package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hdlFunc)
	port := getPORT()
	fmt.Println("serve on ", port)
	http.ListenAndServe(getPORT(), nil)
}

func hdlFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hello you how are you</h1>")
}

func getPORT() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	port = ":" + port
	return port
}
