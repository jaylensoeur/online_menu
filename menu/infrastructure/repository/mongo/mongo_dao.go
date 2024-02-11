package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDao struct {
	uri    string
	client *mongo.Client
}

func NewMongoDao(uri string) *MongoDao {
	return &MongoDao{
		uri: uri,
	}
}

func (m *MongoDao) Connect() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.uri))
	m.client = client
}

func (m *MongoDao) Disconnect() {
	err := m.client.Disconnect(context.TODO())
	if err != nil {
		return
	}
}

func (m *MongoDao) Query(database string, collection string) *mongo.Collection {
	return m.client.Database(database).Collection(collection)
}
