package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Chien0903/Go-ToDo-App/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const UsernameKey contextKey = "username"

// JWTAuthMiddleware xác thực JWT token và lưu thông tin user vào context
func JWTAuthMiddleware(cfg config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Lấy token từ header Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error": "Authorization header required"}`))
				return
			}

			// Kiểm tra format "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error": "Invalid authorization header format. Use: Bearer <token>"}`))
				return
			}

			tokenString := parts[1]

			// Parse và validate token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Kiểm tra signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(cfg.JWTSecret), nil
			})

			if err != nil || !token.Valid {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error": "Invalid or expired token"}`))
				return
			}

			// Lấy claims từ token
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error": "Invalid token claims"}`))
				return
			}

			// Lưu thông tin user vào context
			userID, ok := claims["user_id"].(float64)
			if !ok {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error": "Invalid user ID in token"}`))
				return
			}

			username, ok := claims["username"].(string)
			if !ok {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error": "Invalid username in token"}`))
				return
			}

			// Thêm thông tin vào context
			ctx := context.WithValue(r.Context(), UserIDKey, uint(userID))
			ctx = context.WithValue(ctx, UsernameKey, username)

			// Tiếp tục với request đã có context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserID lấy user ID từ context (dùng trong handlers)
func GetUserID(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(UserIDKey).(uint)
	return userID, ok
}

// GetUsername lấy username từ context (dùng trong handlers)
func GetUsername(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(UsernameKey).(string)
	return username, ok
}
