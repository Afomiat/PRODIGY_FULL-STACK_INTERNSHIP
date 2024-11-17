package router

import (
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/delivery/controller"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/delivery/middleware"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/repository"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLogInRouter(env *config.Env, timeout time.Duration, db *mongo.Database, Group *gin.RouterGroup) {
	userR := repository.NewLoginRepo(db, domain.CollectionUser)
	tokenR := repository.NewTokenRepository(db, domain.TokenCollection)

	LoginUsecase := usecase.NewLoginUsecase(userR, tokenR, timeout, env)
	LoginController := controller.NewTokenController(LoginUsecase, env)

	Group.POST("/login", LoginController.Login)
	Group.POST("/refresh", LoginController.RefreshTokenHandler)

	adminGroup := Group.Group("/admin") 
	adminGroup.Use(middleware.RoleRequired(env, domain.AdminRole)) 
	{
		adminGroup.GET("/all_users", )
	}

	employeeGroup := Group.Group("/employee") 
	employeeGroup.Use(middleware.RoleRequired(env, domain.EmployeeRole)) 
	{
		employeeGroup.GET("/dashboard", )
	}
}

