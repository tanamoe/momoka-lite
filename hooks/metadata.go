package hooks

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

func appendMetadata(m *models.Record, data map[string]interface{}) {
	rawJson := m.Get("metadata")
	if rawJson == nil {
		rawJson = types.JsonRaw{}
	}
	metadata := map[string]interface{}{}
	if len(metadata) > 2 {
		if err := json.Unmarshal(rawJson.(types.JsonRaw), &metadata); err != nil {
			panic(err)
		}
	}
	for key, value := range data {
		metadata[key] = value
	}
	m.Set("metadata", metadata)
}
