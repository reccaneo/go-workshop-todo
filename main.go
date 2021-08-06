package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"todo/repository"
	"todo/todo"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// db.AutoMigrate(&todo.Task{})

	service := todo.NewService(repository.NewMemoryRepository())
	app := todo.NewApp(service)

	//	API should aware to
	//	1.Shutdown graceful
	//	2.Race conditions
	r := gin.Default()
	api := r.Group("/")
	api.Use(authMiddleware)

	r.GET("/auth/", generateAuth)

	r.PUT("/todos/", app.AddTask)

	r.GET("/todos/", app.GetTask)

	api.PUT("/todos/:id", app.DoneTask)

	r.Run(":9090")
}

func authMiddleware(c *gin.Context) {
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
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func generateAuth(c *gin.Context) {
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
}
