package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() (*mongo.Client, error) {
	connection := os.Getenv("DB_CONNECTION")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	return client, nil
}

type IsEmptier interface {
	IsEmpty() bool
}

type MongoStore[T IsEmptier] struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func (s MongoStore[T]) One(q Query) (*T, error) {
	panic("not implemented") // TODO: Implement
}

func (s MongoStore[T]) Many(q Query) ([]T, error) {
	panic("not implemented") // TODO: Implement
}

func (s MongoStore[T]) Put(doc T) error {
	panic("not implemented") // TODO: Implement
}

func (s MongoStore[T]) Update(q Query, doc T) error {
	panic("not implemented") // TODO: Implement
}

func (s MongoStore[T]) Delete(q Query) error {
	panic("not implemented") // TODO: Implement
}
