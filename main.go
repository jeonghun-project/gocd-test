package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const txt = "some txt"
	fmt.Println("dumdum: ", txt)
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Fatalln(http.ListenAndServe(":3200", nil))
}
