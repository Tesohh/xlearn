package main

import (
	"fmt"
	"net/http"

	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.DecorateHTTPFunc(handler.Greet))
	fmt.Println("Server runnning")
	http.ListenAndServe(":8080", r)
}
