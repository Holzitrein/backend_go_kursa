package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DGroup struct {
	Id int `json:"id"`
}

type Group struct {
	Name string `json:"name"`
}
type Status struct {
	Status string `json:"status"`
}

func GetGroup(c *gin.Context) {
	var response Group
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), "select nameGroup from Groups where idGroup=$1", requestBody).Scan(&response.Name)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddGroup(c *gin.Context) {
	var response Group
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO Groups (nameGroup) values ($1)", response.Name)
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

func DelGroup(c *gin.Context) {
	var response DGroup
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Groups where idGroup=$1", response.Id)
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
