package db

import (
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql" // mysql package
	"github.com/jmoiron/sqlx"
	"github.com/ruairigibney/buff/internal/data"
	"github.com/ruairigibney/buff/internal/util"
)

// DB struct holds sqlx.DB for our DB
type DB struct {
	*sqlx.DB
}

// NewDB returns the DB connection based on a DSN passed in
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func getVideoStreamsQuery() sq.SelectBuilder {
	return sq.Select(
		"id",
		"title",
		"created_time",
		"updated_time",
	).
		From("video_stream")
}

func getVideoStreamQuestionsQuery(questionID []int64) sq.SelectBuilder {
	query := sq.Select(
		"id",
		"video_stream_id",
		"text",
		"correct_answer_id",
	).
		From("video_stream_question")

	if len(questionID) > 0 {
		query = query.Where(sq.Eq{"id": questionID})
	}

	return query
}

func getVideoStreamAnswersQuery(questionID int64) sq.SelectBuilder {
	query := sq.Select(
		"id",
		"text",
	).
		From("video_stream_answer").
		Where(sq.Eq{"question_id": questionID})

	return query
}

func processPagination(baseQuery sq.SelectBuilder, pagination data.Pagination) sq.SelectBuilder {
	if pagination.Limit > 0 {
		baseQuery = baseQuery.Limit(pagination.Limit)
	}

	if pagination.Offset > 0 {
		baseQuery = baseQuery.Offset(pagination.Offset)
	}

	return baseQuery
}

// GetVideoStreams returns a list of video streams.
// It accepts a data.Pagination struct for limit and offset
func (db *DB) GetVideoStreams(pagination data.Pagination) (videoStreams []data.VideoStream, err error) {
	baseQuery := getVideoStreamsQuery()
	baseQuery = processPagination(baseQuery, pagination)

	getVideoStreamsSQL, _, err := baseQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("err getting getVideoStreamsSQL ToSQL, %v", err.Error())
	}

	err = db.Select(&videoStreams, getVideoStreamsSQL)
	if err != nil {
		return nil, fmt.Errorf("err doing db.Select for videoStream '%v', %v", getVideoStreamsSQL, err.Error())
	}

	var vsIDs []int64
	for _, videoStream := range videoStreams {
		found := util.Int64Contains(vsIDs, videoStream.ID)
		if !found {
			vsIDs = append(vsIDs, videoStream.ID)
		}
	}

	videoStreamQuestions, err := db.getVideoStreamQuestions(vsIDs, nil)
	if err != nil {
		return nil, fmt.Errorf("err doing db.getVideoStreamQuestions %v", err.Error())
	}

	for i := range videoStreamQuestions {
		for j := range videoStreams {
			if videoStreams[j].ID == videoStreamQuestions[i].VideoStreamID {
				videoStreams[j].QuestionIDs = append(videoStreams[j].QuestionIDs, videoStreamQuestions[i].ID)
			}
		}
	}

	return videoStreams, nil
}

func (db DB) getVideoStreamQuestions(videoStreamIDs []int64, questionIDs []int64) (videoStreamQuestions []data.VideoStreamQuestion, err error) {
	getVideoStreamQuestionsSQL, getVideoStreamQuestionsArgs, err := getVideoStreamQuestionsQuery(questionIDs).ToSql()
	if err != nil {
		return nil, fmt.Errorf("err getting getVideoStreamQuestionsSQL ToSQL, %v", err.Error())
	}

	err = db.Select(&videoStreamQuestions, getVideoStreamQuestionsSQL, getVideoStreamQuestionsArgs...)
	if err != nil {
		return nil, fmt.Errorf(
			"err doing db.Select for videoStreamQuestion '%v', args: '%v', %v",
			getVideoStreamQuestionsSQL,
			getVideoStreamQuestionsArgs,
			err.Error(),
		)
	}

	return videoStreamQuestions, nil
}

// GetVideoStreamQuestionAndAnswer returns a list of video stream questions and answers
// based on a questionID passed
func (db DB) GetVideoStreamQuestionAndAnswer(questionID int64) (videoStreamQuestion *data.VideoStreamQuestion, err error) {
	videoStreamQuestions, err := db.getVideoStreamQuestions(nil, []int64{questionID})
	if err != nil {
		return nil, fmt.Errorf("err in db.GetVideoStreamQuestions, %v", err.Error())
	}

	if len(videoStreamQuestions) > 0 {
		videoStreamQuestion = &videoStreamQuestions[0]
	} else {
		return nil, errors.New(data.RecordNotFoundError)
	}

	getVideoStreamAnswersSQL, getVideoStreamAnswersArgs, err := getVideoStreamAnswersQuery(videoStreamQuestion.ID).ToSql()
	if err != nil {
		return nil, fmt.Errorf("err getting getVideoStreamAnswersQuery ToSQL, %v", err.Error())
	}

	var videoStreamAnswers []data.VideoStreamAnswer
	err = db.Select(&videoStreamAnswers, getVideoStreamAnswersSQL, getVideoStreamAnswersArgs...)
	if err != nil {
		return nil, fmt.Errorf(
			"err doing db.Select for videoStreamQuestion '%v', args: '%v', %v",
			getVideoStreamAnswersSQL,
			getVideoStreamAnswersArgs,
			err.Error(),
		)
	}

	if len(videoStreamAnswers) > 0 {
		for i := range videoStreamAnswers {
			if videoStreamQuestion.CorrectAnswerID != nil &&
				*videoStreamQuestion.CorrectAnswerID == videoStreamAnswers[i].ID {
				videoStreamAnswers[i].CorrectAnswer = true
			}
		}

		videoStreamQuestion.Answers = videoStreamAnswers
	}

	return videoStreamQuestion, nil
}
