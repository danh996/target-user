package grpc

import (
	"context"
	"time"

	"user/common/errors"
	"user/internal/domain"
	"user/internal/models"
	"user/pb"

	"github.com/danh996/go-shop-kit/token"
)

type authDelivery struct {
	authDomain    domain.AuthDomain
	authenticator token.Authenticator
	pb.UnimplementedAuthServiceServer
}

// NewAuthDelivery ...
func NewAuthDelivery(authDomain domain.AuthDomain, authenticator token.Authenticator) pb.AuthServiceServer {
	return &authDelivery{
		authDomain:    authDomain,
		authenticator: authenticator,
	}
}

func (d *authDelivery) RegisterCustomer(ctx context.Context, req *pb.RegisterCustomerRequest) (*pb.RegisterCustomerResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}
	t := time.Now()
	if err := d.authDomain.RegisterCustomer(ctx, &models.User{
		Fullname: req.Fullname,
		Password: req.Password,
		Email:    req.Email,

		CreatedAt: &t,
		UpdatedAt: &t,
	}); err != nil {
		return nil, err
	}
	return &pb.RegisterCustomerResponse{}, nil
}

func (d *authDelivery) LoginCustomer(ctx context.Context, req *pb.LoginCustomerRequest) (*pb.LoginCustomerResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.ErrInvalidArgument
	}
	u, tkn, err := d.authDomain.LoginByEmail(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	res := &pb.LoginCustomerResponse{
		Token:    tkn,
		Fullname: u.Fullname,

		Email: u.Email,
	}
	return res, nil
}
