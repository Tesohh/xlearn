package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/Tesohh/xlearn/handler/adventurehandler"
	"github.com/Tesohh/xlearn/handler/generalhandler"
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

	r := router{
		r:      mux.NewRouter().NewRoute().PathPrefix("/api").Subrouter(),
		stores: stores,
		routes: map[string]route{
			"/user/signup": {handler: userhandler.Signup, modifiers: "unprotected", methods: "POST"},
			"/user/login":  {handler: userhandler.Login, modifiers: "unprotected", methods: "POST"},
			"/user/logout": {handler: userhandler.Logout, methods: "GET"},

			"/user/me":                  {handler: userhandler.Me, methods: "GET"},
			"/user/me/settings/edit":    {handler: userhandler.EditSettings, methods: "POST"},
			"/user/org/join/{code}":     {handler: userhandler.JoinOrg, methods: "POST"},
			"/user/org/leave/@{orgtag}": {handler: userhandler.LeaveOrg, methods: "POST"},
			"/user/org/joined":          {handler: userhandler.JoinedOrgs, methods: "GET"},
			"/user/org/joined/tags":     {handler: userhandler.JoinedOrgsTags, methods: "GET"},

			"/org/new":            {handler: orghandler.New, methods: "POST"},
			"/org/@{orgtag}":      {handler: orghandler.One, modifiers: "protectorg", methods: "GET"},
			"/org/@{orgtag} ":     {handler: orghandler.Edit, modifiers: "admin,protectorg", methods: "POST"},
			"/org/@{orgtag}/meta": {handler: orghandler.Meta, modifiers: "protectorg", methods: "GET"},

			"/org/@{orgtag}/code/{uses}":       {handler: orghandler.Code, modifiers: "admin,protectorg", methods: "GET"},
			"/org/@{orgtag}/revokecode/{code}": {handler: orghandler.RevokeCode, modifiers: "admin,protectorg", methods: "POST"},

			"/org/@{orgtag}/adventure/new":                {handler: adventurehandler.New, modifiers: "admin,protectorg", methods: "POST"},
			"/org/@{orgtag}/adventure/all":                {handler: adventurehandler.All, modifiers: "protectorg", methods: "GET"},
			"/org/@{orgtag}/adventure/@{advtag}":          {handler: adventurehandler.One, modifiers: "protectorg", methods: "GET"},
			"/org/@{orgtag}/adventure/@{advtag} ":         {handler: adventurehandler.Edit, modifiers: "admin,protectorg", methods: "POST"},
			"/org/@{orgtag}/adventure/@{advtag}/movestep": {handler: adventurehandler.MoveStep, modifiers: "teacher,protectorg", methods: "POST"},

			"/step/new":         {handler: stephandler.New, modifiers: "teacher", methods: "POST"},
			"/step/many":        {handler: stephandler.Many, methods: "GET"},
			"/step/@{steptag}":  {handler: stephandler.One, methods: "GET"},
			"/step/@{steptag} ": {handler: stephandler.Edit, modifiers: "teacher", methods: "POST"},

			"/danger/mockdb": {handler: generalhandler.MockDB, modifiers: "unprotected", methods: "POST"},
			"/isdbempty":     {handler: generalhandler.IsDBEmptyEndpoint, modifiers: "unprotected", methods: "GET"},
		},
	}

	fmt.Println("Server running on http://localhost:8080")
	r.serve(":8080")
}
