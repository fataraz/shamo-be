package session

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	usersDomain "shamo-be/internal/domain/users"
	"time"
)

const signingSecret = "minumkopimakankuepancong"
const apiKey = "DDE6733F7EB5157FD3C62972120F9D6727852EFA4058DCA0041F637F5CF70379DA903FE5272E35F561EF4CAA1827695AA233C30D065BD1DD7B18F152088B22E8"

// claims ...
type claims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
type Session struct{}

// CheckAPIKey ...
func (session *Session) CheckAPIKey(ctx echo.Context) {
	apiHeader := ctx.Request().Header.Get("X-Token-Key")
	if apiHeader != apiKey {
		ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "not authorised",
		})
		return
	}
}

// NewBearerToken ...
func NewBearerToken(user *usersDomain.User) (string, time.Time) {
	expiry := time.Now().Add(time.Hour * 24 * 365).Unix()
	claims := &claims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(signingSecret)); err == nil {
		return tokenString, time.Unix(expiry, 0)
	}
	return "", time.Unix(expiry, 0)
}

// RefreshToken ...
func RefreshToken(user *usersDomain.User) (string, time.Time) {
	expiry := time.Now().Add(time.Hour * 24 * 370).Unix()
	claims := &claims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(signingSecret)); err == nil {
		return tokenString, time.Unix(expiry, 0)
	}
	return "", time.Unix(expiry, 0)
}
