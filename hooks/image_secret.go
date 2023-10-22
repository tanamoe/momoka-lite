package hooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/models"
	"tana.moe/momoka-lite/models"
)

func registerAppendImageSecretHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
        // check for imagor secret
        if _, ok := os.LookupEnv("IMAGOR_SECRET"); !ok {
                return errors.New("IMAGOR_SECRET not set")
        }

        app.OnRecordViewRequest("titles").Add(func(e *core.RecordViewEvent) error {
                return appendImageSecretHook(e.Record)
        })

	app.OnRecordsListRequest("titles").Add(func(e *core.RecordsListEvent) error {
                for _, record := range e.Records {
                    if err := appendImageSecretHook(record); err != nil {
                        return err
                    }
                }

                return nil
	})
        
        app.OnRecordViewRequest("books", "publications").Add(func(e *core.RecordViewEvent) error {
                return appendImageArraySecretHook(e.Record)
        })

        app.OnRecordsListRequest("books", "publications").Add(func(e *core.RecordsListEvent) error {
                for _, record := range e.Records {
                    if err := appendImageArraySecretHook(record); err != nil {
                        return err
                    }
                }

                return nil
            })
        
	return nil
}

func appendImageSecretHook(record *m.Record) error {
        cover := record.GetString("cover")

        if cover == "" {
                return nil
        }

        path := record.Collection().Id + "/" + record.Id + "/" + cover

        record.Set("metadata", map[string]interface{}{
                "coverSizes": getCoverSizes(path),
            },
        )

        return nil
}

func appendImageArraySecretHook(record *m.Record) error {
        covers := record.GetStringSlice("covers")

        var coverSizes []map[string]string

        for _, cover := range covers {
            coverSizes = append(coverSizes, getCoverSizes(cover))
        }
        
        record.Set("metadata", map[string]interface{}{
            "coverSizes": coverSizes,
        })

        return nil
}

func getCoverSizes(path string) map[string]string {
        return map[string]string{
                "160w": signImageHmac("120x0/filters:quality(90)/" + path),
                "320w": signImageHmac("320x0/filters:quality(90)/" + path),
                "480w": signImageHmac("480x0/filters:quality(90)/" + path),
                "640w": signImageHmac("640x0/filters:quality(90)/" + path),
                "1280w": signImageHmac("1280x0/filters:quality(90)/" + path),
                "1920w": signImageHmac("1920x0/filters:quality(90)/" + path),
        }
}

func signImageHmac(path string) string {
        secret := os.Getenv("IMAGOR_SECRET")

        mac := hmac.New(sha256.New, []byte(secret))
        mac.Write([]byte(path))

        return base64.URLEncoding.EncodeToString(mac.Sum(nil))[:40] + "/" + path
}
