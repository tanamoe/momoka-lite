package apis

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

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
	e *core.ServeEvent,
	c echo.Context,
	expand models.ExpandMap,
) (T, error)

type listHandlerFunction[T comparable] func(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []T, rpage uint, rperPage int, totalItems uint, totalPages uint, err error)

type upsertHandlerFunction[T comparable] func(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	expand models.ExpandMap,
) (item T, err error)

type bulkUpsertHandlerFunction[T comparable] func(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	expand models.ExpandMap,
) (items []T, err error)

type deleteHandlerFunction func(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
) (err error)

func extractExpandMap(c echo.Context) (models.ExpandMap, error) {
	expandJson := c.QueryParam("expand")
	if expandJson == "" {
		return nil, nil
	}
	expand := make(models.ExpandMap)
	if err := json.NewDecoder(strings.NewReader(expandJson)).Decode(&expand); err != nil {
		return nil, err
	}
	return expand, nil
}

func viewRouteHandler[T comparable](
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	handler viewHandlerFunction[T],
) echo.HandlerFunc {
	return func(c echo.Context) error {
		expand, err := extractExpandMap(c)
		if err != nil {
			return handleError(app, e, c, errors.Join(err, invalidRequestError))
		}
		r, err := handler(app, e, c, expand)
		if err != nil {
			return handleError(app, e, c, err)
		}
		return c.JSON(
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
) echo.HandlerFunc {
	type listQuery struct {
		Page    uint `query:"page"`
		PerPage int  `query:"perPage"`
	}
	return func(c echo.Context) error {
		listQueryForm := &listQuery{}
		if err := c.Bind(listQueryForm); err != nil {
			return handleError(app, e, c, errors.Join(err, invalidRequestError))
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
		expand, err := extractExpandMap(c)
		if err != nil {
			return handleError(app, e, c, errors.Join(err, invalidRequestError))
		}
		items, page, perPage, totalItems, totalPages, err := handler(
			app,
			e,
			c,
			listQueryForm.Page,
			listQueryForm.PerPage,
			expand,
		)
		if err != nil {
			return handleError(app, e, c, err)
		}
		return c.JSON(
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
) echo.HandlerFunc {
	return func(c echo.Context) error {
		expand, err := extractExpandMap(c)
		if err != nil {
			return handleError(app, e, c, errors.Join(err, invalidRequestError))
		}
		r, err := handler(app, e, c, expand)
		if err != nil {
			return handleError(app, e, c, err)
		}
		return c.JSON(
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
) echo.HandlerFunc {
	return func(c echo.Context) error {
		expand, err := extractExpandMap(c)
		if err != nil {
			return handleError(app, e, c, errors.Join(err, invalidRequestError))
		}
		r, err := handler(app, e, c, expand)
		if err != nil {
			return handleError(app, e, c, err)
		}
		return c.JSON(
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

func deleteRouteHandler[T comparable](
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	handler deleteHandlerFunction,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := handler(app, e, c)
		if err != nil {
			return handleError(app, e, c, err)
		}
		return c.JSON(
			http.StatusOK,
			response{
				Success: true,
			},
		)
	}
}
