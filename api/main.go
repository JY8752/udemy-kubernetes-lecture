package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello World!")) })

	server.ListenAndServe()
}
