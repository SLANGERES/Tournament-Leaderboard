package middleware

import (
	"context"
	"net/http"

	jwt_token "github.com/SLANGERES/Tournament-Lederboard/internal/common/jwt"
)

type contextKey string

const ClaimsKey contextKey = "Claims"

func AdminMiddlewareAdmin(jwtMaker *jwt_token.JwtMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("admin-access-token")
			if err != nil {
				http.Error(w, "Unauthorized: Missing token cookie", http.StatusUnauthorized)
				return
			}
			accessToken := cookie.Value
			claims, err := jwtMaker.VerifyToken(accessToken)
			if err != nil {
				http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
				return
			}
			if claims.Role != "admin" {
				http.Error(w, "Unauthorized: Admin only", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), ClaimsKey, &claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func UserMiddlewareAdmin(jwtMaker *jwt_token.JwtMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("admin-access-token")
			if err != nil {
				http.Error(w, "Unauthorized: Missing token cookie", http.StatusUnauthorized)
				return
			}
			accessToken := cookie.Value
			claims, err := jwtMaker.VerifyToken(accessToken)
			if err != nil {
				http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
				return
			}
			if claims.Role != "user" {
				http.Error(w, "Unauthorized: You are not a user", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), ClaimsKey, &claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func ClaimsJWT(ctx context.Context) (*jwt_token.TournamentClaims, bool) {

	claims, ok := ctx.Value(ClaimsKey).(*jwt_token.TournamentClaims)

	return claims, ok
}
