package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"../models"
	"../db"
	"strconv"
	"../lib"
)

func GetUser(c echo.Context) error {
	// get params
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	conn := db.ConnectPG()
	defer db.ClosePG(conn)
	user := new(models.Users)
	user.GetUserById(conn,userId)
	if user.Id == 0 {
		c.Logger().Debug("----------- %v ", user)
		return lib.Resp(c, http.StatusNotFound, "not found", "")
	}
	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) (err error) {
	// get params
	conn := db.ConnectPG()
	defer db.ClosePG(conn)
	user := new(models.Users)
	if err = c.Bind(user); err != nil {
		return lib.Resp(c,http.StatusBadRequest,"param err", err)
	}
	user.CreateUser(conn)
	// return user
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	// get params
	userId, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	conn := db.ConnectPG()
	defer db.ClosePG(conn)
	user := new(models.Users)
	user.DeleteUser(conn,userId)
	// return user
	return c.JSON(http.StatusOK, user)
}
