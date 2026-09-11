package countdown

import (
	"testing"
	"time"
)

func TestDaysUntilNewYear(t *testing.T) {
	tests := []struct {
		name string
		now  time.Time
		want int
	}{
		{
			name: "1 января — Новый год сегодня",
			now:  time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: 366, 
		},
		{
			name: "31 декабря — остался 1 день",
			now:  time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "Середина года — 1 июля 2024",
			now:  time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC),
			want: 184, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DaysUntilNewYear(tt.now)
			if got != tt.want {
				t.Errorf("DaysUntilNewYear(%v) = %d, ожидалось %d",
					tt.now.Format("2006-01-02"), got, tt.want)
			}
		})
	}
}
