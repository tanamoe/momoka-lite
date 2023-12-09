package tools

import (
	"encoding/json"
	"strings"

	"github.com/labstack/echo/v5"
	"tana.moe/momoka-lite/models"
)

func flatToExpandMap(mapping models.ExpandMap, field string) models.ExpandMap {
	if len(strings.Split(field, ".")) > 6 {
		return models.ExpandMap{}
	}
	if mapping == nil {
		mapping = models.ExpandMap{}
	}
	dotPos := strings.Index(field, ".")
	if dotPos <= 1 {
		mapping[field] = models.ExpandMap{}
		return mapping
	}
	target := field[:dotPos]
	child := field[(dotPos + 1):]
	mapping[target] = flatToExpandMap(mapping[target], child)
	return mapping
}

func ExtractExpandMap(c echo.Context) (models.ExpandMap, error) {
	type expandField struct {
		Expand models.ExpandMap `json:"expand"`
	}
	expandJson := c.QueryParam("expand")
	if expandJson == "" {
		return nil, nil
	}
	if !strings.HasPrefix(expandJson, "{") {
		expand := make(models.ExpandMap)
		for _, field := range strings.Split(expandJson, ",") {
			expand = flatToExpandMap(expand, field)
		}
		return expand, nil
	}
	expand := make(models.ExpandMap)
	if err := json.NewDecoder(strings.NewReader(expandJson)).Decode(&expand); err != nil {
		return nil, err
	}
	return expand, nil
}
