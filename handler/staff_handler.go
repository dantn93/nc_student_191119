package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang191119/nc_student/db"
	"github.com/labstack/echo/v4"
)

func AddStudent(c echo.Context) error {
	var student db.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	res, err := db.AddStudent(&student)
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateStudent(c echo.Context) error {
	var student db.StudentUpdateRequest
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	fmt.Println(student)
	res, err := db.UpdateStudent(&student)
	if err != nil {
		log.Printf("update error :%v", err)
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteStudentById(c echo.Context) error {
	req := make(map[string]string)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	id, err := strconv.Atoi(req["id"])
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: err.Error()})
	}
	result, err := db.DeleteStudentById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteStudentByIdReq(c echo.Context) error {
	req := struct {
		ID int
	}{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	result, err := db.DeleteStudentByIdReq(req.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, result)
}
