package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("http://localhost:4000")
	ctx := context.Background()
	rh := rejson.NewReJSONHandler()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6397",
		DB:   0,
	})
	rh.SetGoRedisClientWithContext(ctx, rdb)

	router := echo.New()
	router.Use(middleware.CORS())
	router.POST("/users", createUser)
	router.GET("/users/:id", getUser)
	router.Start(":4000")
}
