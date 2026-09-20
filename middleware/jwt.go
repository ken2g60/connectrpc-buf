package middleware

import (
	"context"
	"net/http"
	"strings"
)

// Define your public routes here
var publicProcedures = map[string]bool{
	"/greet.v1.GreetService/CreateAccount": true,
	"/greet.v1.GreetService/Login":         true,
	"/greet.v1.GreetService/Greet":         true,
	"/greet.v1.GreetService/Header":        true,
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the route is public
		if publicProcedures[r.URL.Path] {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 1. Validate the JWT (using a library like ://github.com)
		claims, err := validateJWT(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// 2. Inject claims or user details into the context
		ctx := context.WithValue(r.Context(), "user_claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateJWT(tokenString string) (map[string]interface{}, error) {
	// Implement JWT validation logic here (e.g., using a library like github.com/dgrijalva/jwt-go)
	// For demonstration purposes, we'll just return a dummy claim if the token is "valid-token"
	if tokenString == "valid-token" {
		return map[string]interface{}{
			"user_id": "12345",
			"role":    "admin",
		}, nil
	}
	return nil, http.ErrNoCookie // or any other error indicating invalid token
}
