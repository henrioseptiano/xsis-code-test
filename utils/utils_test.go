package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"xsis-code-test/models/request"
)

func TestReadJson(t *testing.T) {
	testCases := []struct {
		name          string
		body          string
		expectedData  request.CreateMovie
		expectedError error
	}{
		{
			name: "Valid JSON",
			body: `{"title":"a","description":"a","rating":8,"image":"a.jpg"}`,
			expectedData: request.CreateMovie{
				Title:       "a",
				Description: "a",
				Rating:      8,
				Image:       "a.jpg",
			},
			expectedError: nil,
		},
		{
			name:          "Invalid JSON",
			body:          `{"name":}`,
			expectedData:  request.CreateMovie{},
			expectedError: errors.New("invalid character '}' looking for beginning of value"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "localhost:1000/Movie", bytes.NewBufferString(tc.body))
			w := httptest.NewRecorder()

			data := request.CreateMovie{}
			err := ReadJson(w, req, &data)

			if (err != nil || tc.expectedError != nil) && (err == nil || tc.expectedError == nil || err.Error() != tc.expectedError.Error()) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}

			if data != tc.expectedData {
				t.Errorf("expected data %v, got %v", tc.expectedData, data)
			}
		})
	}
}

func TestWriteJson(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
	}

	testCases := []struct {
		name           string
		status         int
		data           TestData
		headers        http.Header
		expectedBody   string
		expectedStatus int
	}{
		{
			name:           "Valid JSON",
			status:         http.StatusOK,
			data:           TestData{Name: "test"},
			headers:        http.Header{"X-Test-Header": []string{"test value"}},
			expectedBody:   `{"name":"test"}`,
			expectedStatus: http.StatusOK,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			err := WriteJson(w, tc.status, tc.data, tc.headers)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result := w.Result()
			body, _ := io.ReadAll(result.Body)

			if result.StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, result.StatusCode)
			}

			if string(body) != tc.expectedBody {
				t.Errorf("expected body %s, got %s", tc.expectedBody, string(body))
			}

			for key, value := range tc.headers {
				if !reflect.DeepEqual(result.Header[key], value) {
					t.Errorf("expected header %s to be %v, got %v", key, value, result.Header[key])
				}
			}
		})
	}
}

func TestErrorJson(t *testing.T) {
	type JSONResponse struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	testCases := []struct {
		name           string
		err            error
		status         []int
		expectedBody   string
		expectedStatus int
	}{
		{
			name:           "Error with default status",
			err:            errors.New("test error"),
			status:         nil,
			expectedBody:   `{"error":true,"message":"test error"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Error with custom status",
			err:            errors.New("test error"),
			status:         []int{http.StatusNotFound},
			expectedBody:   `{"error":true,"message":"test error"}`,
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			err := ErrorJson(w, tc.err, tc.status...)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			result := w.Result()
			body, _ := io.ReadAll(result.Body)

			if result.StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, result.StatusCode)
			}

			if string(body) != tc.expectedBody {
				t.Errorf("expected body %s, got %s", tc.expectedBody, string(body))
			}
		})
	}
}
