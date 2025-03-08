package hooks

import (
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/tools"
)

const (
	imageCoversField           = "covers"
	imageCoverField            = "cover"
	imageCoversCollectionField = "parentCollection"
	imageCoversRecordField     = "parentId"
)

func registerAppendImageSecretHook(
	app *pocketbase.PocketBase,
) error {
	targetCollections := []string{"titles", "books", "publications", "bookDetails", "releaseDetails", "titleCovers"}
	secret := app.Store().Get(models.ImagorSecretKey).(string)

	app.OnRecordViewRequest(targetCollections...).BindFunc(func(e *core.RecordRequestEvent) error {
		if err := appendImageSecret(secret, e.Record); err != nil {
			return err
		}
		return e.Next()
	})

	app.OnRecordsListRequest(targetCollections...).BindFunc(func(e *core.RecordsListRequestEvent) error {
		for _, record := range e.Records {
			if err := appendImageSecret(secret, record); err != nil {
				return err
			}
		}

		return nil
	})

	return nil
}

func appendImageSecret(secret string, record *core.Record) error {
	if _, exist := record.FieldsData()[imageCoversField]; exist {
		return appendImageSliceSecret(secret, record)
	}

	path := getCoverImagePath(record)

	if path == "" {
		return nil
	}

	appendMetadata(
		record,
		map[string]interface{}{
			"images": getImageSizes(secret, path),
		},
	)
	return nil
}

func appendImageSliceSecret(secret string, record *core.Record) error {
	covers := record.GetStringSlice(imageCoversField)

	var images []map[string]string

	for _, cover := range covers {
		path := getCoverImagePath(record, cover)

		images = append(images, getImageSizes(secret, path))
	}

	appendMetadata(
		record,
		map[string]interface{}{
			"images": images,
		},
	)
	return nil
}

// Return the cover image path from a record.
// Since PocketBase is having a bug on resolving type with every view record that using UNION,
//
//	(https://github.com/pocketbase/pocketbase/discussions/1938#discussioncomment-5143723)
//	we implement a temporary fix by removing the opening and ending double-quote.
func getCoverImagePath(record *core.Record, originalCover ...string) string {
	cover := ""
	if len(originalCover) <= 0 {
		cover = record.GetString(imageCoverField)
	} else {
		cover = originalCover[0]
	}
	if cover == "" {
		return ""
	}

	collectionId := record.Collection().Id
	id := record.Id

	if id := record.GetString(imageCoversCollectionField); id != "" {
		collectionId = string(id[1 : len(id)-1])
	}

	if rId := record.GetString(imageCoversRecordField); rId != "" {
		id = rId
	}

	return fmt.Sprintf("%s/%s/%s", collectionId, id, cover)
}

func appendImageSizeMetadata(secret string, record *core.Record) error {
	covers := record.GetStringSlice(imageCoversField)

	var images []map[string]string

	for _, cover := range covers {
		path := fmt.Sprintf(
			"%s/%s/%s",
			record.GetString(imageCoversCollectionField),
			record.GetString(imageCoversRecordField),
			cover,
		)
		images = append(images, getImageSizes(secret, path))
	}

	appendMetadata(
		record,
		map[string]interface{}{
			"images": images,
		},
	)
	return nil
}

func getImageSizes(secret string, path string) map[string]string {
	return map[string]string{
		"160w":  tools.SignImageUrl(secret, fmt.Sprintf("120x0/filters:quality(90)/%s", path)),
		"320w":  tools.SignImageUrl(secret, fmt.Sprintf("320x0/filters:quality(90)/%s", path)),
		"480w":  tools.SignImageUrl(secret, fmt.Sprintf("480x0/filters:quality(90)/%s", path)),
		"640w":  tools.SignImageUrl(secret, fmt.Sprintf("640x0/filters:quality(90)/%s", path)),
		"1280w": tools.SignImageUrl(secret, fmt.Sprintf("1280x0/filters:quality(90)/%s", path)),
		"1920w": tools.SignImageUrl(secret, fmt.Sprintf("1920x0/filters:quality(90)/%s", path)),
	}
}
