package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DSubject struct {
	Id int `json:"id"`
}

type Subject struct {
	Name  string `json:"name"`
	TypeS int    `json:"type"`
}

func GetSubject(c *gin.Context) {
	var response Subject
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), "select name, type from Subjects where idSubject=$1", requestBody).Scan(&response.Name, &response.TypeS)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddSubject(c *gin.Context) {
	var response Subject
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO Subjects (name, type) values ($1, $2)", response.Name, response.TypeS)
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

func DelSubject(c *gin.Context) {
	var response DSubject
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Subjects where idSubject=$1", response.Id)
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
