package repository

import (
	"main/domain"

	"github.com/jmoiron/sqlx"
)

type roomRepository struct {
	Conn *sqlx.DB
}

func NewRoomRepository(conn *sqlx.DB) *roomRepository {
	return &roomRepository{
		Conn: conn,
	}
}

func (r *roomRepository) FindByID(id int64) (*domain.Room, error) {
	var room domain.Room

	err := r.Conn.Get(&room, "SELECT * FROM rooms WHERE rooms.id=?", id)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *roomRepository) FindAll() (*[]domain.Room, error) {
	var rooms []domain.Room

	err := r.Conn.Select(&rooms, "SELECT * FROM rooms")
	if err != nil {
		return nil, err
	}

	return &rooms, nil
}
