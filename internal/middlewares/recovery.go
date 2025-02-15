package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, err := fmt.Fprint(w, err)
				if err != nil {
					log.Fatal(err)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
