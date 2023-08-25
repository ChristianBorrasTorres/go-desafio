package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"./tickets"
)

func main() {
	storage := tickets.Storage{
		Tickets: readTicketsFile("./tickets.csv"),
	}

	totalTickets, _ := storage.GetTotalTickets("USA")
	fmt.Printf("Total tickets to USA: %d\n", totalTickets)

	countMorning, _ := storage.GetCountByPeriod("mañana")
	fmt.Printf("Total tickets in the morning: %d\n", countMorning)

	percentage, _ := storage.PercentageDestination("USA", len(storage.Tickets))
	fmt.Printf("Percentage of tickets to USA: %.2f%%\n", percentage)
}

func readTicketsFile(filename string) []tickets.Ticket {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var ticketsList []tickets.Ticket
	for _, each := range rawCSVdata {
		price, _ := strconv.ParseFloat(each[5], 64)
		ticket := tickets.Ticket{
			ID:           each[0],
			Nombre:       each[1],
			Email:        each[2],
			PaisDestino:  each[3],
			HoraDelVuelo: each[4],
			Precio:       price,
		}
		ticketsList = append(ticketsList, ticket)
	}

	return ticketsList
}
