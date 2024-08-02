package domainModel

import (
	"codeGolang/global"
	"codeGolang/models"
)

type SortableScreenings []models.Screening

func (a SortableScreenings) Len() int      { return len(a) }
func (a SortableScreenings) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// SortByPrice this is Price sorting strategy
type SortByPrice struct {
	// composition of SortableScreenings
	SortableScreenings
}

func (a SortByPrice) Less(i, j int) bool {
	return a.SortableScreenings[i].Prices[global.Silver] < a.SortableScreenings[j].Prices[global.Silver]
}

// SortByStartTime this is Price sorting strategy
type SortByStartTime struct {
	SortableScreenings
}

func (a SortByStartTime) Less(i, j int) bool {
	return a.SortableScreenings[i].StartTime.UnixMilli() < a.SortableScreenings[j].StartTime.UnixMilli()
}
