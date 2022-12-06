package routers

import (
	"retake_stat/internal/entity"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.GET("/group/:id", entity.GetGroup)
	route.POST("/group", entity.AddGroup)
	route.DELETE("/group", entity.DelGroup)

	route.GET("/subject/:id", entity.GetSubject)
	route.POST("/subject", entity.AddSubject)
	route.DELETE("/subject", entity.DelSubject)

	route.GET("/teacher/:id", entity.GetTeacher)
	route.POST("/teacher", entity.AddTeacher)
	route.DELETE("/teacher", entity.DelTeacher)

	route.GET("/student/:id", entity.GetStudent)
	route.POST("/student", entity.AddStudent)
	route.DELETE("/student", entity.DelStudent)

	route.GET("/place/:id", entity.GetPlace)
	route.POST("/place", entity.AddPlace)
	route.DELETE("/place", entity.DelPlace)

	route.GET("/retake/:id", entity.GetRetake)
	route.POST("/retake", entity.AddRetake)
	route.DELETE("/retake", entity.DelRetake)
}
