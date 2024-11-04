package controller

import (
	"net/http"
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/internal/tokenutil"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *config.Env
}

func NewTokenController(LoginUsecase domain.LoginUsecase,env *config.Env) *LoginController {
	return &LoginController{
		LoginUsecase: LoginUsecase,
		Env:           env,
	}
}

func (lc *LoginController) Login(c *gin.Context) {
	var loginUser domain.AuthLogin
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := lc.LoginUsecase.AuthenticateUser(c, &loginUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tkn := domain.Token{
		ID:           primitive.NewObjectID(),
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Hour * 24 * time.Duration(lc.Env.RefreshTokenExpiryHour)),
		CreatedAt:    time.Now(),
	}
	err = lc.LoginUsecase.SaveRefreshToken(c, &tkn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := domain.LoginResponse{
		ID:           user.ID,
		AcessToken:   accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (lc *LoginController) RefreshTokenHandler(c *gin.Context) {

	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	claims, err := tokenutil.VerifyToken(req.RefreshToken, lc.Env.RefreshTokenSecret)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	_, err = lc.LoginUsecase.CheckRefreshToken(c, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "the user is logged out."})
		return
	}

	user := domain.SignupForm{
		Username: claims.Username,
		Email:    claims.Email,
		ID:   claims.UserID,
	}
	newaccessToken, err := tokenutil.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newaccessToken,
	})
}
