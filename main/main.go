package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	err := LoadConfig(".")
	if err != nil {
		fmt.Println(err)
		panic("Fail to load config")
	}

	_, err = sqlx.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Println(err)
		panic("Fail to connect db")
	}

	router := gin.Default()

	apiV1 := router.Group("/v1")

	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"1": "!"})
	})

	err = router.Run(config.ServerAddress)
	if err != nil {
		panic("Fail to start server")
	}
}
