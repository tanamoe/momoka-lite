package hooks

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

func appendMetadata(m *core.Record, data map[string]interface{}) {
	rawJson := m.Get("metadata")
	if rawJson == nil {
		rawJson = types.JSONRaw{}
	}
	metadata := map[string]interface{}{}
	if len(rawJson.(types.JSONRaw)) > 2 {
		if err := json.Unmarshal(rawJson.(types.JSONRaw), &metadata); err != nil {
			panic(err)
		}
		if metadata == nil {
			metadata = map[string]interface{}{}
		}
	}
	for key, value := range data {
		metadata[key] = value
	}
	m.Set("metadata", metadata)
}
