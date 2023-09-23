package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	_, err := db.NewPGStore[string]("pgx", os.Getenv("DB_CONNECTION"))
	fmt.Println(err)

	r := mux.NewRouter()
	r.HandleFunc("/", handler.DecorateHTTPFunc(handler.Greet))
	fmt.Println("Server runnning")
	http.ListenAndServe(":8080", r)
}
