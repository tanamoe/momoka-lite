package hooks

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/tools"
)

func registerResizeAssetImageHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	targetCollections := []string{(&models.Asset{}).TableName()}
	app.OnModelAfterCreate(targetCollections...).Add(func(e *core.ModelEvent) error {
		return resizeAssetImage(app, context, e.Model.GetId())
	})
	app.OnModelAfterUpdate(targetCollections...).Add(func(e *core.ModelEvent) error {
		return resizeAssetImage(app, context, e.Model.GetId())
	})
	app.OnModelAfterUpdate((&models.State{}).TableName()).Add(func(e *core.ModelEvent) error {
		targetStateIds := map[string]bool{
			models.ImagorSecretStateId:      true,
			models.AssetImageResizedStateId: true,
		}
		if _, exist := targetStateIds[e.Model.GetId()]; !exist {
			return nil
		}
		assets := []*models.Asset{}
		err := models.AssetQuery(app.Dao()).Select("id").All(&assets)
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		if err != nil {
			return err
		}
		for _, asset := range assets {
			if err := resizeAssetImage(app, context, asset.Id); err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func resizeAssetImage(
	app *pocketbase.PocketBase,
	context *models.AppContext,
	assetId string,
) error {
	dao := app.Dao()
	asset, err := models.FindAssetById(dao, assetId)
	if err != nil {
		return err
	}
	id, err := tools.GetCollectionId(dao, asset.TableName())
	if err != nil {
		return err
	}
	path := fmt.Sprintf(
		"%s/%s/%s",
		id,
		asset.Id,
		asset.Image,
	)
	asset.ResizedImage = types.JsonMap{}
	resizedImages := getImageSizes(context.ImagorSecret, path)
	for id, resizedPath := range resizedImages {
		asset.ResizedImage[id] = resizedPath
	}
	return dao.WithoutHooks().Save(asset)
}
