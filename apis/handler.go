package apis

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/tools"
)

type handlerFunc func(e *core.RequestEvent) error

type response struct {
	Success bool `json:"success"`
}

type viewResponse[T comparable] struct {
	response

	Item T `json:"item"`
}

type listResponse[T comparable] struct {
	response

	Page       uint `json:"page"`
	PerPage    int  `json:"perPage"`
	TotalItems uint `json:"totalItems"`
	TotalPages uint `json:"totalPages"`
	Items      []T  `json:"items"`
}

type upsertResponse[T comparable] struct {
	response

	Item T `json:"item"`
}

type bulkUpsertResponse[T comparable] struct {
	response

	Items []T `json:"items"`
}

type errorResponse struct {
	response
	Message string `json:"message"`
}

type viewHandlerFunction[T comparable] func(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	expand models.ExpandMap,
) (T, error)

type listHandlerFunction[T comparable] func(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []T, rpage uint, rperPage int, totalItems uint, totalPages uint, err error)

type upsertHandlerFunction[T comparable] func(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	expand models.ExpandMap,
) (item T, err error)

type bulkUpsertHandlerFunction[T comparable] func(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	expand models.ExpandMap,
) (items []T, err error)

type deleteHandlerFunction func(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
) (err error)

func viewRouteHandler[T comparable](
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	handler viewHandlerFunction[T],
) handlerFunc {
	return func(e *core.RequestEvent) error {
		info, err := e.RequestInfo()
		if err != nil {
			return err
		}
		expand, err := tools.ExtractExpandMap(info)
		if err != nil {
			return handleError(app, e, errors.Join(err, invalidRequestError))
		}
		r, err := handler(app, e, expand)
		if err != nil {
			return handleError(app, e, err)
		}
		return e.JSON(
			http.StatusOK,
			viewResponse[T]{
				response: response{
					Success: true,
				},
				Item: r,
			},
		)
	}
}

func listRouteHandler[T comparable](
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	handler listHandlerFunction[T],
) handlerFunc {
	type listQuery struct {
		Page    uint `query:"page"`
		PerPage int  `query:"perPage"`
	}
	return func(e *core.RequestEvent) error {
		info, err := e.RequestInfo()
		if err != nil {
			return err
		}
		rPage := info.Query["page"]
		qPage, err := strconv.Atoi(rPage)
		if err != nil {
			if rPage != "" {
				return handleError(app, e, errors.Join(err, invalidRequestError))
			} else {
				qPage = 0
			}
		}
		rPerPage := info.Query["perPage"]
		qPerPage, err := strconv.Atoi(rPerPage)
		if err != nil {
			if rPerPage != "" {
				return handleError(app, e, errors.Join(err, invalidRequestError))
			} else {
				qPerPage = 0
			}
		}
		listQueryForm := &listQuery{
			Page:    uint(qPage),
			PerPage: qPerPage,
		}
		if listQueryForm.Page <= 0 {
			listQueryForm.Page = 1
		}
		if listQueryForm.PerPage <= 0 {
			listQueryForm.PerPage = -1
		}
		if listQueryForm.PerPage > 150 {
			listQueryForm.PerPage = 150
		}
		expand, err := tools.ExtractExpandMap(info)
		if err != nil {
			return handleError(app, e, errors.Join(err, invalidRequestError))
		}
		items, page, perPage, totalItems, totalPages, err := handler(
			app,
			e,
			listQueryForm.Page,
			listQueryForm.PerPage,
			expand,
		)
		if err != nil {
			return handleError(app, e, err)
		}
		return e.JSON(
			http.StatusOK,
			listResponse[T]{
				response: response{
					Success: true,
				},
				Page:       page,
				PerPage:    perPage,
				TotalItems: totalItems,
				TotalPages: totalPages,
				Items:      items,
			},
		)
	}
}

func upsertRouteHandler[T comparable](
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	handler upsertHandlerFunction[T],
) handlerFunc {
	return func(e *core.RequestEvent) error {
		info, err := e.RequestInfo()
		if err != nil {
			return err
		}
		expand, err := tools.ExtractExpandMap(info)
		if err != nil {
			return handleError(app, e, errors.Join(err, invalidRequestError))
		}
		r, err := handler(app, e, expand)
		if err != nil {
			return handleError(app, e, err)
		}
		return e.JSON(
			http.StatusOK,
			upsertResponse[T]{
				response: response{
					Success: true,
				},
				Item: r,
			},
		)
	}
}

func bulkUpsertRouteHandler[T comparable](
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	handler bulkUpsertHandlerFunction[T],
) handlerFunc {
	return func(e *core.RequestEvent) error {
		info, err := e.RequestInfo()
		if err != nil {
			return err
		}
		expand, err := tools.ExtractExpandMap(info)
		if err != nil {
			return handleError(app, e, errors.Join(err, invalidRequestError))
		}
		r, err := handler(app, e, expand)
		if err != nil {
			return handleError(app, e, err)
		}
		return e.JSON(
			http.StatusOK,
			bulkUpsertResponse[T]{
				response: response{
					Success: true,
				},
				Items: r,
			},
		)
	}
}

func deleteRouteHandler(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	handler deleteHandlerFunction,
) handlerFunc {
	return func(e *core.RequestEvent) error {
		err := handler(app, e)
		if err != nil {
			return handleError(app, e, err)
		}
		return e.JSON(
			http.StatusOK,
			response{
				Success: true,
			},
		)
	}
}
