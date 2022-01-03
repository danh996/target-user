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

		Fullname: u.Fullname,

		Email: u.Email,
	}, nil
}
