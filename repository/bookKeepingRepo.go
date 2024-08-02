package repository

import (
	"codeGolang/global"
	"codeGolang/models"
	"errors"
	"sync"
	"time"
)

// todo lock when needed
// screens
// 100  -- 110 -- 120
//	trancction
//	screend for multiplex == 1 updteing seat 10
//	commits x
//	commit OK
//
//	trancction
//	screend for multiplex == 1 updteing seat 10
//	commits

// BookKeepingService manages multiplex data.
type BookKeepingService struct {
	multiplexes map[string]*models.Multiplex
	mu          sync.Mutex
}

// NewBookKeepingService creates a new instance of BookKeepingService.
func NewBookKeepingService() IRepository {
	return &BookKeepingService{
		multiplexes: make(map[string]*models.Multiplex),
	}
}

// AddMultiplex adds a new multiplex.
func (s *BookKeepingService) AddMultiplex(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.multiplexes[name] = &models.Multiplex{Name: name, Screens: []models.Screen{}}
}

// AddScreen adds a new screen to a multiplex.
func (s *BookKeepingService) AddScreen(multiplexName string, screenID int, seats map[global.SeatCategory]int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	multiplex, ok := s.multiplexes[multiplexName]
	if !ok {
		return errors.New("multiplex not found")
	}
	screen := models.Screen{ID: screenID, Seats: seats, Screenings: []models.Screening{}}
	multiplex.Screens = append(multiplex.Screens, screen)
	return nil
}

// AddSchedule adds a new schedule to a screen.
func (s *BookKeepingService) AddSchedule(multiplexName string, screenID int, movie string, startTime, endTime time.Time, prices map[global.SeatCategory]float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	multiplex, ok := s.multiplexes[multiplexName]
	if !ok {
		return errors.New("multiplex not found")
	}
	for i := range multiplex.Screens {
		if multiplex.Screens[i].ID == screenID {
			schedule := models.Screening{
				Movie:     movie,
				StartTime: startTime,
				EndTime:   endTime,
				Prices:    prices,
			}
			multiplex.Screens[i].Screenings = append(multiplex.Screens[i].Screenings, schedule)
			return nil
		}
	}
	return errors.New("screen not found")
}

// GetMultiplexes retrieves all multiplexes.
func (s *BookKeepingService) GetMultiplexes() []*models.Multiplex {
	s.mu.Lock()
	defer s.mu.Unlock()
	var multiplexes []*models.Multiplex
	for _, m := range s.multiplexes {
		multiplexes = append(multiplexes, m)
	}
	return multiplexes
}
