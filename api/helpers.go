package api

import (
	"github.com/labstack/echo/v4"
)

type Response[T any, E any] struct {
	Message string `json:"message"`
	Code    int    `json:"status_code"`
	Details T      `json:"details"`
	Error   E      `json:"errors"`
	Ok      bool   `json:"ok"`
}

// NewResponse creates a new Response object.
func NewResponse[T any, E any](ok bool, message string, code int, details T, err E) Response[T, E] {
	return Response[T, E]{
		Message: message,
		Code:    code,
		Details: details,
		Error:   err,
		Ok:      ok,
	}
}

// ResponseProvider sends a JSON response using the echo context.
func ResponseProvider[T any, E any](response Response[T, E], c echo.Context) error {
	return c.JSON(response.Code, map[string]interface{}{
		"message":     response.Message,
		"errors":      response.Error,
		"ok":          response.Ok,
		"status_code": response.Code,
		"details":     response.Details,
	})
}
