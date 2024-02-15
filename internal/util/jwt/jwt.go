package jwt

import (
	"time"

	"github.com/Karanth1r3/grpc_learn/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

// NewToken generates token for user
func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// Wrapping info to token
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix() // Expiration
	claims["app_id"] = app.ID

	// Signing token
	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
