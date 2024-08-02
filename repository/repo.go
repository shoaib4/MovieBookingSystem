package repository

import (
	"codeGolang/global"
	"codeGolang/models"
	"time"
)

// todo make fuctions with entitys
// todo add get
// IRepository defines the methods that our repository service must implement.
type IRepository interface {
	AddMultiplex(name string)
	AddScreen(multiplexName string, screenID int, seats map[global.SeatCategory]int) error
	AddSchedule(multiplexName string, screenID int, movie string, startTime, endTime time.Time, prices map[global.SeatCategory]float64) error
	GetMultiplexes() []*models.Multiplex
	//todo
	//AddBooking(book models.Booking) error
	//GetMulitplex() []*models.Multiplex
	////all getmethords
	//update
}
