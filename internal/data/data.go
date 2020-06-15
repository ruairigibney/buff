package data

type DB struct {
	DBSource Source
}

// RecordNotFoundError const to denote records not found in DB
const RecordNotFoundError = "record not found"

// Pagination struct allows user to specify pagination limit and offset
type Pagination struct {
	Limit  uint64
	Offset uint64
}

// VideoStream holds a single video stream record
type VideoStream struct {
	ID          int64   `db:"id" json:"ID"`
	Title       string  `db:"title" json:"title"`
	CreatedTime int64   `db:"created_time" json:"createdTime"`
	UpdatedTime int64   `db:"updated_time" json:"updatedTime"`
	QuestionIDs []int64 `json:"questionIDs,omitempty"`
}

// VideoStreamQuestion holds a single video stream question record
type VideoStreamQuestion struct {
	ID              int64               `db:"id" json:"ID"`
	VideoStreamID   int64               `db:"video_stream_id" json:"VideoStreamID"`
	Text            string              `db:"text" json:"Text"`
	CorrectAnswerID *int64              `db:"correct_answer_id" json:"-"`
	Answers         []VideoStreamAnswer `json:"Answers,omitempty"`
}

// VideoStreamAnswer holds a single video stream answer record
type VideoStreamAnswer struct {
	ID            int64  `db:"id" json:"ID"`
	Text          string `db:"text" json:"Text"`
	CorrectAnswer bool   `json:"CorrectAnswer,omitempty"`
}

// Source interface
type Source interface {
	GetVideoStreams(pagination Pagination) (videoStreams []VideoStream, err error)
	GetVideoStreamQuestionAndAnswer(questionID int64) (*VideoStreamQuestion, error)
}
