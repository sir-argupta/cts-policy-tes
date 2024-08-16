package services

import (
	"context"
	"log"
	"policy-search-api/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PolicyService handles business logic for policy-related operations
type PolicyService struct {
	collection *mongo.Collection
}

// NewPolicyService creates a new PolicyService
func NewPolicyService(collection *mongo.Collection) *PolicyService {
	return &PolicyService{
		collection: collection,
	}
}

// SearchPolicies allows searching policies based on various criteria
func (ps *PolicyService) SearchPolicies(criteria map[string]string) ([]domain.Policy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}

	if criteria["type"] != "" {
		filter["type"] = criteria["type"]
	}
	if criteria["company"] != "" {
		filter["company"] = criteria["company"]
	}
	if criteria["id"] != "" {
		filter["_id"] = criteria["id"]
	}
	if criteria["name"] != "" {
		filter["name"] = criteria["name"]
	}
	if criteria["years"] != "" {
		filter["durationyears"] = criteria["years"]
	}

	var policies []domain.Policy

	cursor, err := ps.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var policy domain.Policy
		if err := cursor.Decode(&policy); err != nil {
			log.Println("Error decoding policy:", err)
			continue
		}
		policies = append(policies, policy)
	}

	if len(policies) == 0 {
		return nil, nil
	}

	return policies, nil
}
