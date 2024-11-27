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

func NewAttendanceRouter(env *config.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
    attendanceRepo := repository.NewAttendanceRepository(db,domain.Attendance)
    attendanceUsecase := usecase.NewAttendanceUsecase(attendanceRepo)
    attendanceController := controller.NewAttendanceController(attendanceUsecase)

    attendanceGroup := group.Group("/attendance")
    attendanceGroup.Use(middleware.AuthMiddleware(env)) // Use authentication middleware
    attendanceGroup.POST("/clock_in", attendanceController.ClockIn)
    attendanceGroup.POST("/clock_out", attendanceController.ClockOut)
    attendanceGroup.GET("/records", attendanceController.GetAllAttendanceRecords)
}
