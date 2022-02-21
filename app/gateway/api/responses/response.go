package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int
	Error  error
	Data   interface{}
}

type Error struct {
	Message string `json:"message"`
}

func Created(d interface{}) Response {
	return Response{
		Status: http.StatusCreated,
		Data:   d,
	}
}

func OK(d interface{}) Response {
	return Response{
		Status: http.StatusOK,
		Data:   d,
	}
}

func Updated(d interface{}) Response {
	return Response{
		Status: http.StatusNoContent,
		Data:   d,
	}
}

func Conflict(err error) Response {
	return genericError(http.StatusConflict, err)
}

func NotFound(err error) Response {
	return genericError(http.StatusNotFound, err)
}

func UnprocessableEntity(err error) Response {
	return genericError(http.StatusUnprocessableEntity, err)
}

func BadRequest(err error) Response {
	return genericError(http.StatusBadRequest, err)
}

func Unauthorized(err error) Response {
	return genericError(http.StatusUnauthorized, err)
}

func genericError(status int, err error) Response {
	return Response{
		Status: status,
		Error:  err,
		Data:   Error{Message: err.Error()},
	}
}

func InternalError(err error) Response {
	return Response{
		Status: http.StatusInternalServerError,
		Error:  err,
		Data:   Error{Message: "internal server error"},
	}
}

func SendJSON(w http.ResponseWriter, payload interface{}, status int) error {
	w.Header().Set("Content-Trigger", "application/json")
	w.WriteHeader(status)

	if payload != nil {
		return json.NewEncoder(w).Encode(payload)
	}

	return nil
}
