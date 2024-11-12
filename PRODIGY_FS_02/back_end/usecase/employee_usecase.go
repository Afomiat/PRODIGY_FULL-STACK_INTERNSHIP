package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/internal/employeeUtil"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	env *config.Env
}

func NewEmployeeUsecase(userRepository domain.UserRepository, timeout time.Duration, env *config.Env) domain.UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		env: env,
	}
}

func (uc *UserUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	return uc.userRepository.GetUserByEmail(ctx, email)
}

func (uc *UserUsecase) GetUserByUsername(c context.Context, username string) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	return uc.userRepository.GetUserByUsername(ctx, username)
}

func (uc *UserUsecase) CreateUser(c context.Context, user *domain.SignupForm, claims *domain.JwtCustomClaims) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	aduser := domain.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		Role:     "user",
	}
	err := employeeUtil.CanManipulateUser(claims, &aduser, "add")
	if err != nil {
		return errors.New(err.Message)
	}
	if !employeeUtil.ValidateEmail(user.Email) {
		return errors.New("invalid email")
	}
	if !employeeUtil.ValidatePassword(user.Password) {
		return errors.New("password must be at least 8 characters long")
	}
	return uc.userRepository.CreateUser(ctx, &aduser)
}
func (uc *UserUsecase) GetUserByID(c context.Context, id primitive.ObjectID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	return uc.userRepository.GetUserByID(ctx, id)
}


func (uc *UserUsecase) UpdateUser(c context.Context, user *domain.User, claims *domain.JwtCustomClaims, existinguser *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	user.ID = existinguser.ID
	_err := employeeUtil.CanManipulateUser(claims, user, "update")
	if _err != nil {
		return nil, errors.New(_err.Message)
	}
	if user.Email == ""{
		user.Email = existinguser.Email
	}
	if user.Username == ""{
		user.Username = existinguser.Username
	}
	if user.Password == ""{
		user.Password = existinguser.Password
	}
	user.Password,_ = employeeUtil.HashPassword(user.Password)
	if user.Role == ""{
		user.Role = existinguser.Role
	}
	if user.First_Name == ""{
		user.First_Name = existinguser.First_Name
	}
	if user.Last_Name == ""{
		user.Last_Name = existinguser.Last_Name
	}
	if user.Bio == ""{
		user.Bio = existinguser.Bio
	}
	if user.Profile_Picture == ""{
		user.Profile_Picture = existinguser.Profile_Picture
	}
	if len(user.Contact_Info) == 0{
		user.Contact_Info = existinguser.Contact_Info
	}
	if !employeeUtil.ValidateEmail(user.Email) {
		return nil, errors.New("invalid email")
	}
	if !employeeUtil.ValidatePassword(user.Password) {
		return nil, errors.New("password must be at least 8 characters long")
	}
	err := uc.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) DeleteUser(c context.Context, id primitive.ObjectID, claims *domain.JwtCustomClaims) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	user := domain.User{
		ID: id,
	}
	_err := employeeUtil.CanManipulateUser(claims, &user, "update")
	if _err != nil {
		return errors.New(_err.Message)
	}
	return uc.userRepository.DeleteUser(ctx, id)
}

func (uc *UserUsecase) GetAllUsers(c context.Context) ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.GetAllUsers(ctx)
}