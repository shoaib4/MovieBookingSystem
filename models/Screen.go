package models

import (
	"codeGolang/global"
	"fmt"
	"time"
)

type Screening struct {
	Movie     string
	StartTime time.Time
	EndTime   time.Time
	Prices    map[global.SeatCategory]float64
}

func (sc *Screening) String() string {
	return fmt.Sprintf("Movie: %s, StartTime: %s, EndTime: %s, Prices: %v \n", sc.Movie,
		sc.StartTime.Format(time.RFC3339), sc.EndTime.Format(time.RFC3339), sc.Prices)
}

type Screen struct {
	ID         int
	Seats      map[global.SeatCategory]int // number of seats available per category
	Screenings []Screening
}

func (s *Screen) String() string {
	return fmt.Sprintf("Screen ID: %d \nSeats: %v\nScreening: %v\n", s.ID, s.Seats, s.Screenings)
}
