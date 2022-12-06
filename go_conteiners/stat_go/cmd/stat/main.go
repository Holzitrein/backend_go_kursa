package main

import (
	"context"
	"log"
	"retake_stat/internal/entity"
	"retake_stat/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	router := gin.Default()
	routers.Routes(router)
	info_db := "postgres://mirea:mirea@database:5432/mirea"
	conn, err := pgx.Connect(context.Background(), info_db)
	if err != nil {
		log.Print("lol")
		//os.Exit(1)
	}
	entity.SetDb(conn)

	defer conn.Close(context.Background())
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
