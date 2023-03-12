package usecase

import (
	"main/domain"
)

type roomUseCase struct {
	roomRepository domain.RoomRepository
}

func NewRoomUseCase(roomRepository domain.RoomRepository) *roomUseCase {
	return &roomUseCase{
		roomRepository,
	}
}

// TODO: return pointer rooms or rooms ? check if query return empty array or not
func (r *roomUseCase) GetAll() (*[]domain.Room, error) {
	rooms, err := r.roomRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
