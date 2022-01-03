package mongo

import (
	"context"
	"log"

	"user/common/errors"
	"user/internal/models"
	"user/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	coll *mongo.Collection
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	err := user.HashPassword()
	if err != nil {
		return err
	}
	_, err = r.coll.InsertOne(ctx, user)
	return err
}

func (r *userRepository) UpdateInfo(ctx context.Context, userID primitive.ObjectID, user *models.User) error {
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{
		"$set": user,
	})
	return err
}

func (r *userRepository) UpdatePassword(ctx context.Context, userID, password string) error {
	query := bson.M{
		"$set": bson.M{"password": password},
	}
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": userID}, query)
	return err
}

func (r *userRepository) FindByID(ctx context.Context, userID primitive.ObjectID) (*models.User, error) {
	result := &models.User{}
	err := r.coll.FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) FindByPhone(ctx context.Context, phone string) (*models.User, error) {
	result := &models.User{}
	err := r.coll.FindOne(ctx, bson.M{"phone": phone}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, errors.ErrPhoneNotFound
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	result := &models.User{}
	err := r.coll.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, errors.ErrEmailNotFound
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) FindAll(ctx context.Context, offset int64, limit int64) ([]*models.User, error) {
	results := []*models.User{}
	cur, err := r.coll.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result *models.User
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

// NewUserRepository ...
func NewUserRepository(coll *mongo.Collection) repository.UserRepository {
	return &userRepository{
		coll: coll,
	}
}
