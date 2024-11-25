package router

import (
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/delivery/middleware"
	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Env, timeout time.Duration, db *mongo.Database, r *gin.Engine) {
	PublicRout := r.Group("")
	NewSignUpRouter(env, timeout, db, PublicRout)
	NewLogInRouter(env, timeout, db, PublicRout)

	protectedRouter := r.Group("")
	protectedRouter.Use(middleware.RoleRequired(env, domain.AdminRole))
	NewEmployeeRouter(env, timeout, db, protectedRouter)


}