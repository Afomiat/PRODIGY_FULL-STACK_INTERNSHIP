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

func NewEmployeeRouter(env *config.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	userR := repository.NewEmployeeRepo(db, domain.CollectionUser)
	empUsecase := usecase.NewEmployeeUsecase(userR, timeout, env)
	empController := controller.NewEmployeeController(empUsecase, env)

	group.POST("/create_employee", empController.CreateUser)
	group.PATCH("/update_employee/:id", empController.UpdateUser)
	group.DELETE("/delete_employee/:id", empController.DeleteUser)
	group.GET("/get_employee/:id", empController.GetUser)
	group.GET("/get_all_employee", empController.GetUsers)

}