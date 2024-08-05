package tools

import "github.com/pocketbase/pocketbase/daos"

func GetCollectionId(dao *daos.Dao, tableName string) (string, error) {
	collection, err := dao.FindCollectionByNameOrId(tableName)
	if err != nil {
		return "", err
	}
	return collection.Id, nil
}
