package mock

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/Tesohh/xlearn/data"
	"github.com/Tesohh/xlearn/db"
	"github.com/joho/godotenv"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	godotenv.Load()
}

func AddDataToStores(stores db.StoreHolder) error {
	for _, u := range users {
		err := stores.Users.Put(u)
		if err != nil {
			return err
		}
	}
	for _, o := range orgs {
		err := stores.Orgs.Put(o)
		if err != nil {
			return err
		}
	}
	for _, a := range adventures {
		err := stores.Adventures.Put(a)
		if err != nil {
			return err
		}
	}
	for _, s := range steps {
		err := stores.Steps.Put(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func Stores() (db.StoreHolder, error) {
	fmt.Println(os.Getwd())
	client, err := db.NewMongoClient(os.Getenv("DB_CONNECTION"))
	if err != nil { // if it doesnt connect to mongo, it needs to panic out
		return db.StoreHolder{}, err
	}

	// reset the database
	err = client.Database("mock").Drop(context.Background())
	if err != nil {
		return db.StoreHolder{}, err
	}

	stores := db.StoreHolder{
		Users:      db.MongoStore[data.User]{Client: client, Coll: client.Database("mock").Collection("users")},
		Orgs:       db.MongoStore[data.Org]{Client: client, Coll: client.Database("mock").Collection("orgs")},
		Adventures: db.MongoStore[data.Adventure]{Client: client, Coll: client.Database("mock").Collection("adventures")},
		Steps:      db.MongoStore[data.Step]{Client: client, Coll: client.Database("mock").Collection("steps")},
	}

	err = AddDataToStores(stores)
	if err != nil {
		return db.StoreHolder{}, err
	}

	return stores, nil
}
