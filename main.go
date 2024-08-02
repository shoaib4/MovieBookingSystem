package main

import (
	"codeGolang/domainModel"
	"codeGolang/global"
	"codeGolang/repository"
	"codeGolang/services"
	"fmt"
	"sort"
	"time"
)

func main() {
	bookkeepingService := repository.NewBookKeepingService()
	multiplexService := services.NewMultiplexService(bookkeepingService)

	// Setup data using BookKeepingServiceInterface
	bookkeepingService.AddMultiplex("Central Mall")
	bookkeepingService.AddScreen("Central Mall", 1, map[global.SeatCategory]int{
		global.Silver:   100,
		global.Gold:     50,
		global.Platinum: 25,
	})
	bookkeepingService.AddSchedule("Central Mall", 1, "Movie Title",
		time.Date(2024, time.June, 10, 14, 0, 0, 0, time.UTC),
		time.Date(2024, time.June, 10, 17, 0, 0, 0, time.UTC),
		map[global.SeatCategory]float64{
			global.Silver:   150,
			global.Gold:     200,
			global.Platinum: 250,
		})

	bookkeepingService.AddMultiplex("Inorbit Mall")
	bookkeepingService.AddScreen("Inorbit Mall", 1, map[global.SeatCategory]int{
		global.Silver:   120,
		global.Gold:     60,
		global.Platinum: 30,
	})
	bookkeepingService.AddSchedule("Inorbit Mall", 1, "Another Movie",
		time.Date(2024, time.June, 11, 10, 0, 0, 0, time.UTC),
		time.Date(2024, time.June, 11, 13, 0, 0, 0, time.UTC),
		map[global.SeatCategory]float64{
			global.Silver:   100,
			global.Gold:     180,
			global.Platinum: 220,
		})

	bookkeepingService.AddScreen("Inorbit Mall", 2, map[global.SeatCategory]int{
		global.Silver:   80,
		global.Gold:     40,
		global.Platinum: 20,
	})
	bookkeepingService.AddSchedule("Inorbit Mall", 2, "Movie Title",
		time.Date(2024, time.June, 12, 16, 0, 0, 0, time.UTC),
		time.Date(2024, time.June, 12, 19, 0, 0, 0, time.UTC),
		map[global.SeatCategory]float64{
			global.Silver:   120,
			global.Gold:     220,
			global.Platinum: 270,
		})

	// Filtering and sorting
	multiplexes := bookkeepingService.GetMultiplexes()
	fmt.Println(multiplexes)
	fmt.Println("-------")

	movieFilter := domainModel.MovieFilter{MovieTitle: "Movie Title"}
	screeningsFilteredByMovies := movieFilter.Apply(multiplexes)
	fmt.Println(screeningsFilteredByMovies)
	fmt.Println("-------")

	multiplexFilter := domainModel.MultiplexFilter{MultiplexName: "Central Mall"}
	screeningsFilteredByMultiplex := multiplexFilter.Apply(multiplexes)
	fmt.Println(screeningsFilteredByMultiplex)
	fmt.Println("-------")

	sortableScreensByPrice := domainModel.SortByPrice{SortableScreenings: screeningsFilteredByMovies}
	sort.Sort(sortableScreensByPrice)
	fmt.Println(sortableScreensByPrice)
	fmt.Println("-------")

	// Booking
	seats := map[global.SeatCategory]int{
		global.Silver: 2,
		global.Gold:   1,
	}
	booking, err := multiplexService.BookSeats("Central Mall", 1, "Movie Title",
		time.Date(2024, time.June, 10, 14, 0, 0, 0, time.UTC), seats)
	if err != nil {
		fmt.Println("Error booking seats:", err)
	} else {
		fmt.Printf("Booking successful: %+v\n", booking)
	}
}
