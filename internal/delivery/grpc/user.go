package grpc

import (
	"context"
	"user/internal/domain"
	"user/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userDelivery struct {
	userDomain domain.UserDomain
	pb.UnimplementedUserServiceServer
}

// NewUserDelivery ...
func NewUserDelivery(userDomain domain.UserDomain) pb.UserServiceServer {
	return &userDelivery{
		userDomain: userDomain,
	}
}

func (d *userDelivery) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return nil, err
	}
	u, err := d.userDomain.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIDResponse{
		UserID:        u.ID.Hex(),
		Role:          pb.UserRole(u.Role),
		Fullname:      u.Fullname,
		Phone:         u.Phone,
		Age:           u.Age,
		IdentityCard:  u.IdentityCard,
		Job:           u.Job,
		Sex:           pb.UserSex(u.Sex),
		ProvinceID:    u.ProvinceID,
		DistrictID:    u.DistrictID,
		WardID:        u.WardID,
		LocationScore: u.LocationScore,
		// CreatedAt:      u.CreatedAt.String(),
		// UpdatedAt:      u.UpdatedAt.String(),
		BlockedUserIDs: u.BlockedUserIDS,
		Email:          u.Email,
	}, nil
}
