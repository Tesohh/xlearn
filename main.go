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

func handle(r *mux.Router, path string, f handler.APIFunc, stores db.StoreHolder) {
	r.HandleFunc("/api"+path, handler.DecorateHTTPFunc(f, stores))
}

func main() {
	client, err := db.NewMongoClient()
	if err != nil { // if it doesnt connect to mongo, it needs to panic out
		log.Fatal(err)
	}

	stores := db.StoreHolder{
		Users: db.MongoStore[data.User]{Client: client, Coll: client.Database("main").Collection("users")},
		Orgs:  db.MongoStore[data.Org]{Client: client, Coll: client.Database("main").Collection("orgs")},
	}

	r := mux.NewRouter()
	// auth
	handle(r, "/unprotected/user/signup", handler.UserSignup, stores)
	handle(r, "/unprotected/user/login", handler.UserLogin, stores)

	// user
	handle(r, "/user/one", handler.OneUser, stores)
	handle(r, "/user/me", handler.UserMe, stores)

	// org
	handle(r, "/admin/org/new", handler.OrgNew, stores)
	handle(r, "/org/@{tag}", handler.Org, stores)
	handle(r, "/org/@{tag}/meta", handler.OrgMeta, stores)

	// org adventures
	handle(r, "/org/@{tag}/adventure/all", handler.OrgAdventuresAll, stores)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
