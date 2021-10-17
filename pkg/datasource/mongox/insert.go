package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (c *mongoCollection) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	collection := c.GetCollection()
	result, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *mongoCollection) InsertMany(documents []interface{}) (*mongo.InsertManyResult, error) {
	collection := c.GetCollection()
	result, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return nil, err
	}
	return result, nil
}
