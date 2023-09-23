package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func decoratedHandle(r *mux.Router, path string, f handler.APIFunc, stores db.StoreHolder) {
	r.HandleFunc(path, handler.DecorateHTTPFunc(handler.Greet, stores))
}

func main() {
	client, err := db.NewMongoClient()
	if err != nil { // if it doesnt connect to mongo, it needs to panic out
		log.Fatal(err)
	}

	stores := db.StoreHolder{
		Users: db.MongoStore[data.User]{Client: client, Coll: client.Database("main").Collection("users")},
	}

	r := mux.NewRouter()
	decoratedHandle(r, "/", handler.Greet, stores)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
