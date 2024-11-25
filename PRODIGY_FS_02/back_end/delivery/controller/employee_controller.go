package controller

import (
	"fmt"
	"net/http"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeController struct {
    EmployeeUsecase domain.UserUsecase
    Env          *config.Env
}

func NewEmployeeController(EmployeeUsecase domain.UserUsecase, env *config.Env) *EmployeeController {
    return &EmployeeController{
        EmployeeUsecase: EmployeeUsecase,
        Env:           env,
    }
}

func (uc *EmployeeController) CreateUser(c *gin.Context) {
    claims := c.MustGet("claim").(*domain.JwtCustomClaims)
    var user domain.SignupForm

    if err := c.ShouldBindJSON(&user); err != nil {
        fmt.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Received user data:", user)

    returnedUser, _ := uc.EmployeeUsecase.GetUserByEmail(c, user.Email)
    if returnedUser != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
        return
    }
    returnedUser, _ = uc.EmployeeUsecase.GetUserByUsername(c, user.Username)
    if returnedUser != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }

    err := uc.EmployeeUsecase.CreateUser(c, &user, claims)
    if err != nil {
        fmt.Println("Error creating user:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("User created successfully")
    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}



func (uc *EmployeeController) UpdateUser(c *gin.Context) {
	claims := c.MustGet("claim").(*domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	existingUser, _ := uc.EmployeeUsecase.GetUserByID(c, objectID)
	if existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Email != existingUser.Email {
		print(user.Email)
		euser, _ := uc.EmployeeUsecase.GetUserByEmail(c, user.Email)
		if euser != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
	}
	if user.Username != existingUser.Username {
		euser,_ := uc.EmployeeUsecase.GetUserByUsername(c,user.Username)
		if euser!=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
	}
	resp, err := uc.EmployeeUsecase.UpdateUser(c, &user, claims, existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": resp})
}

func (uc *EmployeeController) DeleteUser(c *gin.Context) {
    claims := c.MustGet("claim").(*domain.JwtCustomClaims)
    id := c.Param("id")
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    existingUser, err := uc.EmployeeUsecase.GetUserByID(c, objectID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if existingUser == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
        return
    }

    err = uc.EmployeeUsecase.DeleteUser(c, objectID, claims)
    if err != nil {
        fmt.Println("Error deleting user:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("User deleted successfully")
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uc *EmployeeController) GetUser(c *gin.Context) {
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	user, _ := uc.EmployeeUsecase.GetUserByID(c, objectID)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *EmployeeController) GetUsers(c *gin.Context) {
	users, err := uc.EmployeeUsecase.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No users found"})
		return
	}
	c.JSON(http.StatusOK, users)
}
