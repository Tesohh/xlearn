package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler"
	"github.com/Tesohh/xlearn/handler/adventurehandler"
	"github.com/Tesohh/xlearn/handler/orghandler"
	"github.com/Tesohh/xlearn/handler/stephandler"
	"github.com/Tesohh/xlearn/handler/userhandler"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client, err := db.NewMongoClient(os.Getenv("DB_CONNECTION"))
	if err != nil { // if it doesnt connect to mongo, it needs to panic out
		log.Fatal(err)
	}

	stores := db.StoreHolder{
		Users:      db.MongoStore[data.User]{Client: client, Coll: client.Database("main").Collection("users")},
		Orgs:       db.MongoStore[data.Org]{Client: client, Coll: client.Database("main").Collection("orgs")},
		Adventures: db.MongoStore[data.Adventure]{Client: client, Coll: client.Database("main").Collection("adventures")},
		Steps:      db.MongoStore[data.Step]{Client: client, Coll: client.Database("main").Collection("steps")},
	}

	r := mux.NewRouter().NewRoute().PathPrefix("/api").Subrouter()

	// auth
	auth := r.NewRoute().PathPrefix("/user").Subrouter()
	auth.HandleFunc("/signup", handler.MW(userhandler.Signup, stores, "unprotected")).Methods("POST")
	auth.HandleFunc("/login", handler.MW(userhandler.Login, stores, "unprotected")).Methods("POST")
	auth.HandleFunc("/logout", handler.MW(userhandler.Logout, stores)).Methods("GET")

	// user
	user := r.NewRoute().PathPrefix("/user").Subrouter()
	user.HandleFunc("/me", handler.MW(userhandler.Me, stores)).Methods("GET")
	user.HandleFunc("/org/join/{code}", handler.MW(userhandler.JoinOrg, stores)).Methods("POST")
	user.HandleFunc("/org/leave/@{orgtag}", handler.MW(userhandler.LeaveOrg, stores)).Methods("POST")
	user.HandleFunc("/org/joined", handler.MW(userhandler.JoinedOrgs, stores)).Methods("GET")
	user.HandleFunc("/org/joined/tags", handler.MW(userhandler.JoinedOrgsTags, stores)).Methods("GET")

	// org
	orgGeneric := r.NewRoute().PathPrefix("/org").Subrouter()
	org := orgGeneric.NewRoute().PathPrefix("/@{orgtag}").Subrouter()

	orgGeneric.HandleFunc("/new", handler.MW(orghandler.New, stores, "admin")).Methods("POST")
	org.HandleFunc("", handler.MW(orghandler.One, stores, "protectorg")).Methods("GET")
	org.HandleFunc("", handler.MW(orghandler.Edit, stores, "admin", "protectorg")).Methods("POST")
	org.HandleFunc("/meta", handler.MW(orghandler.Meta, stores, "protectorg")).Methods("GET")
	org.HandleFunc("/code/{uses}", handler.MW(orghandler.Code, stores, "protectorg", "admin")).Methods("GET")
	org.HandleFunc("/revokecode/{code}", handler.MW(orghandler.RevokeCode, stores, "protectorg", "admin")).Methods("POST")

	// org adventures
	advGeneric := r.NewRoute().PathPrefix("/org/@{orgtag}/adventure").Subrouter()
	adv := advGeneric.NewRoute().PathPrefix("/@{advtag}").Subrouter()

	advGeneric.HandleFunc("/new", handler.MW(adventurehandler.New, stores, "admin", "protectorg")).Methods("POST")
	advGeneric.HandleFunc("/all", handler.MW(adventurehandler.All, stores, "protectorg")).Methods("GET")
	adv.HandleFunc("/movestep", handler.MW(adventurehandler.MoveStep, stores, "teacher", "protectorg")).Methods("POST")
	adv.HandleFunc("", handler.MW(adventurehandler.One, stores, "protectorg")).Methods("GET")
	adv.HandleFunc("", handler.MW(adventurehandler.Edit, stores, "admin", "protectorg")).Methods("POST")

	// steps
	stepGeneric := r.NewRoute().PathPrefix("/step").Subrouter()
	step := stepGeneric.NewRoute().PathPrefix("/@{steptag}").Subrouter()

	stepGeneric.HandleFunc("/new", handler.MW(stephandler.New, stores, "teacher")).Methods("POST")
	stepGeneric.HandleFunc("/many", handler.MW(stephandler.Many, stores)).Methods("GET")
	step.HandleFunc("", handler.MW(stephandler.One, stores)).Methods("GET")
	step.HandleFunc("", handler.MW(stephandler.Edit, stores, "teacher")).Methods("POST")

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
