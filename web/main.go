package main

import (
	"fmt"
	"net/http"
    "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
	//formats and writes
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", router)
	
    if err != nil {
        fmt.Println(err)
    }
}
