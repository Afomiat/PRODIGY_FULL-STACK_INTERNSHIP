package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupUsecase interface {
	VerifyOtp(ctx context.Context, otp *VerifyOtp)(*OTP, error)
	RegisterUser(ctx context.Context, user *SignupForm) (*primitive.ObjectID, error)
	GetUserByUserName(ctx context.Context, username string) (*SignupForm, error)
	GetUserByEmail(ctx context.Context, Email string) (*SignupForm, error)
	SendOtp(cxt context.Context, user *SignupForm,  stmpName, stmpPass string) error

}

type SignupRepository interface {
	CreateUser(ctx context.Context, user *SignupForm) error
	GetUserByUserName(ctx context.Context, username string) (*SignupForm, error)
	GetUserByEmail(ctx context.Context, Email string) (*SignupForm, error)
}

type OtpRepository interface{
	GetOtpByEmail(ctx context.Context, email string) (*OTP, error)
	DeleteOTP(ctx context.Context, email string) error
	SaveOTP(ctx context.Context, otp *OTP) error
}