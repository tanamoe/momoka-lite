package hooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/models"
	"tana.moe/momoka-lite/models"
)

const (
	imageCoversField           = "covers"
	imageCoverField            = "cover"
	imageCoversCollectionField = "coversCollection"
	imageCoversRecordField     = "coversRecord"
)

func registerAppendImageSecretHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	targetCollections := []string{"titles", "books", "publications", "bookDetails"}
	secret := context.ImagorSecret

	app.OnRecordViewRequest(targetCollections...).Add(func(e *core.RecordViewEvent) error {
		return appendImageSecret(secret, e.Record)
	})

	app.OnRecordsListRequest(targetCollections...).Add(func(e *core.RecordsListEvent) error {
		for _, record := range e.Records {
			if err := appendImageSecret(secret, record); err != nil {
				return err
			}
		}

		return nil
	})

	return nil
}

func appendImageSecret(secret string, record *m.Record) error {
	if _, exist := record.SchemaData()[imageCoversField]; exist {
		return appendImageSliceSecret(secret, record)
	}

	cover := record.GetString(imageCoverField)

	if cover == "" {
		return nil
	}

        collectionId := record.Collection().GetId()
        id := record.GetId()

        // cId will be return as "rawJSON", therefore [, ] and " will need to be removed
        if cId := record.GetString(imageCoversCollectionField); cId != "" {
                replacer := strings.NewReplacer("[", "", "]", "", "\"", "")
                collectionId = replacer.Replace(cId)
        }

        if rId := record.GetString(imageCoversRecordField); rId != "" {
                id = rId
        }

	path := fmt.Sprintf("%s/%s/%s", collectionId, id, cover)

	record.Set(
		"metadata",
		map[string]interface{}{
			"images": getImageSizes(secret, path),
		},
	)

	return nil
}

func appendImageSliceSecret(secret string, record *m.Record) error {
	covers := record.GetStringSlice(imageCoversField)

	var images []map[string]string

	for _, cover := range covers {
                collectionId := record.Collection().GetId()
                id := record.GetId()

                if cId := record.GetString(imageCoversCollectionField); cId != "" {
                        replacer := strings.NewReplacer("[", "", "]", "", "\"", "")
                        collectionId = replacer.Replace(cId)
                }

                if rId := record.GetString(imageCoversRecordField); rId != "" {
                        id = rId
                }

                path := fmt.Sprintf("%s/%s/%s", collectionId, id, cover)

                images = append(images, getImageSizes(secret, path))
	}

	record.Set(
		"metadata",
		map[string]interface{}{
			"images": images,
		},
	)

	return nil
}

func appendImageSizeMetadata(secret string, record *m.Record) error {
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

	record.Set("metadata", map[string]interface{}{
		"images": images,
	})

	return nil
}

func getImageSizes(secret string, path string) map[string]string {
	return map[string]string{
		"160w":  signImageUrl(secret, fmt.Sprintf("120x0/filters:quality(90)/%s", path)),
		"320w":  signImageUrl(secret, fmt.Sprintf("320x0/filters:quality(90)/%s", path)),
		"480w":  signImageUrl(secret, fmt.Sprintf("480x0/filters:quality(90)/%s", path)),
		"640w":  signImageUrl(secret, fmt.Sprintf("640x0/filters:quality(90)/%s", path)),
		"1280w": signImageUrl(secret, fmt.Sprintf("1280x0/filters:quality(90)/%s", path)),
		"1920w": signImageUrl(secret, fmt.Sprintf("1920x0/filters:quality(90)/%s", path)),
	}
}

func signImageUrl(secret string, path string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(path))
	sign := base64.URLEncoding.EncodeToString(mac.Sum(nil))[:40]
	return fmt.Sprintf("%s/%s", sign, path)
}
