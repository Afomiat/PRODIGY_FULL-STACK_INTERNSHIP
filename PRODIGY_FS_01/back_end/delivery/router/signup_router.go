package router

import (
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/repository"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/usecase"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/delivery/controller"

)

func NewSignUpRouter(env *config.Env, timeout time.Duration, db *mongo.Database, Group *gin.RouterGroup) {
	userR := repository.NewSignupRepo(db, domain.UserCollection)
	otpR := repository.NewOtpRepository(db, domain.OtpCollection)

	signUsecase := usecase.NewSignupUsecase(userR, otpR, timeout, env)
	signController := controller.NewSignupController(signUsecase, env)

	Group.POST("/signup", signController.Signup)
	// Group.POST("/verify", signController.Verify)

}