package models

import (
	"codeGolang/global"
	"fmt"
	"time"
)

type Booking struct {
	Base
	Multiplex  string
	ScreenID   int
	Movie      string
	StartTime  time.Time
	Seats      map[global.SeatCategory]int
	TotalPrice float64
}

func (b *Booking) String() string {
	return fmt.Sprintf("Multiplex: %s, ScreenID: %d, Movie: %s, StartTime: %s, Seats: %v, TotalPrice: %.2f",
		b.Multiplex, b.ScreenID, b.Movie, b.StartTime.Format(time.RFC3339), b.Seats, b.TotalPrice)
}
