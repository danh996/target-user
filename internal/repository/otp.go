package repository

import (
	"context"
	"time"
)

type OTPRepository interface {
	SetOTPKey(ctx context.Context, phone, otp, token string) error
	VerifyOTP(ctx context.Context, phone, otp string) error
	VerifyToken(ctx context.Context, phone, token string) error
	IncrSMSPerPhone(ctx context.Context, phone string, time time.Time) (int64, error)
	DescSMSPerPhone(ctx context.Context, phone string, time time.Time) (int64, error)
}
