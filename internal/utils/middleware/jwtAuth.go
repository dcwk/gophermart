package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dcwk/gophermart/internal/utils/auth"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getTokenFromRequest(r)
		userID := auth.GetUserID(token)
		if userID <= 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userKey := "userId"
		ctx := context.WithValue(r.Context(), userKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTokenFromRequest(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
