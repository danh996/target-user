package repository

import (
	"context"
	"user/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository ...
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	UpdateInfo(ctx context.Context, userID primitive.ObjectID, user *models.User) error
	UpdatePassword(ctx context.Context, userID string, password string) error
	FindByID(ctx context.Context, userID primitive.ObjectID) (*models.User, error)
	FindByPhone(ctx context.Context, phone string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindAll(ctx context.Context, offset int64, limit int64) ([]*models.User, error)
}
