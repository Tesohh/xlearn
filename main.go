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

// func handle(r *mux.Router, path string, f handler.APIFunc, stores db.StoreHolder, modifiers ...string) {
// 	r.HandleFunc("/api"+path, handler.DecorateHTTPFunc(f, stores, modifiers))
// }

func main() {
	client, err := db.NewMongoClient()
	if err != nil { // if it doesnt connect to mongo, it needs to panic out
		log.Fatal(err)
	}

	stores := db.StoreHolder{
		Users:      db.MongoStore[data.User]{Client: client, Coll: client.Database("main").Collection("users")},
		Orgs:       db.MongoStore[data.Org]{Client: client, Coll: client.Database("main").Collection("orgs")},
		Adventures: db.MongoStore[data.Adventure]{Client: client, Coll: client.Database("main").Collection("adventures")},
	}

	r := mux.NewRouter().NewRoute().PathPrefix("/api").Subrouter()
	// auth
	auth := r.NewRoute().PathPrefix("/user").Subrouter()
	auth.HandleFunc("/signup", handler.MW(handler.UserSignup, stores, "unprotected")).Methods("POST")
	auth.HandleFunc("/login", handler.MW(handler.UserLogin, stores, "unprotected")).Methods("GET")
	auth.HandleFunc("/logout", handler.MW(handler.UserLogout, stores)).Methods("GET")

	// user
	user := r.NewRoute().PathPrefix("/user").Subrouter()
	user.HandleFunc("/one", handler.MW(handler.OneUser, stores)).Methods("GET")
	user.HandleFunc("/me", handler.MW(handler.UserMe, stores)).Methods("GET")

	// org
	orgGeneric := r.NewRoute().PathPrefix("/org").Subrouter()
	org := orgGeneric.NewRoute().PathPrefix("/@{orgtag}").Subrouter()
	orgGeneric.HandleFunc("/new", handler.MW(handler.OrgNew, stores, "admin")).Methods("POST")
	org.HandleFunc("", handler.MW(handler.Org, stores)).Methods("GET")
	org.HandleFunc("/meta", handler.MW(handler.OrgMeta, stores)).Methods("GET")

	// org adventures
	advGeneric := r.NewRoute().PathPrefix("/org/@{orgtag}/adventure").Subrouter()
	adv := advGeneric.NewRoute().PathPrefix("/@{advtag}").Subrouter()

	adv.HandleFunc("", handler.MW(handler.OrgAdventureOne, stores)).Methods("GET")
	advGeneric.HandleFunc("/new", handler.MW(handler.OrgAdventureNew, stores, "admin")).Methods("POST")
	advGeneric.HandleFunc("/all", handler.MW(handler.OrgAdventuresAll, stores)).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
