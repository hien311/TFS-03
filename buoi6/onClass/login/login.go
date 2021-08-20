package login

import (
	"fmt"
	"net/http"
)

var tokens = []string{"12", "123", "12341"}

func CheckTokenMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if !Include(tokens, token) {

			fmt.Fprintln(w, "token invalid")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Include(s []string, a string) bool {
	for _, v := range s {
		if v == a {
			return true
		}
	}
	return false
}
