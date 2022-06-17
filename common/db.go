package common

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database is mongodb adapter
type Database struct {
	cancel context.CancelFunc
	client *mongo.Client
}

// NewDatabase will create new mongodb connection
func NewDatabase() (*Database, error) {
	host := "localhost:27017"

	if m, exists := os.LookupEnv("MONGO_HOST"); exists {
		host = m
	}

	opts := options.Client().ApplyURI("mongodb://" + host)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	// Create database client
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		defer cancel()

		return nil, err
	}

	d := new(Database)
	d.cancel = cancel
	d.client = client

	return d, nil
}

// Client will return active mongodb client
func (d *Database) Client() *mongo.Client {
	if d.client == nil {
		return nil
	}

	return d.client
}

// Close active mongodb connection
func (d *Database) Close() {
	if d.cancel == nil {
		return
	}

	d.cancel()
}
