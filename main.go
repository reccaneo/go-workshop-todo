package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"todo/todo"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/auth/", func(c *gin.Context) {
		mySigningKey := []byte("passowrd")
		claims := &jwt.StandardClaims{
			Issuer:    "test",
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"token": ss,
		})

	})
	authMiddleware := authMiddleware()
	r.PUT("/todos/", todo.AddTask)

	r.GET("/todos/", authMiddleware, todo.GetTask)

	r.PUT("/todos/:id", todo.DoneTask)

	r.Run(":9090")
}

func authMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Token from another example.  This token is expired
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
		mySigningKey := []byte("passowrd")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpect signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

	}
}
