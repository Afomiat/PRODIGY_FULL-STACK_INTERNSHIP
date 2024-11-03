package controller

import (
	"fmt"
	"net/http"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	env           *config.Env
}

func NewSignupController(signupUsecase domain.SignupUsecase, env *config.Env) *SignupController {
	return &SignupController{
		SignupUsecase: signupUsecase,
		env:           env,
	}
}

func (sc *SignupController) Signup(ctx *gin.Context) {
    var user domain.SignupForm

    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	fmt.Println("cheking user name******************",user.Username)
    returnUser, _ := sc.SignupUsecase.GetUserByUserName(ctx, user.Username)
    if returnUser != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username Already Exists!"})
		fmt.Println("Username Already Exists!3******************",user.Username)
        return
    }

    returnUser, _ = sc.SignupUsecase.GetUserByEmail(ctx, user.Email)
    if returnUser != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email Already Exists!"})
        return
    }

    // Send OTP
    err := sc.SignupUsecase.SendOtp(ctx, &user, sc.env.SMTPUsername, sc.env.SMTPPassword)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "OTP sending failed"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "OTP Sent"})
}

