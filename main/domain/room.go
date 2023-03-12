package domain

type Room struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type RoomRepository interface {
	FindByID(id int64) (*Room, error)
	FindAll() (*[]Room, error)
}

type RoomUseCase interface {
	GetAll() (*[]Room, error)
}
