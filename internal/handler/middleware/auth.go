package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/alifrahmadian/personal-finance-tracker/internal/config"
	"github.com/alifrahmadian/personal-finance-tracker/internal/errors"
	"github.com/alifrahmadian/personal-finance-tracker/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func RequireAuth(logger *logrus.Logger, authCfg *config.AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetString(RequestIDKey)

		entry := logger.WithFields(logrus.Fields{
			"client_ip":  c.ClientIP(),
			"request_id": requestID,
			"path":       c.Request.URL.Path,
		})

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			entry.Warn("No Authorization header provided")
			response.Error(c, http.StatusUnauthorized, errors.ErrNoAuthHeader.Error(), "")
			c.Abort()

			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			entry.Warn("Invalid token format")
			response.Error(c, http.StatusUnauthorized, errors.ErrInvalidTokenFormat.Error(), "")
			c.Abort()

			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				entry.Warn("Unexpected signing method")
				return nil, errors.ErrUnexpectedSigningMethod
			}

			return []byte(authCfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			entry.WithError(err).Warn("Invalid token")
			response.Error(c, http.StatusUnauthorized, errors.ErrInvalidToken.Error(), "")
			c.Abort()

			return
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			entry.Warn("Token has expired")
			response.Error(c, http.StatusUnauthorized, errors.ErrTokenExpires.Error(), "")
			c.Abort()

			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)

		entry.Infof("Authenticated user ID %d", claims.UserID)

		c.Next()
	}
}

func OptionalAuth(authCfg *config.AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
