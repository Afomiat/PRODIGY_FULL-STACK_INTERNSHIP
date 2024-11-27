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
    "github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/delivery/middleware"
)

func NewEmployeeRouter(env *config.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
    userR := repository.NewEmployeeRepo(db, domain.CollectionUser)
    empUsecase := usecase.NewEmployeeUsecase(userR, timeout, env)
    empController := controller.NewEmployeeController(empUsecase, env)

    employeeGroup := group.Group("/employees")
    employeeGroup.Use(middleware.RoleRequired(env, domain.AdminRole)) 

    employeeGroup.POST("/create_employee", empController.CreateUser)
    employeeGroup.PUT("/update_employee/:id", empController.UpdateUser)
    employeeGroup.DELETE("/delete_employee/:id", empController.DeleteUser)
    employeeGroup.GET("/get_employee/:id", empController.GetUser)
    employeeGroup.GET("/get_all_employee", empController.GetUsers)
}
