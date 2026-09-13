package main

import (
	"testing"
	"time"
)

func TestDaysUntilNewYear(t *testing.T) {
	loc := time.UTC

	tests := []struct {
		name     string
		from     time.Time
		expected int
		wantErr  bool
	}{
		{
			name:     "1 января 2024 — до 2025",
			from:     time.Date(2024, 1, 1, 0, 0, 0, 0, loc),
			expected: 366,
		},
		{
			name:     "15 июня 2024",
			from:     time.Date(2024, 6, 15, 0, 0, 0, 0, loc),
			expected: 200,
		},
		{
			name:     "31 декабря 2024 — Новый год завтра",
			from:     time.Date(2024, 12, 31, 0, 0, 0, 0, loc),
			expected: 1,
		},
		{
			name:     "1 июля 2023",
			from:     time.Date(2023, 7, 1, 0, 0, 0, 0, loc),
			expected: 184,
		},
		{
			name:    "нулевое время — ошибка",
			from:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DaysUntilNewYear(tt.from)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ожидалась ошибка, получили nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("неожиданная ошибка: %v", err)
			}
			if got != tt.expected {
				t.Errorf("DaysUntilNewYear(%v) = %d, ожидалось %d", tt.from, got, tt.expected)
			}
		})
	}
}
