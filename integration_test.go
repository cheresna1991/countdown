package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIDaysEndpoint(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	client := srv.Client()

	tests := []struct {
		name         string
		query        string
		wantStatus   int
		wantDays     int
		wantErrEmpty bool
	}{
		{
			name:         "конкретная дата — 15 июня 2024",
			query:        "?date=2024-06-15",
			wantStatus:   http.StatusOK,
			wantDays:     200,
			wantErrEmpty: true,
		},
		{
			name:         "31 декабря 2024",
			query:        "?date=2024-12-31",
			wantStatus:   http.StatusOK,
			wantDays:     1,
			wantErrEmpty: true,
		},
		{
			name:         "без параметра — сегодня",
			query:        "",
			wantStatus:   http.StatusOK,
			wantDays:     -1,
			wantErrEmpty: true,
		},
		{
			name:       "некорректный формат",
			query:      "?date=not-a-date",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.Get(srv.URL + "/days" + tt.query)
			if err != nil {
				t.Fatalf("запрос не удался: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Fatalf("статус = %d, ожидался %d", resp.StatusCode, tt.wantStatus)
			}

			var body Response
			if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
				t.Fatalf("не удалось распарсить JSON: %v", err)
			}

			if tt.wantStatus != http.StatusOK {
				if body.Err == "" {
					t.Errorf("ожидалась ошибка в теле ответа")
				}
				return
			}

			if tt.wantErrEmpty && body.Err != "" {
				t.Errorf("не ожидалась ошибка, получили: %s", body.Err)
			}

			if tt.wantDays >= 0 && body.Days != tt.wantDays {
				t.Errorf("Days = %d, ожидалось %d", body.Days, tt.wantDays)
			}

			if tt.wantDays < 0 && body.Days < 1 {
				t.Errorf("Days = %d, ожидалось >= 1", body.Days)
			}
		})
	}
}

func TestAPIRealListener(t *testing.T) {
	srv := httptest.NewServer(NewMux())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/days?date=2025-01-01")
	if err != nil {
		t.Fatalf("запрос не удался: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("статус = %d", resp.StatusCode)
	}
}
