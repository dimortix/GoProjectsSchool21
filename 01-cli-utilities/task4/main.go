package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Visit struct {
	Specialization string
	Date           time.Time
}

type PatientNotFoundError struct {
	Name string
}

func (e *PatientNotFoundError) Error() string {
	return "patient not found"
}

type VisitLog struct {
	entries map[string][]Visit
}

func NewVisitLog() *VisitLog {
	return &VisitLog{
		entries: make(map[string][]Visit),
	}
}

func (v *VisitLog) Save(name, specialization, dateStr string) error {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}
	visit := Visit{
		Specialization: specialization,
		Date:           date,
	}
	v.entries[name] = append(v.entries[name], visit)
	return nil
}

func (v *VisitLog) GetHistory(name string) ([]Visit, error) {
	visits, ok := v.entries[name]
	if !ok || len(visits) == 0 {
		return nil, &PatientNotFoundError{Name: name}
	}
	return visits, nil
}

func (v *VisitLog) GetLastVisit(name, specialization string) (time.Time, error) {
	visits, ok := v.entries[name]
	if !ok || len(visits) == 0 {
		return time.Time{}, &PatientNotFoundError{Name: name}
	}

	var (
		found bool
		last  time.Time
	)

	for _, visit := range visits {
		if visit.Specialization != specialization {
			continue
		}
		if !found || visit.Date.After(last) {
			found = true
			last = visit.Date
		}
	}

	if !found {
		return time.Time{}, &PatientNotFoundError{Name: name}
	}

	return last, nil
}

func main() {
	log := NewVisitLog()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			return
		}
		command := strings.TrimSpace(scanner.Text())
		if command == "" {
			continue
		}

		switch command {
		case "Save":
			if !scanner.Scan() {
				return
			}
			name := strings.TrimSpace(scanner.Text())

			if !scanner.Scan() {
				return
			}
			spec := strings.TrimSpace(scanner.Text())

			if !scanner.Scan() {
				return
			}
			dateStr := strings.TrimSpace(scanner.Text())

			if err := log.Save(name, spec, dateStr); err != nil {
				fmt.Println("Invalid input")
			}

		case "GetHistory":
			if !scanner.Scan() {
				return
			}
			name := strings.TrimSpace(scanner.Text())

			visits, err := log.GetHistory(name)
			if err != nil {
				var notFound *PatientNotFoundError
				if errors.As(err, &notFound) {
					fmt.Println(notFound.Error())
				} else {
					fmt.Println("Invalid input")
				}
				continue
			}

			for _, visit := range visits {
				fmt.Printf("%s %s\n", visit.Specialization, visit.Date.Format("2006-01-02"))
			}

		case "GetLastVisit":
			if !scanner.Scan() {
				return
			}
			name := strings.TrimSpace(scanner.Text())

			if !scanner.Scan() {
				return
			}
			spec := strings.TrimSpace(scanner.Text())

			date, err := log.GetLastVisit(name, spec)
			if err != nil {
				var notFound *PatientNotFoundError
				if errors.As(err, &notFound) {
					fmt.Println(notFound.Error())
				} else {
					fmt.Println("Invalid input")
				}
				continue
			}

			fmt.Println(date.Format("2006-01-02"))

		default:
			fmt.Println("Invalid input")
		}
	}
}

