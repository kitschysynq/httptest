package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%#v\n", r)
		//fmt.Fprintf(w, "ok\n")
		http.Error(w, "internal server error", 500)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
