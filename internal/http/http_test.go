package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ruairigibney/buff/internal/mock"
)

func TestVideoStreamsHandler(t *testing.T) {
	var mockDBSource mock.DB
	httpEnv := httpEnv{&mockDBSource}
	mock.MockError.ErrMock = nil

	req, err := http.NewRequest("GET", "/api/videoStreams", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpEnv.VideoStreamsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"ID":1,"title":"Mock Video Stream 1","createdTime":12345678,"updatedTime":12345678,"questionIDs":[1,2,3]},{"ID":2,"title":"Mock Video Stream 2","createdTime":12345678,"updatedTime":12345678},{"ID":3,"title":"Mock Video Stream 3","createdTime":12345678,"updatedTime":12345678}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestVideoStreamsHandlerWithPagination(t *testing.T) {
	var mockDBSource mock.DB
	httpEnv := httpEnv{&mockDBSource}
	mock.MockError.ErrMock = nil

	req, err := http.NewRequest("GET", "/api/videoStreams?limit=10&offset=0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpEnv.VideoStreamsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"ID":1,"title":"Mock Video Stream 1","createdTime":12345678,"updatedTime":12345678,"questionIDs":[1,2,3]},{"ID":2,"title":"Mock Video Stream 2","createdTime":12345678,"updatedTime":12345678},{"ID":3,"title":"Mock Video Stream 3","createdTime":12345678,"updatedTime":12345678}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestVideoStreamsHandlerWithErrorFromDB(t *testing.T) {
	var mockDBSource mock.DB
	httpEnv := httpEnv{&mockDBSource}
	mock.MockError.ErrMock = mock.ErrMock

	req, err := http.NewRequest("GET", "/api/videoStreams?limit=2&offset=0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpEnv.VideoStreamsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `err while getting video streams`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestVideoStreamQuestionHandler(t *testing.T) {
	var mockDBSource mock.DB
	httpEnv := httpEnv{&mockDBSource}
	mock.MockError.ErrMock = nil

	req, err := http.NewRequest("GET", "/api/videoStreams/questions/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpEnv.VideoStreamQuestionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"ID":1,"VideoStreamID":1,"Text":"Question 1","Answers":[{"ID":1,"Text":"Answer 1","CorrectAnswer":true},{"ID":2,"Text":"Answer 2"}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestVideoStreamQuestionHandlerWithDBError(t *testing.T) {
	var mockDBSource mock.DB
	httpEnv := httpEnv{&mockDBSource}
	mock.MockError.ErrMock = mock.ErrMock

	req, err := http.NewRequest("GET", "/api/videoStreams/questions/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpEnv.VideoStreamQuestionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestVideoStreamQuestionHandlerWithURIPathError(t *testing.T) {
	var mockDBSource mock.DB
	httpEnv := httpEnv{&mockDBSource}
	mock.MockError.ErrMock = nil

	req, err := http.NewRequest("GET", "/api/videoStreams/questions/", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id": "NaN",
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpEnv.VideoStreamQuestionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
