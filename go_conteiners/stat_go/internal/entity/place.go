package entity

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type DPlace struct {
	Id int `json:"id"`
}

type Place struct {
	NamePlace    string `json:"name"`
	Proector     bool   `json: "proector"`
	NumSeats     int    `json: "numseats"`
	NumComputers int    `json: "numcomputers"`
}

func GetPlace(c *gin.Context) {
	var response Place
	var err_r error_res
	requestBody := c.Param("id")
	log.Print(requestBody)
	err := db.QueryRow(context.Background(), "select nameplace, proector, numseats, numcomputers from places where idplace=$1", requestBody).Scan(&response.NamePlace, &response.Proector, &response.NumSeats, &response.NumComputers)
	if err != nil {
		log.Print("db error")
		err_r.Err_s = "failed data from db"
		c.JSON(400, err_r)
		return
	}
	c.JSON(200, response)
}

func AddPlace(c *gin.Context) {
	var response Place
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "INSERT INTO Places (nameplace, proector, numseats, numcomputers) values ($1, $2, $3, $4)", response.NamePlace, response.Proector, response.NumSeats, response.NumComputers)
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

func DelPlace(c *gin.Context) {
	var response DSubject
	var err_r error_res
	if err := c.BindJSON(&response); err != nil {
		err_r.Err_s = "failed request"
		c.JSON(400, err_r)
		return
	}
	_, err := db.Exec(context.Background(), "delete from Places where idplace=$1", response.Id)
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
