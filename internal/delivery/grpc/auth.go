package grpc

import (
	"context"
	"fmt"
	"time"

	enums "user/common/enum"
	"user/common/errors"
	"user/internal/domain"
	"user/internal/models"
	"user/pb"

	"github.com/danh996/go-shop-kit/token"
	"github.com/danh996/go-shop-kit/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (d *authDelivery) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.ErrInvalidArgument
	}
	u, tkn, err := d.authDomain.Login(ctx, req.GetPhone(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	res := &pb.LoginResponse{
		UserID:        u.ID.Hex(),
		Token:         tkn,
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
	}
	fmt.Println(res)
	return res, nil
}

func (d *authDelivery) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	t := time.Now()
	if err := d.authDomain.Register(ctx, &models.User{
		Fullname:       req.Fullname,
		Password:       req.Password,
		Phone:          req.Phone,
		Role:           enums.User,
		Sex:            enums.UserSex(req.Sex),
		ProvinceID:     req.ProvinceID,
		DistrictID:     req.DistrictID,
		WardID:         req.WardID,
		Job:            req.Job,
		IdentityCard:   req.IdentityCard,
		Age:            req.Age,
		BlockedUserIDS: req.BlockedUserIDs,
		CreatedAt:      &t,
		UpdatedAt:      &t,
		LocationScore:  (utils.GetMasterDataCompare(int(req.ProvinceID), int(req.DistrictID), int(req.WardID))),
	}); err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{}, nil
}

func (d *authDelivery) RegisterCustomer(ctx context.Context, req *pb.RegisterCustomerRequest) (*pb.RegisterCustomerResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}
	t := time.Now()
	if err := d.authDomain.RegisterCustomer(ctx, &models.User{
		Fullname:       req.Fullname,
		Password:       req.Password,
		Email:          req.Email,
		Role:           enums.Customer,
		Sex:            enums.UserSex(req.Sex),
		ProvinceID:     req.ProvinceID,
		DistrictID:     req.DistrictID,
		WardID:         req.WardID,
		Job:            req.Job,
		IdentityCard:   req.IdentityCard,
		Age:            req.Age,
		BlockedUserIDS: req.BlockedUserIDs,
		CreatedAt:      &t,
		UpdatedAt:      &t,
		LocationScore:  (utils.GetMasterDataCompare(int(req.ProvinceID), int(req.DistrictID), int(req.WardID))),
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
		UserID:         u.ID.Hex(),
		Token:          tkn,
		Role:           pb.UserRole(u.Role),
		Fullname:       u.Fullname,
		Phone:          u.Phone,
		Age:            u.Age,
		IdentityCard:   u.IdentityCard,
		Job:            u.Job,
		Sex:            pb.UserSex(u.Sex),
		ProvinceID:     u.ProvinceID,
		DistrictID:     u.DistrictID,
		WardID:         u.WardID,
		LocationScore:  u.LocationScore,
		CreatedAt:      u.CreatedAt.String(),
		UpdatedAt:      u.UpdatedAt.String(),
		BlockedUserIDs: u.BlockedUserIDS,
		Email:          u.Email,
	}
	return res, nil
}

func (d *authDelivery) LoginAdmin(ctx context.Context, req *pb.LoginAdminRequest) (*pb.LoginAdminResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.ErrInvalidArgument
	}
	u, tkn, err := d.authDomain.LoginByEmail(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	res := &pb.LoginAdminResponse{
		UserID:         u.ID.Hex(),
		Token:          tkn,
		Role:           pb.UserRole(u.Role),
		Fullname:       u.Fullname,
		Phone:          u.Phone,
		Age:            u.Age,
		IdentityCard:   u.IdentityCard,
		Job:            u.Job,
		Sex:            pb.UserSex(u.Sex),
		ProvinceID:     u.ProvinceID,
		DistrictID:     u.DistrictID,
		WardID:         u.WardID,
		LocationScore:  u.LocationScore,
		CreatedAt:      u.CreatedAt.String(),
		UpdatedAt:      u.UpdatedAt.String(),
		BlockedUserIDs: u.BlockedUserIDS,
		Email:          u.Email,
	}
	return res, nil
}

// func (d *authDelivery) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, errors.ErrInvalidArgument
// 	}

// 	token, err := d.authDomain.ForgotPassword(ctx, req.Phone)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.ForgotPasswordResponse{
// 		Token: token,
// 	}, nil
// }

// func (d *authDelivery) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, errors.ErrInvalidArgument
// 	}

// 	if err := d.authDomain.VerifyOTP(ctx, req.Phone, req.Otp); err != nil {
// 		return nil, err
// 	}

// 	return &pb.VerifyOTPResponse{}, nil
// }

func (d *authDelivery) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.ErrInvalidArgument
	}

	if err := d.authDomain.ResetPassword(ctx, req.Phone, req.NewPassword, req.Token); err != nil {
		return nil, err
	}

	return &pb.ResetPasswordResponse{}, nil
}

func (d *authDelivery) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	err := d.authDomain.ChangePassword(ctx, req.Phone, req.Password, req.NewPassword)
	if err != nil {
		return nil, err
	}

	return &pb.ChangePasswordResponse{}, nil
}
