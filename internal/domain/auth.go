package domain

import (
	"context"
	"user/common/errors"
	"user/internal/models"
	"user/internal/repository"

	"github.com/danh996/go-shop-kit/requestinfo"
	"github.com/danh996/go-shop-kit/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthDomain ...
type AuthDomain interface {
	Login(ctx context.Context, phone, password string) (*models.User, string, error)
	LoginByEmail(ctx context.Context, email, password string) (*models.User, string, error)
	Register(ctx context.Context, user *models.User) error
	RegisterCustomer(ctx context.Context, user *models.User) error
	ChangePassword(ctx context.Context, phone, password, newPassword string) error
	ResetPassword(ctx context.Context, phone, newPassword, token string) error
	//ForgotPassword(ctx context.Context, phone string) (string, error)
	//VerifyOTP(ctx context.Context, phone, otp string) error
}

type authDomain struct {
	userRepository repository.UserRepository
	authenticator  token.Authenticator
}

// NewAuthDomain ...
func NewAuthDomain(
	userRepository repository.UserRepository,
	authenticator token.Authenticator,

) AuthDomain {
	return &authDomain{
		userRepository: userRepository,
		authenticator:  authenticator,
	}
}

func (d *authDomain) Login(ctx context.Context, phone, password string) (*models.User, string, error) {
	user, err := d.userRepository.FindByPhone(ctx, phone)
	if err != nil {
		return nil, "", status.Errorf(codes.InvalidArgument, "find by phone %w", err)
	}

	if !user.IsCorrectPassword(password) {
		return nil, "", status.Errorf(codes.InvalidArgument, "verify password %w", errors.ErrPasswordIsNotCorrect)
	}
	tkn, err := d.authenticator.Generate(&requestinfo.Info{
		UserID: user.ID.Hex(),
	})

	if err != nil {
		return nil, "", status.Errorf(codes.Internal, "gen token %w", err)
	}
	return user, tkn.Token, nil
}

func (d *authDomain) LoginByEmail(ctx context.Context, email, password string) (*models.User, string, error) {

	user, err := d.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", status.Errorf(codes.InvalidArgument, "find by email %w", err)
	}
	if !user.IsCorrectPassword(password) {
		return nil, "", status.Errorf(codes.InvalidArgument, "verify password %w", errors.ErrPasswordIsNotCorrect)
	}

	tkn, err := d.authenticator.Generate(&requestinfo.Info{
		UserID: user.ID.Hex(),
	})

	if err != nil {
		return nil, "", status.Errorf(codes.Internal, "gen token %w", err)
	}

	return user, tkn.Token, nil
}

func (d *authDomain) Register(ctx context.Context, user *models.User) error {
	_, err := d.userRepository.FindByPhone(ctx, user.Phone)
	if err != nil && err != errors.ErrPhoneNotFound {
		return err
	}

	if err == nil {
		return status.Errorf(codes.AlreadyExists, "find by phone: %w", errors.ErrPhoneAlreadyExists)
	}

	if err := d.userRepository.Create(ctx, user); err != nil {
		return status.Errorf(codes.Internal, "create user %w", err)
	}

	return nil
}

func (d *authDomain) RegisterCustomer(ctx context.Context, user *models.User) error {
	_, err := d.userRepository.FindByEmail(ctx, user.Email)
	if err != nil && err != errors.ErrEmailNotFound {
		return status.Errorf(codes.InvalidArgument, "find by email %w", err)
	}

	if err == nil {
		return status.Errorf(codes.AlreadyExists, "find by email: %w", errors.ErrEmailAlreadyExists)
	}

	if err := d.userRepository.Create(ctx, user); err != nil {
		return status.Errorf(codes.Internal, "create user %w", err)
	}

	return nil
}

func (d *authDomain) ChangePassword(ctx context.Context, phone, password, newPassword string) error {
	u, err := d.userRepository.FindByPhone(ctx, phone)
	if err != nil {
		return status.Errorf(codes.NotFound, "find by phone %w", err)
	}

	if !u.IsCorrectPassword(password) {
		return status.Errorf(codes.InvalidArgument, "verify password %w", errors.ErrPasswordIsNotCorrect)
	}

	u.Password = newPassword
	u.HashPassword()
	if err := d.userRepository.UpdateInfo(ctx, u.ID, u); err != nil {
		return status.Errorf(codes.Internal, "update user info %w", err)
	}
	return nil
}

func (d *authDomain) ResetPassword(ctx context.Context, phone, newPassword, token string) error {
	// if err := d.otpRepository.VerifyToken(ctx, phone, token); err != nil {
	// 	return status.Errorf(codes.InvalidArgument, "verify token %w", err)
	// }

	// u, err := d.userRepository.FindByPhone(ctx, phone)
	// if err != nil {
	// 	return status.Errorf(codes.NotFound, "find by phone %w", err)
	// }

	// u.Password = newPassword
	// u.HashPassword()
	// if err := d.userRepository.UpdateInfo(ctx, u.ID, u); err != nil {
	// 	return status.Errorf(codes.Internal, "update user info %w", err)
	// }
	return nil
}

// func (d *authDomain) ForgotPassword(ctx context.Context, phone string) (string, error) {
// 	res, err := d.otpRepository.IncrSMSPerPhone(ctx, phone, time.Now())
// 	if err != nil {
// 		return "", status.Errorf(codes.Internal, "increase sms per phone %w", err)
// 	}

// 	if res > int64(d.maxSMSPerDay) {
// 		return "", status.Errorf(codes.OutOfRange, "exceed sms today")
// 	}

// 	token := utils.RandStringBytes(d.tokenLength)
// 	otp := utils.GenOTPCode(d.otpLength)
// 	body := &pb.SMSNotification{
// 		Body: fmt.Sprintf("Day la ma OTP cua ban: %v", otp),
// 		To:   phone,
// 	}
// 	data, err := proto.Marshal(body)
// 	if err != nil {
// 		return "", err
// 	}

// 	if err := d.postNotificationPublisher.Publish(
// 		&kafka_utils.Topic{
// 			Name: d.postNotificationTopic,
// 		},
// 		nil,
// 		data,
// 	); err != nil {
// 		d.otpRepository.DescSMSPerPhone(ctx, phone, time.Now())
// 	}

// 	if err := d.otpRepository.SetOTPKey(ctx, phone, otp, token); err != nil {
// 		return "", status.Errorf(codes.Internal, "set otp key %w", err)
// 	}

// 	return token, nil
// }

// func (d *authDomain) VerifyOTP(ctx context.Context, phone, otp string) error {
// 	if err := d.otpRepository.VerifyOTP(ctx, phone, otp); err != nil {
// 		return status.Errorf(codes.InvalidArgument, "verify otp %w", err)
// 	}
// 	return nil
// }
