package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-07-17 03:36:06.154Z",
				"updated": "2023-08-11 03:05:01.566Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "displayName",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "duuqo7tx",
						"name": "bio",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 280,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 1048576,
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/webp"
							],
							"thumbs": [
								"24x24"
							],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "nlmcbsrx",
						"name": "banner",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 2097152,
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/webp"
							],
							"thumbs": [
								"420x180"
							],
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "zpr3heo6mae3h1w",
				"created": "2023-08-11 02:48:18.818Z",
				"updated": "2023-08-11 03:05:01.621Z",
				"name": "format",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "y19nx09k",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "eac1kqsz",
						"name": "slug",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
						}
					},
					{
						"system": false,
						"id": "1vm4japp",
						"name": "color",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 4,
							"max": 7,
							"pattern": "^#(?:[0-9a-fA-F]{3}){1,2}$"
						}
					},
					{
						"system": false,
						"id": "k7dq29id",
						"name": "decription",
						"type": "editor",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "cwcaxuub",
						"name": "thumbnail",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/png",
								"image/vnd.mozilla.apng",
								"image/jpeg",
								"image/jxl",
								"image/jp2",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_YSULLEh` + "`" + ` ON ` + "`" + `format` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "3j32s2l7fdos1e4",
				"created": "2023-08-11 02:48:18.834Z",
				"updated": "2023-08-11 03:05:01.670Z",
				"name": "release",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "4tlee6c6",
						"name": "title",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "s91oidzeo1xm4m7",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "8nglstcz",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mtaohnx5",
						"name": "publisher",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "2lrfiedkzjul4s1",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "1t7lpcuz",
						"name": "status",
						"type": "select",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"WAITING_FOR_APPROVAL",
								"REGISTERED",
								"LICENSED",
								"ON_GOING",
								"COMPLETED",
								"HIATUS",
								"CANCELLED"
							]
						}
					},
					{
						"system": false,
						"id": "ju84js8w",
						"name": "old_id",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "mu2u4hp0vc4dle5",
				"created": "2023-08-11 02:48:18.835Z",
				"updated": "2023-08-11 03:05:01.712Z",
				"name": "book",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wx5t7htt",
						"name": "publication",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "guv9vnyfu5pdz9t",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "m9wcv0mj",
						"name": "edition",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "arr8bmxa",
						"name": "publishDate",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "n99n0fa3",
						"name": "cover",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99,
							"maxSize": 20971520,
							"mimeTypes": [
								"image/png",
								"image/vnd.mozilla.apng",
								"image/jpeg",
								"image/jxl",
								"image/jp2",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "6m7pzsej",
						"name": "price",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "nudhir82",
						"name": "metadata",
						"type": "json",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "ifejwbve",
						"name": "old_id",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_gZH4WB5` + "`" + ` ON ` + "`" + `book` + "`" + ` (` + "`" + `publication` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "guv9vnyfu5pdz9t",
				"created": "2023-08-11 02:48:18.835Z",
				"updated": "2023-08-11 03:05:01.714Z",
				"name": "publication",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "g4g08sqp",
						"name": "release",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "3j32s2l7fdos1e4",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"name"
							]
						}
					},
					{
						"system": false,
						"id": "duzqx65s",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "5oc5bnk3",
						"name": "volume",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "h0okjh8g",
						"name": "cover",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99,
							"maxSize": 20971520,
							"mimeTypes": [
								"image/png",
								"image/vnd.mozilla.apng",
								"image/jpeg",
								"image/jxl",
								"image/jp2",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "joaiicfj",
						"name": "digital",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "wgzhppl8",
						"name": "metadata",
						"type": "json",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "vr9ftnmg",
						"name": "old_id",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_jj5RsfT` + "`" + ` ON ` + "`" + `publication` + "`" + ` (` + "`" + `release` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "kdwverajgytgjpe",
				"created": "2023-08-11 02:48:18.840Z",
				"updated": "2023-08-11 03:05:01.715Z",
				"name": "author",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "aewdsjta",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "s91oidzeo1xm4m7",
				"created": "2023-08-11 02:48:18.841Z",
				"updated": "2023-08-11 03:05:01.715Z",
				"name": "title",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "axcok2ww",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "khr7f9me",
						"name": "description",
						"type": "editor",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "sllntio1",
						"name": "author",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "kdwverajgytgjpe",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "oxs4pmme",
						"name": "format",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "zpr3heo6mae3h1w",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "anl7vmmb",
						"name": "cover",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [],
							"thumbs": [],
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "2lrfiedkzjul4s1",
				"created": "2023-08-11 02:48:18.890Z",
				"updated": "2023-08-11 03:05:01.759Z",
				"name": "publisher",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "e5b8x7mo",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "djiyotvx",
						"name": "logo",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": [
								"24x24"
							],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "atfsttrk",
						"name": "slug",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
						}
					},
					{
						"system": false,
						"id": "w8uj8pzd",
						"name": "color",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 4,
							"max": 7,
							"pattern": "^#(?:[0-9a-fA-F]{3}){1,2}$"
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_RmOvURr` + "`" + ` ON ` + "`" + `publisher` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "0kmdlx1ukh9p454",
				"created": "2023-08-11 02:55:50.444Z",
				"updated": "2023-08-11 03:05:01.968Z",
				"name": "book_data",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "cbneknte",
						"name": "titleId",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "s91oidzeo1xm4m7",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "9zwom33y",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "hujprnvw",
						"name": "description",
						"type": "editor",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "40hiioka",
						"name": "format",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "zpr3heo6mae3h1w",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "yrycjpwn",
						"name": "publicationName",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "tlehvbwy",
						"name": "baseCover",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99,
							"maxSize": 20971520,
							"mimeTypes": [
								"image/png",
								"image/vnd.mozilla.apng",
								"image/jpeg",
								"image/jxl",
								"image/jp2",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "nkkzntnp",
						"name": "cover",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99,
							"maxSize": 20971520,
							"mimeTypes": [
								"image/png",
								"image/vnd.mozilla.apng",
								"image/jpeg",
								"image/jxl",
								"image/jp2",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "yuruzup4",
						"name": "publishDate",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT\n  publication.id,\n  title.id AS titleId,\n  title.name,\n  title.description,\n  title.format,\n  publication.name AS publicationName,\n  publication.cover baseCover,\n  book.cover,\n  book.publishDate,\n  title.created,\n  title.updated\nFROM publication\nLEFT JOIN book ON book.publication = publication.id\nJOIN release ON release.id = publication.release\nJOIN title ON title.id = release.title\nORDER BY (CASE publication.cover\n  WHEN '' THEN 0\n  WHEN NULL THEN 0\n  WHEN '[]' THEN 0\n  ELSE 1\nEND) DESC,\npublishDate ASC;"
				}
			},
			{
				"id": "0l473ttmx8o31i9",
				"created": "2023-08-11 03:05:01.810Z",
				"updated": "2023-08-11 03:05:02.022Z",
				"name": "book_detailed",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "fcjnc9ds",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "psyhky5u",
						"name": "volume",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "a4tjyrsh",
						"name": "release",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "3j32s2l7fdos1e4",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"name"
							]
						}
					},
					{
						"system": false,
						"id": "var1kuv8",
						"name": "publishDate",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "vrhymccs",
						"name": "baseCover",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99,
							"maxSize": 20971520,
							"mimeTypes": [
								"image/png",
								"image/vnd.mozilla.apng",
								"image/jpeg",
								"image/jxl",
								"image/jp2",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "luv3auxt",
						"name": "cover",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99,
							"maxSize": 20971520,
							"mimeTypes": [
								"image/png",
								"image/vnd.mozilla.apng",
								"image/jpeg",
								"image/jxl",
								"image/jp2",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "aje5dx54",
						"name": "edition",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "4d6yohkj",
						"name": "price",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "qxjxzsob",
						"name": "digital",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "srun4ibb",
						"name": "publisher",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "2lrfiedkzjul4s1",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "uez1288b",
						"name": "title",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "s91oidzeo1xm4m7",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT \n  book.id, \n  publication.name, \n  publication.volume,\n  publication.release,\n  book.publishDate, \n  publication.cover as baseCover, \n  book.cover as cover, \n  book.edition, \n  book.price,\n  publication.digital,\n  release.publisher,\n  release.title,\n  book.created,\n  book.updated\nFROM book, publication, release\nWHERE \n  book.publication = publication.id AND \n  publication.release = release.id\nORDER BY \nbook.publishDate ASC,\n(CASE book.edition\n  WHEN '' THEN 0\n  ELSE 1\nEND) ASC;"
				}
			},
			{
				"id": "rdtjoq5nhzk9kgj",
				"created": "2023-08-11 03:05:01.812Z",
				"updated": "2023-08-11 03:05:02.067Z",
				"name": "title_with_cover",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wsw5khye",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "op2ckrsc",
						"name": "description",
						"type": "editor",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "heqnk1l5",
						"name": "format",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "zpr3heo6mae3h1w",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "ocbcvo7t",
						"name": "baseCover",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 99,
							"maxSize": 20971520,
							"mimeTypes": [
								"image/png",
								"image/vnd.mozilla.apng",
								"image/jpeg",
								"image/jxl",
								"image/jp2",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": [],
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT\n  titleId AS id,\n  name,\n  description,\n  format,\n  baseCover,\n  created,\n  updated\nFROM book_data\nGROUP BY titleId;"
				}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
