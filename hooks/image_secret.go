package hooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/models"
	"tana.moe/momoka-lite/models"
)

func registerAppendImageSecretHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	secret := context.ImagorSecret

	app.OnRecordViewRequest("titles").Add(func(e *core.RecordViewEvent) error {
		return appendImageSecret(secret, e.Record)
	})

	app.OnRecordsListRequest("titles").Add(func(e *core.RecordsListEvent) error {
		for _, record := range e.Records {
			if err := appendImageSecret(secret, record); err != nil {
				return err
			}
		}

		return nil
	})

	app.OnRecordViewRequest("books", "publications").Add(func(e *core.RecordViewEvent) error {
		return appendImageSliceSecret(secret, e.Record)
	})

	app.OnRecordsListRequest("books", "publications").Add(func(e *core.RecordsListEvent) error {
		for _, record := range e.Records {
			if err := appendImageSliceSecret(secret, record); err != nil {
				return err
			}
		}

		return nil
	})

	app.OnRecordViewRequest("bookDetails").Add(func(e *core.RecordViewEvent) error {
		return appendImageSliceSecret(secret, e.Record)
	})

	app.OnRecordsListRequest("bookDetails").Add(func(e *core.RecordsListEvent) error {
		for _, record := range e.Records {
			if err := appendImageSliceSecret(secret, record); err != nil {
				return err
			}
		}

		return nil
	})

	return nil
}

func appendImageSecret(secret string, record *m.Record) error {
	cover := record.GetString("cover")

	if cover == "" {
		return nil
	}

	path := record.Collection().Id + "/" + record.Id + "/" + cover

	record.Set("metadata", map[string]interface{}{
		"images": getImageSizes(secret, path),
	},
	)

	return nil
}

func appendImageSliceSecret(secret string, record *m.Record) error {
	covers := record.GetStringSlice("covers")

	var images []map[string]string

	for _, cover := range covers {
		path := record.Collection().Id + "/" + record.Id + "/" + cover
		images = append(images, getImageSizes(secret, path))
	}

	record.Set("metadata", map[string]interface{}{
		"images": images,
	})

	return nil
}

func appendImageSizeMetadata(secret string, record *m.Record) error {
	covers := record.GetStringSlice("covers")

	var images []map[string]string

	for _, cover := range covers {
		path := record.GetString("coversCollection") + "/" + record.GetString("coversRecord") + "/" + cover
		images = append(images, getImageSizes(secret, path))
	}

	record.Set("metadata", map[string]interface{}{
		"images": images,
	})

	return nil
}

func getImageSizes(secret string, path string) map[string]string {
	return map[string]string{
		"160w":  signImageHmac(secret, "120x0/filters:quality(90)/"+path),
		"320w":  signImageHmac(secret, "320x0/filters:quality(90)/"+path),
		"480w":  signImageHmac(secret, "480x0/filters:quality(90)/"+path),
		"640w":  signImageHmac(secret, "640x0/filters:quality(90)/"+path),
		"1280w": signImageHmac(secret, "1280x0/filters:quality(90)/"+path),
		"1920w": signImageHmac(secret, "1920x0/filters:quality(90)/"+path),
	}
}

func signImageHmac(secret string, path string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(path))

	return base64.URLEncoding.EncodeToString(mac.Sum(nil))[:40] + "/" + path
}
