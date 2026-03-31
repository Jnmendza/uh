package repository

import (
	"context"

	"github.com/Jnmendza/uh/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MembershipTierRepo struct {
	collection *mongo.Collection
}

func NewMembershipTierRepo(client *mongo.Client) *MembershipTierRepo {
	return &MembershipTierRepo{
		collection: client.Database("unionhub").Collection("membership_tiers"),
	}
}

func (r *MembershipTierRepo) Create(ctx context.Context, tier *models.MembershipTier) (*models.MembershipTier, error) {
	if tier.ID.IsZero() {
		tier.ID = primitive.NewObjectID()
	}

	_, err := r.collection.InsertOne(ctx, tier)
	if err != nil {
		return nil, err
	}
	return tier, nil
}
