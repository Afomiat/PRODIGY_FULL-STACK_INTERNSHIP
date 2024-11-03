package router

import (
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *config.Env, timeout time.Duration, db *mongo.Database, r *gin.Engine) {
	PublicRout := r.Group("")
	NewSignUpRouter(env, timeout, db, PublicRout)

}