package countdown

import "time"


func DaysUntilNewYear(now time.Time) int {
	
	year := now.Year()
	newYear := time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
	if !newYear.After(now) {
		newYear = time.Date(year+1, time.January, 1, 0, 0, 0, 0, now.Location())
	}

	duration := newYear.Sub(now)
	return int(duration.Hours() / 24)
}
