package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func createUser(c echo.Context) error {
	newUser := new(User)
	if newUser.Id == uuid.Nil {
		newUser.Id = uuid.New()
	}
	c.Bind(newUser)
	key := fmt.Sprintf("User:%s", newUser.Id)
	rh.JSONSet(key, ".", newUser)
	fmt.Println(newUser)
	return c.JSON(http.StatusCreated, newUser.Id)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	key := fmt.Sprintf("User:%s", id)
	userJSON, err := redis.Bytes(rh.JSONGet(key, "."))
	user := new(User)
	json.Unmarshal(userJSON, &user)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, &user)
}
