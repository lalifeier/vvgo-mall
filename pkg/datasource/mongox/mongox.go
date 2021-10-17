package mongox

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	mongoCollection struct {
		client     *mongo.Client
		db         *mongo.Database
		collection string
	}

	MongoCollection interface {
		GetCollection() *mongo.Collection

		InsertOne(document interface{}) (*mongo.InsertOneResult, error)
		InsertMany(documents []interface{}) (*mongo.InsertManyResult, error)

		DeleteOne(filter interface{}) error
		DeleteMany(filter interface{}) error

		UpdateOne(filter, update interface{}) error
		UpdateMany(filter, update interface{}) error

		FindOne(filter interface{}) *mongo.SingleResult
		FindMany(filter interface{}, options *options.FindOptions) (*mongo.Cursor, error)

		Count() (int64, error)
	}
)

func New(client *mongo.Client, database, collection string) MongoCollection {
	return &mongoCollection{
		client:     client,
		db:         client.Database(""),
		collection: collection,
	}
}

func NewClient(url string, opts *options.ClientOptions) *mongo.Client {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (c *mongoCollection) GetCollection() *mongo.Collection {
	return c.db.Collection(c.collection)
}
