package middlewares

import (
	"api/controllers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RejectNonJsonRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Accept"][0] != "application/json" {
			w.WriteHeader(http.StatusNotAcceptable)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s (%s)", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func JsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusInternalServerError) //500
				err := controllers.ErrorMessage{Message:fmt.Sprint(r)}
				json.NewEncoder(w).Encode(err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
