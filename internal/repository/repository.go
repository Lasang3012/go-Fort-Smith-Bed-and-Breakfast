package repository

import "github.com/Lasang3012/go-Fort-Smith-Bed-and-Breakfast/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
}
