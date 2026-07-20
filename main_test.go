package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "adds two positive integers",
			query:      "param1=1&param2=2",
			wantStatus: http.StatusOK,
			wantBody:   "3",
		},
		{
			name:       "adds a negative integer",
			query:      "param1=-3&param2=1",
			wantStatus: http.StatusOK,
			wantBody:   "-2",
		},
		{
			name:       "rejects invalid param1",
			query:      "param1=invalid&param2=2",
			wantStatus: http.StatusBadRequest,
			wantBody:   "param1 must be an integer\n",
		},
		{
			name:       "rejects invalid param2",
			query:      "param1=1&param2=invalid",
			wantStatus: http.StatusBadRequest,
			wantBody:   "param2 must be an integer\n",
		},
		{
			name:       "rejects missing param1",
			query:      "param2=2",
			wantStatus: http.StatusBadRequest,
			wantBody:   "param1 must be an integer\n",
		},
		{
			name:       "rejects missing param2",
			query:      "param1=1",
			wantStatus: http.StatusBadRequest,
			wantBody:   "param2 must be an integer\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/add?"+tt.query, nil)
			recorder := httptest.NewRecorder()

			addHandler(recorder, req)

			if recorder.Code != tt.wantStatus {
				t.Errorf("status code = %d, want %d", recorder.Code, tt.wantStatus)
			}

			if recorder.Body.String() != tt.wantBody {
				t.Errorf("body = %q, want %q", recorder.Body.String(), tt.wantBody)
			}
		})
	}
}
