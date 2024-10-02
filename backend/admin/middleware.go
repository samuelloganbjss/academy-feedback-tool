package admin

import (
	"net/http"
)

func AdminMiddleware(getTutorRole func(*http.Request) (string, error)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, err := getTutorRole(r)
			if err != nil {
				http.Error(w, "Forbidden: Unable to determine user role", http.StatusForbidden)
				return
			}

			if role != "admin" {
				http.Error(w, "Forbidden: Only admins can access this resource", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
