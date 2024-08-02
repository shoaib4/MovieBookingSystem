package models

import (
	"codeGolang/global"
	"fmt"
)

type Seat struct {
	Category global.SeatCategory
	Price    float64
}

func (s *Seat) String() string {
	return fmt.Sprintf("Category: %s, Price: %.2f", s.Category, s.Price)
}
