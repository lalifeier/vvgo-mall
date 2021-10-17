package mongox

import "context"

func (c *mongoCollection) UpdateOne(filter, update interface{}) error {
	collection := c.GetCollection()
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (c *mongoCollection) UpdateMany(filter, update interface{}) error {
	collection := c.GetCollection()
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
