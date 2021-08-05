package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"todo/todo"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	privateRoute := r.NewRoute().Subrouter()
	privateRoute.Use(authMiddleware)

	r.HandleFunc("/auth/", func(rw http.ResponseWriter, r *http.Request) {
		mySigningKey := []byte("passowrd")
		claims := &jwt.StandardClaims{
			Issuer:    "test",
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(rw).Encode(map[string]string{
			"token": ss,
		})

	}).Methods(http.MethodGet)

	r.HandleFunc("/todos/", todo.AddTask).Methods(http.MethodPut)

	privateRoute.HandleFunc("/todos/", todo.GetTask).Methods(http.MethodGet)

	r.HandleFunc("/todos/{id}", todo.DoneTask).Methods(http.MethodPut)

	http.ListenAndServe(":9090", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Token from another example.  This token is expired
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
		mySigningKey := []byte("passowrd")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpect signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
