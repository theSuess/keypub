package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	logf "github.com/theSuess/keypub/pkg/log"
	"github.com/theSuess/keypub/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

var log = logf.Log.WithName("auth")

var userCtxKey = &contextKey{"user"}

// TODO: Provide a way to change this
const jwtSecret = "LlzRluVCsOPPWY9OWhFmtplBowgKImmh"

type contextKey struct {
	name string
}

type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticationClaim struct {
	UserID string
	Roles  []string
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.TrimPrefix(c.GetHeader("Authentication"), "Bearer ")
		if token != "" {
			token, err := jwt.ParseString(token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "authentication token")})
				return
			}
			claim := &AuthenticationClaim{UserID: token.Subject()}
			log.Info("authenticated user", "userId", claim.UserID)
			ctx := context.WithValue(c.Request.Context(), userCtxKey, claim)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}
		c.Next()
	}
}

func Authenticate(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Actually implement propper auth
		req := AuthenticationRequest{}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username must be specified"})
			return
		}
		user := &model.User{}
		if err := db.Where(&model.User{Username: req.Username}).First(user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Info("issuing token for user", "username", user.Username, "userId", user.ID)
		token := jwt.New()
		token.Set(jwt.SubjectKey, user.ID)
		token.Set(jwt.IssuedAtKey, time.Now())
		signed, err := jwt.Sign(token, jwa.HS512, []byte("foobar"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": string(signed)})
	}
}

func ForContext(ctx context.Context) *AuthenticationClaim {
	raw, _ := ctx.Value(userCtxKey).(*AuthenticationClaim)
	return raw
}
