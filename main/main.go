package main

import (
	"fmt"

	"main/database"
	"main/modules/room/delivery/http"
	"main/modules/room/repository"
	useCase "main/modules/room/use_case"
	"main/tools"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	err := tools.LoadConfig(".")
	if err != nil {
		fmt.Println(err)
		panic("Fail to load config")
	}

	err = database.Connect()
	if err != nil {
		fmt.Println(err)
		panic("Fail to load database")
	}

	router := gin.Default()

	apiV1 := router.Group("/v1")

	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"1": "!"})
	})

	roomRepository := repository.NewRoomRepository(database.Conn)
	roomUseCase := useCase.NewRoomUseCase(roomRepository)
	http.NewRoomHandler(apiV1, roomUseCase)

	err = router.Run(tools.Config.ServerAddress)
	if err != nil {
		panic("Fail to start server")
	}
}
