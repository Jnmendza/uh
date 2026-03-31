package repository

import (
	"context"
	"time"

	"github.com/Jnmendza/uh/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MembershipRepo struct {
	collection *mongo.Collection
}

func NewMembershipRepo(client *mongo.Client) *MembershipRepo {
	return &MembershipRepo{
		collection: client.Database("unionhub").Collection("memberships"),
	}
}

func (r *MembershipRepo) Create(ctx context.Context, membership *models.Membership) (*models.Membership, error) {
	if membership.ID.IsZero() {
		membership.ID = primitive.NewObjectID()
	}
	membership.JoinedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, membership)
	if err != nil {
		return nil, err
	}
	return membership, nil
}
