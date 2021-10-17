package mongox

import "context"

func (c *mongoCollection) DeleteOne(filter interface{}) error {
	collection := c.GetCollection()
	_, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *mongoCollection) DeleteMany(filter interface{}) error {
	collection := c.GetCollection()
	_, err := collection.DeleteMany(context.TODO(), filter, nil)
	if err != nil {
		return err
	}
	return nil
}
