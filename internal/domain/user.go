package domain

import (
	"context"
	"user/internal/models"
	"user/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDomain ...
type UserDomain interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID primitive.ObjectID) (*models.User, error)
	GetListUser(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, userID string, user *models.User) error
	DeleteUser(ctx context.Context, userID string) error
}

type userDomain struct {
	userRepository repository.UserRepository
}

// NewUserDomain ...
func NewUserDomain(userRepository repository.UserRepository) UserDomain {
	return &userDomain{
		userRepository,
	}
}

func (d *userDomain) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (d *userDomain) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}

func (d *userDomain) GetUserByID(ctx context.Context, userId primitive.ObjectID) (*models.User, error) {
	user, err := d.userRepository.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *userDomain) GetListUser(ctx context.Context) ([]*models.User, error) {
	return []*models.User{}, nil
}

func (d *userDomain) UpdateUser(ctx context.Context, userID string, user *models.User) error {
	return nil
}

func (d *userDomain) DeleteUser(ctx context.Context, userID string) error {
	return nil
}
