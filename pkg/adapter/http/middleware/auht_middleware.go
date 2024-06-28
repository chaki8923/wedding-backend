package middleware

import (
	"fmt"
	"strings"
	"log"
	"golang.org/x/xerrors"

	"github.com/chaki8923/wedding-backend/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type Auth interface {
	AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type AuthMiddleware struct {
	AuthUseCase usecase.Auth
}

func NewAuthMiddleware(ju usecase.Auth) Auth {
	AuthMiddleware := AuthMiddleware{
		AuthUseCase: ju,
	}
	return &AuthMiddleware
}

func (j *AuthMiddleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if j.isSkippedPath(c.Request().URL.Path, c.Request().Referer()) {
			if err := next(c); err != nil {
				return xerrors.Errorf("AuthMiddleware error path: %s: %w", c.Request().URL.Path, err)
			}
			return nil
		}
		
		// Log all request headers for debugging
		for name, values := range c.Request().Header {
			for _, value := range values {
				log.Printf("Header: %s=%s", name, value)
			}
		}
		
		cookie, err := c.Cookie("token")
		if err != nil {
			return xerrors.Errorf("AuthMiddleware！！ not extract cookie: %w", err)
		}

		claims, err := j.AuthUseCase.JwtParser(cookie.Value)
		if err != nil {
			j.AuthUseCase.DeleteCookie(c, cookie)
			return fmt.Errorf("failed to parse jwt claims: %w", err)
		}

		log.Printf("新規クライム作成")
		var cl = *claims
		uId := cl["user_id"].(string)
		if err := j.AuthUseCase.IdentifyJwtUser(uId); err != nil {
			j.AuthUseCase.DeleteCookie(c, cookie)
			return fmt.Errorf("failed to personal authentication: %w", err)
		}

		log.Printf("クライム作成後")
		if err := next(c); err != nil {
			return xerrors.Errorf("failed to AuthMiddleware err: %w", err)
		}
		log.Printf("auth middlewareは通過")
		return nil
	}
}

func (j *AuthMiddleware) isSkippedPath(reqPath, refPath string) bool {
	skippedPaths := []string{"/healthcheck", "/csrf-cookie", "/login","/signup", "/logout", "/playground"}
	for _, path := range skippedPaths {
		if strings.Contains(reqPath, path) || strings.Contains(refPath, path) {
			return true
		}
	}

	return false
}
