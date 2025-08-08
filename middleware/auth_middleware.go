package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("some_value")

type Claims struct {
	UserName string `json:username`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer")

			claims := &Claims{}

			token, err := jwt.ParseWithClaims(
				tokenString,
				claims,
				func(t *jwt.Token) (any, error) {
					return jwtKey, nil
				})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "username", claims.UserName)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
