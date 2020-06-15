package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ruairigibney/buff/internal/data"
	"github.com/ruairigibney/buff/internal/data/db"
)

type httpEnv struct {
	DBSource data.Source
}

// StartAndServe will start the server
func StartAndServe(db *db.DB) {
	env := &httpEnv{db}

	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/videoStreams", env.VideoStreamsHandler).Methods("GET")
	rtr.HandleFunc("/api/videoStreams/questions/{id:[0-9]+}", env.VideoStreamQuestionHandler).Methods("GET")
	http.Handle("/", rtr)
	log.Println("Listening...")
	panic(http.ListenAndServe(":8080", nil))
}

func getPagination(request *http.Request) (limit uint64, offset uint64, err error) {
	reqQuery := request.URL.Query()
	l, o := reqQuery.Get("limit"), reqQuery.Get("offset")
	if l != "" {
		limit, err = strconv.ParseUint(l, 10, 64)
	}
	if o != "" {
		offset, err = strconv.ParseUint(o, 10, 64)
	}

	return limit, offset, nil
}

// VideoStreamsHandler is the handler for /api/videoStreams.
// It can accept pagination query params limit & offset.
func (env *httpEnv) VideoStreamsHandler(writer http.ResponseWriter, request *http.Request) {
	log.Print("processing videoStreamsHandler")

	limit, offset, err := getPagination(request)
	if err != nil {
		errorString := "err while parsing pagination query"
		returnError(writer, http.StatusInternalServerError, errorString)
		log.Printf("%v\n%v", errorString, err)
	}
	videoStreams, err := env.DBSource.GetVideoStreams(data.Pagination{Limit: limit, Offset: offset})
	if err != nil {
		errorString := "err while getting video streams"
		returnError(writer, http.StatusInternalServerError, errorString)
		log.Printf("%v\n%v", errorString, err)
		return
	}

	body, err := json.Marshal(videoStreams)
	if err != nil {
		errorString := "err while marshalling json"
		returnError(writer, http.StatusInternalServerError, errorString)
		log.Printf("%v\n%v", errorString, err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(body))
}

// VideoStreamQuestionHandler is the handler for /api/videoStreams/questions/{id:[0-9]+}
// It requires the question ID to be specified in the URI
func (env *httpEnv) VideoStreamQuestionHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	questionID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		errorString := "err while parsing video stream question ID"
		returnError(writer, http.StatusInternalServerError, errorString)
		log.Printf("%v\n%v", errorString, err)
	}
	log.Printf("processing videoStreamQuestionHandler for %v", questionID)
	videoStreamQuestion, err := env.DBSource.GetVideoStreamQuestionAndAnswer(questionID)
	if err != nil {
		if err.Error() == data.RecordNotFoundError {
			returnError(writer, http.StatusNotFound, data.RecordNotFoundError)
			log.Printf("video stream question %v not found", questionID)
			return
		}

		errorString := "err while getting video stream question"
		returnError(writer, http.StatusInternalServerError, errorString)
		log.Printf("%v\n%v", errorString, err)
		return
	}

	body, err := json.Marshal(videoStreamQuestion)
	if err != nil {
		errorString := "err while marshalling json"
		returnError(writer, http.StatusInternalServerError, errorString)
		log.Printf("%v\n%v", errorString, err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(body))
}

func returnError(writer http.ResponseWriter, code int, errorString string) {
	writer.WriteHeader(code)
	writer.Write([]byte(errorString))

	return
}
