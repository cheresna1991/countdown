package main

import (
	"errors"
	"time"
)

var ErrNilDate = errors.New("date must not be nil")

func DaysUntilNewYear(from time.Time) (int, error) {
	if from.IsZero() {
		return 0, ErrNilDate
	}
	y, m, d := from.Date()
	loc := from.Location()
	nextYear := y
	current := time.Date(y, m, d, 0, 0, 0, 0, loc)
	newYearThisYear := time.Date(y, time.January, 1, 0, 0, 0, 0, loc)

	if !current.Before(newYearThisYear) {
		nextYear = y + 1
	}

	nextNewYear := time.Date(nextYear, time.January, 1, 0, 0, 0, 0, loc)
	days := int(nextNewYear.Sub(current).Hours() / 24)
	return days, nil
}
