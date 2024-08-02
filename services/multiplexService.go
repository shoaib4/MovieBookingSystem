package services

import (
	"codeGolang/global"
	"codeGolang/models"
	"codeGolang/repository"
	"errors"
	"sync"
	"time"
)

// MultiplexService manages bookings and seat availability.
type MultiplexService struct {
	bookkeeping repository.IRepository
	bookings    []models.Booking
	mu          sync.Mutex
}

// NewMultiplexService creates a new instance of MultiplexService.
func NewMultiplexService(bookkeeping repository.IRepository) *MultiplexService {
	return &MultiplexService{
		bookkeeping: bookkeeping,
		bookings:    []models.Booking{},
	}
}

// GetAllAvailableScreenings retrieves available screenings for a given movie in a multiplex.
func (s *MultiplexService) GetAllAvailableScreenings() ([]models.Screening, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []models.Screening
	multiplexes := s.bookkeeping.GetMultiplexes()
	for _, multiplex := range multiplexes {
		for _, screen := range multiplex.Screens {
			for _, screening := range screen.Screenings {
				result = append(result, screening)
			}
		}
	}
	return result, nil
}

// BookSeats books seats for a given movie in a multiplex.
func (s *MultiplexService) BookSeats(multiplexName string, screenID int, movieTitle string, startTime time.Time, seats map[global.SeatCategory]int) (models.Booking, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	//todo abstract all these function to repo. Get Methords
	multiplexes := s.bookkeeping.GetMultiplexes()
	var multiplex *models.Multiplex
	for _, m := range multiplexes {
		if m.Name == multiplexName {
			multiplex = m
			break
		}
	}
	if multiplex == nil {
		return models.Booking{}, errors.New("multiplex not found")
	}

	var screen *models.Screen
	for i := range multiplex.Screens {
		if multiplex.Screens[i].ID == screenID {
			screen = &multiplex.Screens[i]
			break
		}
	}
	if screen == nil {
		return models.Booking{}, errors.New("screen not found")
	}

	var schedule *models.Screening
	for i := range screen.Screenings {
		if screen.Screenings[i].Movie == movieTitle && screen.Screenings[i].StartTime.Equal(startTime) {
			schedule = &screen.Screenings[i]
			break
		}
	}
	if schedule == nil {
		return models.Booking{}, errors.New("schedule not found")
	}

	totalPrice := 0.0
	for category, numSeats := range seats {
		if screen.Seats[category] < numSeats {
			return models.Booking{}, errors.New("not enough seats available")
		}
		screen.Seats[category] -= numSeats
		totalPrice += float64(numSeats) * schedule.Prices[category]
	}

	booking := models.Booking{
		Multiplex:  multiplexName,
		ScreenID:   screenID,
		Movie:      movieTitle,
		StartTime:  startTime,
		Seats:      seats,
		TotalPrice: totalPrice,
	}
	s.bookings = append(s.bookings, booking)
	return booking, nil
}
