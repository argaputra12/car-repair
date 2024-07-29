package middleware

import (
	"net/http"
)

// RoleMiddleware is a middleware to check the role of the admin
func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the user role from the context
		role := r.Context().Value("role").(string)

		// Check if the user role is admin
		if role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
