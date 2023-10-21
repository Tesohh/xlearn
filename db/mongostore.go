package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrDocumentNotFound = errors.New("couldn't find document")
)

func (q Query) ToMongo() primitive.D {
	d := primitive.D{}
	for k, v := range q {
		d = append(d, primitive.E{Key: k, Value: v})
	}
	return d
}

func NewMongoClient(connection string) (*mongo.Client, error) {
	// connection := os.Getenv("DB_CONNECTION")
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
	res := s.Coll.FindOne(context.Background(), q.ToMongo())
	var document T
	res.Decode(&document)

	if document.IsEmpty() {
		return nil, ErrDocumentNotFound
	}
	return &document, nil
}

func (s MongoStore[T]) Many(q Query) ([]T, error) {
	cur, err := s.Coll.Find(context.Background(), q.ToMongo())
	if err != nil {
		return nil, err
	}

	var results []T
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s MongoStore[T]) Put(doc T) error {
	_, err := s.Coll.InsertOne(context.Background(), doc, nil)
	return err
}

func (s MongoStore[T]) Update(q Query, doc T) error {
	update := bson.M{"$set": doc}
	_, err := s.Coll.UpdateOne(context.Background(), q.ToMongo(), update)
	return err
}

func (s MongoStore[T]) Delete(q Query) error {
	_, err := s.Coll.DeleteOne(context.Background(), q.ToMongo())
	return err
}
