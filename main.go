package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"./tickets"
)

func main() {
	storage := tickets.Storage{
		Tickets: readTicketsFile("./tickets.csv"),
	}

	var wg sync.WaitGroup

	// Requerimiento 1
	wg.Add(1)
	go func() {
		totalTickets, err := storage.GetTotalTickets("USA")
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("Total tickets to USA: %d\n", totalTickets)
		}
		wg.Done()
	}()

	// Requerimiento 2
	wg.Add(1)
	go func() {
		countMorning, err := storage.GetCountByPeriod("mañana")
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("Total tickets in the morning: %d\n", countMorning)
		}
		wg.Done()
	}()

	// Requerimiento 3
	wg.Add(1)
	go func() {
		percentage, err := storage.PercentageDestination("USA", len(storage.Tickets))
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("Percentage of tickets to USA: %.2f%%\n", percentage)
		}
		wg.Done()
	}()

	// Esperar a que todas las goroutines terminen
	wg.Wait()
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
