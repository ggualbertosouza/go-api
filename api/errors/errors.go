package errors

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError: %s (code: %s, status: %d)", e.Message, e.Code, e.StatusCode)
}

var (
	ErrInternalServer = &APIError{
		StatusCode: http.StatusInternalServerError,
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    "Internal server error",
	}

	ErrNotFound = &APIError{
		StatusCode: http.StatusNotFound,
		Code:       "NOT_FOUND",
		Message:    "The requested resource was not found",
	}

	ErrUnauthorized = &APIError{
		StatusCode: http.StatusUnauthorized,
		Code:       "UNAUTHORIZED",
		Message:    "User not authorized to access this resource",
	}

	ErrBadRequest = &APIError{
		StatusCode: http.StatusBadRequest,
		Code:       "BAD_REQUEST",
		Message:    "Invalid request",
	}
)


/*
	Funções auxiliares para mandarmos erros específicos por código
	exemplo:
		{
			StatusCode: 400,
			Code: 001, -> Código referente ao erro específico que podemos colocar nos logs a referência do código e qual o erro
			Message: "Invalid request"
		}

*/
func NewNotFoundError(message string, code string) *APIError {
	return &APIError{
		StatusCode: http.StatusNotFound,
		Code:       code,
		Message:    message,
	}
}

func NewUnauthorizedError(message string, code string) *APIError {
	return &APIError{
		StatusCode: http.StatusUnauthorized,
		Code:       code,
		Message:    message,
	}
}

func NewBadRequestError(message string, code string) *APIError {
	return &APIError{
		StatusCode: http.StatusBadRequest,
		Code:       code,
		Message:    message,
	}
}