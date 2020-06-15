package mock

import (
	"errors"

	"github.com/ruairigibney/buff/internal/data"
)

// MockVideoStreams contains mock data for video streams
var MockVideoStreams = []data.VideoStream{
	{ID: 1, Title: "Mock Video Stream 1", CreatedTime: 12345678, UpdatedTime: 12345678, QuestionIDs: []int64{1, 2, 3}},
	{ID: 2, Title: "Mock Video Stream 2", CreatedTime: 12345678, UpdatedTime: 12345678},
	{ID: 3, Title: "Mock Video Stream 3", CreatedTime: 12345678, UpdatedTime: 12345678},
}

// MockVideoStreamQuestion contains mock data for a video stream question
var MockVideoStreamQuestion = &data.VideoStreamQuestion{ID: 1, VideoStreamID: 1, Text: "Question 1", CorrectAnswerID: nil, Answers: MockVideoStreamAnswers}

// MockVideoStreamAnswers contians mock data for video stream answers
var MockVideoStreamAnswers = []data.VideoStreamAnswer{
	{ID: 1, Text: "Answer 1", CorrectAnswer: true},
	{ID: 2, Text: "Answer 2"},
}

// MockPagination contains mock pagination data
var MockPagination = data.Pagination{Limit: 1, Offset: 1}

// ErrMock contains a mock error
var ErrMock = errors.New("mock error")

// ErrorStruct is the mock struct for containing a mock error
type ErrorStruct struct{ ErrMock error }

// MockError contains a mock error
var MockError = ErrorStruct{}

// DB is a mock BD struct
type DB struct{}

// GetVideoStreams mocked function
func (mockDB *DB) GetVideoStreams(pagination data.Pagination) (videoStreams []data.VideoStream, err error) {
	if MockError.ErrMock != nil {
		err = MockError.ErrMock
	}
	return MockVideoStreams, err
}

// GetVideoStreamQuestions mocked function
func (mockDB *DB) GetVideoStreamQuestions(videoStreamIDs []int64, questionIDs []int64) (videoStreamQuestions []data.VideoStreamQuestion, err error) {
	return nil, nil
}

// GetVideoStreamQuestionAndAnswer mocked function
func (mockDB *DB) GetVideoStreamQuestionAndAnswer(questionID int64) (videoStreamQuestion *data.VideoStreamQuestion, err error) {
	if MockError.ErrMock != nil {
		err = MockError.ErrMock
	}
	return MockVideoStreamQuestion, err
}
