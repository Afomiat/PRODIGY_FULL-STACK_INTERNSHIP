package middleware

import (
    "net/http"
    "strings"

    "github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/config"
)

func RoleRequired(env *config.Env, role domain.Role) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.Request.Header.Get("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        token, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenString, "Bearer "), &domain.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(env.AccessTokenSecret), nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(*domain.JwtCustomClaims); ok && token.Valid {
            userRole := claims.Role
            if !strings.EqualFold(userRole, string(role)) {
                c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
                c.Abort()
                return
            }
            c.Set("claim", claims) 
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        c.Next()
    }
}
