package tickets

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrNotFound = errors.New("no se encontró el destino")
)

type Ticket struct {
	ID           string
	Nombre       string
	Email        string
	PaisDestino  string
	HoraDelVuelo string
	Precio       float64
}

type Storage struct {
	Tickets []Ticket
}

func (s *Storage) GetTotalTickets(destination string) (int, error) {
	count := 0
	for _, ticket := range s.Tickets {
		if ticket.PaisDestino == destination {
			count++
		}
	}
	if count == 0 {
		return 0, ErrNotFound
	}
	return count, nil
}

func (s *Storage) GetCountByPeriod(timePeriod string) (int, error) {
	count := 0
	for _, ticket := range s.Tickets {
		hora, _ := strconv.Atoi(strings.Split(ticket.HoraDelVuelo, ":")[0])
		switch timePeriod {
		case "madrugada":
			if hora >= 0 && hora < 6 {
				count++
			}
		case "mañana":
			if hora >= 7 && hora < 12 {
				count++
			}
		case "tarde":
			if hora >= 13 && hora < 19 {
				count++
			}
		case "noche":
			if hora >= 20 && hora <= 23 {
				count++
			}
		}
	}
	return count, nil
}

func (s *Storage) PercentageDestination(destination string, total int) (float64, error) {
	count, err := s.GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	return (float64(count) / float64(total)) * 100, nil
}
