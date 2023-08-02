package interceptor

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shamo-be/internal/shared/models"
	ctxSess "shamo-be/internal/shared/utils/context"
	"strings"
	"time"
)

// Interceptor ...
type Interceptor struct {
}

// claims ...
type claims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// signingSecret ...
const signingSecret = "minumkopimakankuepancong"

// New ...
func New() *Interceptor {
	return &Interceptor{}
}

// ValidateAccess ...
func (i *Interceptor) ValidateAccess() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if i.isHealthCheck(c) {
				return next(c)
			}

			if i.skipCheckSession(c) {
				return next(c)
			}

			if i.checkSession(c) {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusUnauthorized)
		}
	}
}

// isHealthCheck ...
func (i *Interceptor) isHealthCheck(c echo.Context) bool {
	if c.Request().URL.String() == "/" || c.Request().URL.String() == "/ping" {
		return true
	}
	return false
}

// skipCheckSession ...
func (i *Interceptor) skipCheckSession(c echo.Context) bool {
	if strings.HasPrefix(c.Request().URL.String(), "/api/login") ||
		strings.HasPrefix(c.Request().URL.String(), "/api/public") {
		return true
	}
	return false
}

func (i *Interceptor) checkSession(c echo.Context) bool {
	data := c.Get(ctxSess.AppSession)
	ctxSess := data.(*ctxSess.Context)
	authHeader := c.Request().Header.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return false
	} else {
		tokenString := authHeaderParts[1]
		claimsToken := &claims{}
		token, _ := jwt.ParseWithClaims(tokenString, claimsToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(signingSecret), nil
		})

		if token == nil {
			return false
		}

		if token.Valid && !checkTokenExpiry(claimsToken.StandardClaims.ExpiresAt) {
			user := models.AccountSession{
				ID:    claimsToken.ID,
				Email: claimsToken.Email,
				Name:  claimsToken.Name,
			}
			ctxSess.UserSession = user
			return true
		} else {
			return false
		}
	}

	return false
}

// checkTokenExpiry ...
func checkTokenExpiry(timestamp interface{}) bool {
	if validity, ok := timestamp.(int64); ok {
		tm := time.Unix(int64(validity), 0)
		remainder := tm.Sub(time.Now())
		if remainder > 0 {
			return false
		}
	}
	return true
}
