package nosql

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repository by nosql (mongo)
type Repository struct {
	client     *mongo.Client
	context    context.Context
	collection *mongo.Collection
}
