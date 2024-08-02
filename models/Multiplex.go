package models

import (
	"fmt"
)

type Multiplex struct {
	Base
	Name    string
	Screens []Screen
}

func (m *Multiplex) String() string {
	return fmt.Sprintf("Multiplex Name: %s\n \t Screens: %v\n", m.Name, m.Screens)
}
