package route

import (
	"github.com/golang191119/nc_student/handler"
	"github.com/labstack/echo/v4"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {

}

func Staff(e *echo.Echo) {
	g := e.Group("/api/student/v1/staff")
	g.POST("/student", handler.AddStudent)
	g.PUT("/student", handler.UpdateStudent)
	g.DELETE("/student/id/:id", handler.DeleteStudentById)
	g.DELETE("/student", handler.DeleteStudentByIdReq)
}

func Public(e *echo.Echo) {
	g := e.Group("/api/student/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/test", handler.TestDB)
	g.GET("/student", handler.GetAllStudents)
	g.GET("/student/group/last_name", handler.GroupLastName)
	g.GET("/student/id/:id", handler.GetStudentById)
	g.PATCH("/student", handler.SearchStudent)
	g.PATCH("/student/simple/", handler.SearchStudentSimple)

}
