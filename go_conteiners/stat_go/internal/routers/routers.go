package routers

import (
	"retake_stat/internal/entity"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.GET("/api/group/:id", entity.GetGroup)
	route.POST("/api/group", entity.AddGroup)
	route.DELETE("/api/group", entity.DelGroup)

	route.GET("/api/subject/:id", entity.GetSubject)
	route.POST("/api/subject", entity.AddSubject)
	route.DELETE("/api/subject", entity.DelSubject)

	route.GET("/api/teacher/:id", entity.GetTeacher)
	route.POST("/api/teacher", entity.AddTeacher)
	route.DELETE("/api/teacher", entity.DelTeacher)

	route.GET("/api/student/:id", entity.GetStudent)
	route.POST("/api/student", entity.AddStudent)
	route.DELETE("/api/student", entity.DelStudent)

	route.GET("/api/place/:id", entity.GetPlace)
	route.POST("/api/place", entity.AddPlace)
	route.DELETE("/api/place", entity.DelPlace)

	route.GET("/api/retake/:id", entity.GetRetake)
	route.POST("/api/retake", entity.AddRetake)
	route.DELETE("/api/retake", entity.DelRetake)
}
