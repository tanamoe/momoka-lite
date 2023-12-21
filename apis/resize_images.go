package apis

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/tools"
)

var resizableImagePathRegex = regexp.MustCompile(`^[a-zA-Z0-9-_]+\/[a-zA-Z0-9-_]+\/[a-zA-Z0-9-_]+$`)

type resizedImage struct {
	W160  string `json:"160w"`
	W320  string `json:"320w"`
	W480  string `json:"480w"`
	W640  string `json:"640w"`
	W1280 string `json:"1280w"`
	W1920 string `json:"1920w"`
}

func registerResizeImagesRoute(
	app *pocketbase.PocketBase,
	core *core.ServeEvent,
) error {
	core.Router.GET(
		"/api/resize-images",
		bulkUpsertRouteHandler(app, core, onResizeImagesRequest),
		apis.ActivityLogger(app),
	)
	return nil
}

func onResizeImagesRequest(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	expand models.ExpandMap,
) (items []resizedImage, err error) {
	images, err := extractResizeImageRequest(c)
	if err != nil {
		return nil, err
	}
	context := c.(*models.EchoContext)
	secret := context.ImagorSecret()
	for _, image := range images {
		if !isResizableImagePath(image) {
			return nil, invalidRequestError
		}
		items = append(
			items,
			resizedImage{
				W160:  tools.SignImageUrl(secret, fmt.Sprintf("120x0/filters:quality(90)/%s", image)),
				W320:  tools.SignImageUrl(secret, fmt.Sprintf("320x0/filters:quality(90)/%s", image)),
				W480:  tools.SignImageUrl(secret, fmt.Sprintf("480x0/filters:quality(90)/%s", image)),
				W640:  tools.SignImageUrl(secret, fmt.Sprintf("640x0/filters:quality(90)/%s", image)),
				W1280: tools.SignImageUrl(secret, fmt.Sprintf("1280x0/filters:quality(90)/%s", image)),
				W1920: tools.SignImageUrl(secret, fmt.Sprintf("1920x0/filters:quality(90)/%s", image)),
			},
		)
	}
	return
}

func extractResizeImageRequest(
	c echo.Context,
) ([]string, error) {
	var images []string
	rawReq := c.QueryParam("images")
	if strings.HasPrefix(rawReq, "[") {
		if err := json.Unmarshal([]byte(rawReq), &images); err != nil {
			return nil, err
		}
		return images, nil
	}
	images = strings.Split(rawReq, ",")
	for index, image := range images {
		images[index] = strings.Trim(image, " ")
	}
	return images, nil
}

func isResizableImagePath(path string) bool {
	return resizableImagePathRegex.MatchString(path)
}
