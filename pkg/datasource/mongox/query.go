package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *mongoCollection) FindOne(filter interface{}) *mongo.SingleResult {
	collection := c.GetCollection()
	result := collection.FindOne(context.TODO(), filter)
	return result
}

func (c *mongoCollection) FindMany(filter interface{}, options *options.FindOptions) (*mongo.Cursor, error) {
	collection := c.GetCollection()
	return collection.Find(context.TODO(), filter, options)
}

func (c *mongoCollection) Count() (int64, error) {
	collection := c.GetCollection()
	return collection.EstimatedDocumentCount(context.TODO())
}
