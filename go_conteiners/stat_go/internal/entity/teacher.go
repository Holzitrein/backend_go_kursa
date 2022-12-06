package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DTeacher struct {
	Id int `json:"id"`
}

type Teacher struct {
	Name        string `json:"name"`
	SecondName  string `json:"secondname"`
	Departament string `json:"department"`
}

func GetTeacher(c *gin.Context) {
	var response Teacher
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), "select name, secondName, department from Teachers where idTeacher=$1", requestBody).Scan(&response.Name, &response.SecondName, &response.Departament)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddTeacher(c *gin.Context) {
	var response Teacher
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO Teachers (name, secondName, department) values ($1, $2, $3)", response.Name, response.SecondName, response.Departament)
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

func DelTeacher(c *gin.Context) {
	var response DTeacher
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Teachers where idTeacher=$1", response.Id)
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
