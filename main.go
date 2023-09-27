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
	decoratedHandle(r, "/unprotected/user/signup", handler.UserSignup, stores)
	decoratedHandle(r, "/unprotected/user/login", handler.UserLogin, stores)

	// user
	decoratedHandle(r, "/user/one", handler.OneUser, stores)
	decoratedHandle(r, "/user/me", handler.UserMe, stores)

	// org
	decoratedHandle(r, "/admin/org/new", handler.OrgNew, stores)
	decoratedHandle(r, "/org/@{tag}", handler.Org, stores)
	decoratedHandle(r, "/org/@{tag}/meta", handler.OrgMeta, stores)
	decoratedHandle(r, "/org/@{tag}/adventure/all", handler.OrgAdventuresAll, stores)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
