package apis

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var (
	invalidRequestError = errors.New("Invalid request.")
	unauthorizedError   = errors.New("Unauthorized.")
	forbiddenError      = errors.New("Forbidden.")
	notFoundError       = errors.New("Not found.")
	unimplementedError  = errors.New("Unimplemented.")
)

func handleError(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	err error,
) error {
	if errors.Is(err, invalidRequestError) {
		return e.JSON(
			http.StatusBadRequest,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Invalid request.",
			},
		)
	}

	if errors.Is(err, unauthorizedError) {
		return e.JSON(
			http.StatusBadRequest,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Unauthorized.",
			},
		)
	}

	if errors.Is(err, forbiddenError) {
		return e.JSON(
			http.StatusBadRequest,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Forbidden.",
			},
		)
	}

	if errors.Is(err, notFoundError) {
		return e.JSON(
			http.StatusNotFound,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Not found.",
			},
		)
	}

	info, err2 := e.RequestInfo()
	if err2 != nil {
		app.Logger().Error(
			"Internal server error occur.",
			slog.Any("error", err),
		)
		return e.JSON(
			http.StatusInternalServerError,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Internal server error.",
			},
		)
	}
	app.Logger().Error(
		"Internal server error occur.",
		slog.String("route", e.Request.RequestURI),
		slog.Any("queryParams", info.Query),
		slog.Any("error", err),
	)
	return e.JSON(
		http.StatusInternalServerError,
		errorResponse{
			response: response{
				Success: false,
			},
			Message: "Internal server error.",
		},
	)
}
