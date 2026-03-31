package repository

import (
	"context"
	"time"

	"github.com/Jnmendza/uh/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrganizationRepo struct {
	collection *mongo.Collection
}

func NewOrganizationRepo(client *mongo.Client) *OrganizationRepo {
	return &OrganizationRepo{
		collection: client.Database("unionhub").Collection("organizations"),
	}
}

func (r *OrganizationRepo) Create(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	// Initialize defaults
	if org.ID.IsZero() {
		org.ID = primitive.NewObjectID()
	}
	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (r *OrganizationRepo) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Organization, error) {
	var org models.Organization
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&org)
	if err != nil {
		return nil, err
	}
	return &org, nil
}

func (r *OrganizationRepo) GetChildren(ctx context.Context, parentID primitive.ObjectID) ([]*models.Organization, error) {
	var orgs []*models.Organization
	cursor, err := r.collection.Find(ctx, bson.M{"parentId": parentID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}
