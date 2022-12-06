package entity

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type DRetake struct {
	Id int `json:"id"`
}

type Retake struct {
	Teacher int `json: "teacher"`
	Subject int `json: "subject"`
	TimeOn  int `json: "timeon"`
	Place   int `json: "place"`
}

type GRetake struct {
	Teacher       string `json: "teacher"`
	TeacherSecond string `json: "teachersecond"`
	Subject       string `json: "subject"`
	TimeOn        string `json: "timeon"`
	Place         string `json: "place"`
}

func GetRetake(c *gin.Context) {
	var response GRetake
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	var someTime time.Time
	err := db.QueryRow(context.Background(), `select Teachers.name, Teachers.secondname, subjects.name, timeon, places.nameplace from retakes
	LEFT JOIN Teachers ON teacher = idTeacher 
	LEFT JOIN Subjects ON subject = idsubject
	LEFT JOIN Places ON place = idplace
	where idretake=$1`, requestBody).Scan(&response.Teacher, &response.TeacherSecond, &response.Subject, &someTime, &response.Place)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	response.TimeOn = someTime.Format(time.RFC3339)
	log.Print(response.Place)
	c.JSON(200, response)
}

func AddRetake(c *gin.Context) {
	var response Retake
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	var someTime time.Time
	someTime = time.Now()
	_, err := db.Exec(context.Background(), "INSERT INTO Retakes (teacher, subject, timeon, place) values ($1, $2, $3, $4)", response.Teacher, response.Subject, someTime, response.Place)
	if err != nil {
		log.Print(err)
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	var res Status
	res.Status = "ok"
	c.JSON(200, res)
}

func DelRetake(c *gin.Context) {
	var response DRetake
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Retakes where idretake=$1", response.Id)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	var res Status
	res.Status = "ok"
	c.JSON(200, res)
}
