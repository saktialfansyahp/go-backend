package middleware

import (
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/pkg/utils"
)


func AuthHandler(requiredRole string, next http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
		if (tokenString == "") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := utils.ParseToken(tokenString)
		if (err != nil || !token.Valid) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		next.ServeHTTP(w, r)
    }
}