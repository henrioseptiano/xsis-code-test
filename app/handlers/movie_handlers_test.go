package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"xsis-code-test/app/usecase"
	"xsis-code-test/models/request"
	"xsis-code-test/models/response"
)

var mockAppUsecase = new(usecase.MockAppUsecase)
var appHandler = &AppHandler{AppUsecase: mockAppUsecase}

func TestCreateMovie(t *testing.T) {
	testcases := []struct {
		name           string
		expectedcode   int
		expectedresult error
		input          request.CreateMovie
	}{
		{
			name:           "valid",
			expectedcode:   http.StatusCreated,
			expectedresult: nil,
			input: request.CreateMovie{
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      4,
				Image:       "fafa.jpg",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tc.input)
			r, _ := http.NewRequest("POST", "/Movie", bytes.NewBuffer(requestBody))
			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			mockAppUsecase.Mock.On("CreateMovie", tc.input).Return(tc.expectedresult)
			appHandler.CreateMovie(w, r)

			assert.Equal(t, tc.expectedcode, w.Code)
		})
	}
}

func TestListMovie(t *testing.T) {
	testcases := []struct {
		name            string
		expectedcode    int
		expectedresult1 *[]response.ListMovie
		expectedresult2 error
		path            string
	}{
		{
			name:         "valid",
			expectedcode: http.StatusAccepted,
			expectedresult1: &[]response.ListMovie{
				{
					ID:          1,
					Title:       "Dans 1",
					Description: "Dans 1",
					Rating:      7,
					Image:       "ta.jpg",
					CreatedAt:   "2024-01-13 00:00:00",
					UpdatedAt:   "2024-01-13 00:00:00",
				},
				{
					ID:          2,
					Title:       "Dans 2",
					Description: "Dans 2",
					Rating:      7,
					Image:       "ta2.jpg",
					CreatedAt:   "2024-01-13 00:00:00",
					UpdatedAt:   "2024-01-13 00:00:00",
				},
			},
			expectedresult2: nil,
			path:            "/Movie",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", tc.path, nil)
			r.Header.Set("Content-Type", "application/json")
			mockAppUsecase.Mock.On("ListMovie").Return(tc.expectedresult1, tc.expectedresult2)
			appHandler.ListMovie(w, r)
			assert.Equal(t, tc.expectedcode, w.Code)
		})
	}
}

func TestGetMovie(t *testing.T) {
	testcases := []struct {
		name            string
		expectedcode    int
		expectedresult1 *response.GetMovie
		expectedresult2 error
		path            string
		id              string
	}{
		{
			name:         "valid",
			expectedcode: http.StatusAccepted,
			expectedresult1: &response.GetMovie{
				ID:          1,
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      7,
				Image:       "ta.jpg",
				CreatedAt:   "2024-01-13 00:00:00",
				UpdatedAt:   "2024-01-13 00:00:00",
			},
			expectedresult2: nil,
			path:            "/Movie/{id}",
			id:              "1",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", tc.path, nil)
			r.Header.Set("Content-Type", "application/json")
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.id)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
			idInt, _ := strconv.Atoi(tc.id)
			mockAppUsecase.Mock.On("GetMovie", int64(idInt)).Return(tc.expectedresult1, tc.expectedresult2)
			appHandler.GetMovie(w, r)
			assert.Equal(t, tc.expectedcode, w.Code)
		})
	}
}

func TestUpdateMovie(t *testing.T) {
	testcases := []struct {
		name           string
		expectedcode   int
		expectedresult error
		input          request.UpdateMovie
		path           string
		id             string
	}{
		{
			name:           "valid",
			expectedcode:   http.StatusOK,
			expectedresult: nil,
			input: request.UpdateMovie{
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      4,
				Image:       "fafa.jpg",
			},
			path: "/Movie/{id}",
			id:   "1",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tc.input)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PATCH", tc.path, bytes.NewBuffer(requestBody))
			r.Header.Set("Content-Type", "application/json")
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.id)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
			idInt, _ := strconv.Atoi(tc.id)
			mockAppUsecase.Mock.On("UpdateMovie", int64(idInt), tc.input).Return(tc.expectedresult)
			appHandler.UpdateMovie(w, r)
			assert.Equal(t, tc.expectedcode, w.Code)
		})
	}
}

func TestDeleteMovie(t *testing.T) {
	testcases := []struct {
		name           string
		expectedcode   int
		expectedresult error
		path           string
		id             string
	}{
		{
			name:           "valid",
			expectedcode:   http.StatusOK,
			expectedresult: nil,
			path:           "/Movie/{id}",
			id:             "1",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("DELETE", tc.path, nil)
			r.Header.Set("Content-Type", "application/json")

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.id)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
			idInt, _ := strconv.Atoi(tc.id)
			mockAppUsecase.Mock.On("DeleteMovie", int64(idInt)).Return(tc.expectedresult)
			appHandler.DeleteMovie(w, r)

			assert.Equal(t, tc.expectedcode, w.Code)
		})
	}
}
