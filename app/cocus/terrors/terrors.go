package terrors

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Err    error  `json:"-"`
	Msg    string `json:"message"`
	Status int    `json:"-"`
}

func Handler(w http.ResponseWriter, status int, err error) {

	he := HTTPError{
		Err:    err,
		Msg:    err.Error(),
		Status: status,
	}

	msg, err := json.Marshal(he)
	w.Header().Set("Content-type", "application/json")
	if status != 0 {
		w.WriteHeader(he.Status)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err == nil {
		_, _ = w.Write(msg)
	}
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}
