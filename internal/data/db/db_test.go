package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	mockData "github.com/ruairigibney/buff/internal/mock"
)

// a successful case
func TestGetVideoStreams(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%v' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mockDBSource := DB{sqlxDB}
	mockData.MockError.ErrMock = nil

	rows := []string{"id", "title", "created_time", "updated_time"}
	mock.ExpectQuery(`SELECT id, title, created_time, updated_time FROM video_stream LIMIT 1 OFFSET 1`).
		WillReturnRows(sqlmock.NewRows(rows).AddRow(
			mockData.MockVideoStreams[0].ID,
			mockData.MockVideoStreams[0].Title,
			mockData.MockVideoStreams[0].CreatedTime,
			mockData.MockVideoStreams[0].UpdatedTime,
		))

	rows = []string{"id", "video_stream_id", "text", "correct_answer_id"}
	mock.ExpectQuery(`SELECT id, video_stream_id, text, correct_answer_id FROM video_stream_question`).
		WillReturnRows(sqlmock.NewRows(rows).AddRow(
			mockData.MockVideoStreamQuestion.ID,
			mockData.MockVideoStreamQuestion.VideoStreamID,
			mockData.MockVideoStreamQuestion.Text,
			mockData.MockVideoStreamQuestion.CorrectAnswerID,
		))

	if _, err := mockDBSource.GetVideoStreams(mockData.MockPagination); err != nil {
		t.Errorf("error while getting video streams: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//GetVideoStreamQuestionAndAnswer
func TestGetVideoStreamQuestionAndAnswer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%v' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mockDBSource := DB{sqlxDB}
	mockData.MockError.ErrMock = nil

	rows := []string{"id", "video_stream_id", "text", "correct_answer_id"}
	mock.ExpectQuery(`SELECT id, video_stream_id, text, correct_answer_id FROM video_stream_question WHERE id IN (?)`).
		WithArgs(mockData.MockVideoStreamQuestion.ID).
		WillReturnRows(sqlmock.NewRows(rows).AddRow(
			mockData.MockVideoStreamQuestion.ID,
			mockData.MockVideoStreamQuestion.VideoStreamID,
			mockData.MockVideoStreamQuestion.Text,
			mockData.MockVideoStreamQuestion.CorrectAnswerID,
		))

	rows = []string{"id", "text"}
	mock.ExpectQuery(`SELECT id, text FROM video_stream_answer WHERE question_id = ?`).
		WithArgs(mockData.MockVideoStreamQuestion.ID).
		WillReturnRows(sqlmock.NewRows(rows).AddRow(
			mockData.MockVideoStreamAnswers[0].ID,
			mockData.MockVideoStreamAnswers[0].Text,
		))

	if _, err := mockDBSource.GetVideoStreamQuestionAndAnswer(mockData.MockVideoStreamQuestion.ID); err != nil {
		t.Errorf("error while getting video streams: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
