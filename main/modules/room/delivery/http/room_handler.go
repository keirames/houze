package http

import (
	"fmt"
	"main/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type roomHandler struct {
	roomUseCase domain.RoomUseCase
}

func NewRoomHandler(rg *gin.RouterGroup, roomUseCase domain.RoomUseCase) {
	handler := &roomHandler{
		roomUseCase,
	}

	rg.GET("/rooms", handler.GetRooms)
	rg.GET("/room/:id", handler.GetRoom)
}

func (rh *roomHandler) GetRooms(c *gin.Context) {

	rooms, err := rh.roomUseCase.GetAll()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"fail": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": rooms})
}

func (rh *roomHandler) GetRoom(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "200"})
}
