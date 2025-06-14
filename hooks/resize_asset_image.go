package hooks

import (
	"database/sql"
	"errors"
	"path"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"tana.moe/momoka-lite/models"
)

func registerResizeAssetImageHook(
	app *pocketbase.PocketBase,
) error {
	targetCollections := []string{(&models.Asset{}).TableName()}
	app.OnModelAfterCreateSuccess(targetCollections...).BindFunc(func(e *core.ModelEvent) error {
		if err := resizeAssetImage(app, e.Model.PK().(string)); err != nil {
			return err
		}
		return e.Next()
	})
	app.OnModelAfterUpdateSuccess(targetCollections...).BindFunc(func(e *core.ModelEvent) error {
		if err := resizeAssetImage(app, e.Model.PK().(string)); err != nil {
			return err
		}
		return e.Next()
	})
	app.OnModelAfterUpdateSuccess((&models.State{}).TableName()).BindFunc(func(e *core.ModelEvent) error {
		targetStateIds := map[string]bool{
			models.ImagorSecretStateId:      true,
			models.AssetImageResizedStateId: true,
		}
		if _, exist := targetStateIds[e.Model.PK().(string)]; !exist {
			return e.Next()
		}
		assets := []*models.Asset{}
		err := models.AssetQuery(app.DB()).Select("id").All(&assets)
		if errors.Is(err, sql.ErrNoRows) {
			return e.Next()
		}
		if err != nil {
			return err
		}
		for _, asset := range assets {
			if err := resizeAssetImage(app, asset.Id); err != nil {
				return err
			}
		}
		return e.Next()
	})
	return nil
}

func resizeAssetImage(
	app *pocketbase.PocketBase,
	assetId string,
) error {
	db := app.UnsafeWithoutHooks().DB()
	asset, err := models.FindAssetById(db, assetId)
	if err != nil {
		return err
	}
	assetPath := path.Join(asset.Id, asset.Image)
	asset.ResizedImage = types.JSONMap[string]{}
	resizedImages := getImageSizes(app.Store().Get(models.ImagorSecretKey).(string), assetPath)
	for id, resizedPath := range resizedImages {
		asset.ResizedImage[id] = resizedPath
	}
	return db.Model(asset).Update()
}
