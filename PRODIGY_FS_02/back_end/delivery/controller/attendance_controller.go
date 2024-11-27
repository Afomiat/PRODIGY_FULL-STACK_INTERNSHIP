package controller

import (
	"fmt"
	"net/http"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendanceController struct {
    usecase domain.AttendanceUsecase
}

func NewAttendanceController(usecase domain.AttendanceUsecase) *AttendanceController {
    return &AttendanceController{
        usecase: usecase,
    }
}

func (ctrl *AttendanceController) ClockIn(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
        return
    }

    objectId, err := primitive.ObjectIDFromHex(userID.(string))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
    err = ctrl.usecase.ClockIn(c.Request.Context(), objectId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Clocked in successfully"})
}

func (ctrl *AttendanceController) ClockOut(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
        return
    }

    objectId, err := primitive.ObjectIDFromHex(userID.(string))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
    err = ctrl.usecase.ClockOut(c.Request.Context(), objectId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Clocked out successfully"})
}

func (ctrl *AttendanceController) GetAllAttendanceRecords(c *gin.Context) {
    records, err := ctrl.usecase.GetAllAttendanceRecords(c.Request.Context())
    fmt.Print("records", records)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, records)
}
