package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/internal/employeeUtil"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/internal/tokenutil"
)

type loginUsecase struct {
	userRepository  domain.UserRepository
	tokenRepository domain.TokenRepository
	contextTimeout  time.Duration
	env             *config.Env
}

func NewLoginUsecase(loginRepo domain.UserRepository, tokenRepo domain.TokenRepository, timeout time.Duration, env *config.Env) *loginUsecase {
	return &loginUsecase{
		userRepository:     loginRepo,
		contextTimeout:     timeout,
		tokenRepository:    tokenRepo,
		env:                env,
	}
}

func (lu *loginUsecase) AuthenticateUser(c context.Context, login *domain.AuthLogin) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	user, err := lu.userRepository.GetUserByEmail(ctx, login.Email)
	if err != nil {
		return nil, errors.New("invalid Email")
	}

	err = employeeUtil.ComparePassword(user.Password, login.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateAccessToken(&domain.SignupForm{Email: user.Email, Username: user.Username, ID: user.ID, Role: string(user.Role)}, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateRefreshToken(&domain.SignupForm{Username: user.Username, Email: user.Email, ID: user.ID,Role: string(user.Role)}, secret, expiry)
}

func (lu *loginUsecase) SaveRefreshToken(c context.Context, token *domain.Token) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	return lu.tokenRepository.SaveToken(ctx, token)
}

func (lu *loginUsecase) CheckRefreshToken(c context.Context, refreshToken string) (*domain.Token, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	return lu.tokenRepository.FindTokenByRefreshToken(ctx, refreshToken)

}