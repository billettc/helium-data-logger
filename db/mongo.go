package db

import (
	"context"
	"github.com/billettc/helium-data-logger/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE = "helium-hotspot-tracker"
const COLLECTION_LOG_EVENTS = "log_events"

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB(address string) (*MongoDB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &MongoDB{client: client}, nil

}

func (db *MongoDB) SaveLogEvent(event *models.LogEvent) error {
	collection := db.client.Database(DATABASE).Collection(COLLECTION_LOG_EVENTS)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//filter := bson.M{"datetime": price.DateTime, "symbol": price.Symbol}
	//update := bson.M{"$set": price}
	_, err := collection.InsertOne(ctx, event)
	if err != nil {
		return err
	}

	return nil
}
