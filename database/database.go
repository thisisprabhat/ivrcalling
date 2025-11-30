package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// InitDB initializes MongoDB connection
func InitDB(mongoURI, dbName string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	database := client.Database(dbName)

	// Create indexes
	if err := createIndexes(ctx, database); err != nil {
		return nil, fmt.Errorf("failed to create indexes: %w", err)
	}

	return &MongoDB{
		Client:   client,
		Database: database,
	}, nil
}

// createIndexes creates necessary indexes for collections
func createIndexes(ctx context.Context, db *mongo.Database) error {
	// Campaign indexes
	campaignIndexes := []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"name": 1},
		},
		{
			Keys: map[string]interface{}{"is_active": 1},
		},
	}
	_, err := db.Collection("campaigns").Indexes().CreateMany(ctx, campaignIndexes)
	if err != nil {
		return fmt.Errorf("failed to create campaign indexes: %w", err)
	}

	// Call indexes
	callIndexes := []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"campaign_id": 1},
		},
		{
			Keys: map[string]interface{}{"status": 1},
		},
		{
			Keys: map[string]interface{}{"twilio_call_sid": 1},
		},
		{
			Keys: map[string]interface{}{"phone_number": 1},
		},
	}
	_, err = db.Collection("calls").Indexes().CreateMany(ctx, callIndexes)
	if err != nil {
		return fmt.Errorf("failed to create call indexes: %w", err)
	}

	// CallLog indexes
	callLogIndexes := []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"call_id": 1},
		},
		{
			Keys: map[string]interface{}{"created_at": -1},
		},
	}
	_, err = db.Collection("call_logs").Indexes().CreateMany(ctx, callLogIndexes)
	if err != nil {
		return fmt.Errorf("failed to create call_log indexes: %w", err)
	}

	return nil
}

// Close closes the MongoDB connection
func (m *MongoDB) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

// Collection returns a MongoDB collection
func (m *MongoDB) Collection(name string) *mongo.Collection {
	return m.Database.Collection(name)
}
