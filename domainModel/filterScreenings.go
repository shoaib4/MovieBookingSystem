package domainModel

import "codeGolang/models"

type Filter interface {
	Apply([]*models.Multiplex) []models.Screening
}

type MovieFilter struct {
	MovieTitle string
}

func (f MovieFilter) Apply(multiplexes []*models.Multiplex) []models.Screening {
	var result []models.Screening
	// Repo Get methords could be used
	for _, multiplex := range multiplexes {
		for _, screen := range multiplex.Screens {
			for _, screening := range screen.Screenings {
				if screening.Movie == f.MovieTitle {
					result = append(result, screening)
				}
			}
		}
	}
	return result
}

type MultiplexFilter struct {
	MultiplexName string
}

func (f MultiplexFilter) Apply(multiplexes []*models.Multiplex) []models.Screening {
	var result []models.Screening
	for _, multiplex := range multiplexes {
		if multiplex.Name == f.MultiplexName {
			for _, screen := range multiplex.Screens {
				for _, screening := range screen.Screenings {
					result = append(result, screening)
				}
			}
		}
	}
	return result
}
