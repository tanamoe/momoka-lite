package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `[
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "y19nx09k",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "eac1kqsz",
						"max": 0,
						"min": 0,
						"name": "slug",
						"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "1vm4japp",
						"max": 7,
						"min": 4,
						"name": "color",
						"pattern": "^#(?:[0-9a-fA-F]{3}){1,2}$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"convertURLs": false,
						"hidden": false,
						"id": "k7dq29id",
						"maxSize": 0,
						"name": "decription",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "editor"
					},
					{
						"hidden": false,
						"id": "cwcaxuub",
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
						"name": "thumbnail",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": null,
						"type": "file"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "zpr3heo6mae3h1w",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_YSULLEh` + "`" + ` ON ` + "`" + `formats` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"name": "formats",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "s91oidzeo1xm4m7",
						"hidden": false,
						"id": "4tlee6c6",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "title",
						"presentable": true,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "lzdfyn1k",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "8nglstcz",
						"max": 0,
						"min": 0,
						"name": "type",
						"pattern": "",
						"presentable": true,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "gcb2iw3u",
						"name": "digital",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "bool"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "awceclka",
						"max": 0,
						"min": 0,
						"name": "disambiguation",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "2lrfiedkzjul4s1",
						"hidden": false,
						"id": "mtaohnx5",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "publisher",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "2lrfiedkzjul4s1",
						"hidden": false,
						"id": "w9p6emac",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "partner",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "1t7lpcuz",
						"maxSelect": 1,
						"name": "status",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"WAITING_FOR_APPROVAL",
							"REGISTERED",
							"LICENSED",
							"ON_GOING",
							"COMPLETED",
							"HIATUS",
							"CANCELLED"
						]
					},
					{
						"cascadeDelete": false,
						"collectionId": "fygd9716lhpgmur",
						"hidden": false,
						"id": "kwwcrtzn",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "front",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "fygd9716lhpgmur",
						"hidden": false,
						"id": "blogqwsq",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "banner",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "fygd9716lhpgmur",
						"hidden": false,
						"id": "fxrj5efe",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "logo",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "ju84js8w",
						"max": null,
						"min": null,
						"name": "old_id",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "3j32s2l7fdos1e4",
				"indexes": [],
				"listRule": "",
				"name": "releases",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": true,
						"collectionId": "guv9vnyfu5pdz9t",
						"hidden": false,
						"id": "wx5t7htt",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "publication",
						"presentable": true,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "m9wcv0mj",
						"max": 0,
						"min": 0,
						"name": "edition",
						"pattern": "",
						"presentable": true,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "arr8bmxa",
						"max": "",
						"min": "",
						"name": "publishDate",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "date"
					},
					{
						"hidden": false,
						"id": "n99n0fa3",
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
						"name": "covers",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": null,
						"type": "file"
					},
					{
						"hidden": false,
						"id": "6m7pzsej",
						"max": null,
						"min": null,
						"name": "price",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"convertURLs": false,
						"hidden": false,
						"id": "inz6maav",
						"maxSize": 0,
						"name": "note",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "editor"
					},
					{
						"hidden": false,
						"id": "nudhir82",
						"maxSize": 2000000,
						"name": "metadata",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "ifejwbve",
						"max": 0,
						"min": 0,
						"name": "old_id",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "mu2u4hp0vc4dle5",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_gZH4WB5` + "`" + ` ON ` + "`" + `books` + "`" + ` (` + "`" + `publication` + "`" + `)"
				],
				"listRule": "",
				"name": "books",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "3j32s2l7fdos1e4",
						"hidden": false,
						"id": "g4g08sqp",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "release",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "duzqx65s",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "zvh7s1mm",
						"max": 0,
						"min": 0,
						"name": "subtitle",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"convertURLs": false,
						"hidden": false,
						"id": "oysvkf4d",
						"maxSize": 0,
						"name": "description",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "editor"
					},
					{
						"hidden": false,
						"id": "5oc5bnk3",
						"max": null,
						"min": null,
						"name": "volume",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"cascadeDelete": false,
						"collectionId": "mu2u4hp0vc4dle5",
						"hidden": false,
						"id": "kb6sktmd",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "defaultBook",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "h0okjh8g",
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
						"name": "covers",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": null,
						"type": "file"
					},
					{
						"hidden": false,
						"id": "wgzhppl8",
						"maxSize": 2000000,
						"name": "metadata",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "vr9ftnmg",
						"max": 0,
						"min": 0,
						"name": "old_id",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "guv9vnyfu5pdz9t",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_jj5RsfT` + "`" + ` ON ` + "`" + `publications` + "`" + ` (` + "`" + `release` + "`" + `)"
				],
				"listRule": "",
				"name": "publications",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "aewdsjta",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "kdwverajgytgjpe",
				"indexes": [],
				"listRule": "",
				"name": "staffs",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "cqxzavfw",
						"max": 0,
						"min": 0,
						"name": "slugGroup",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "tlb30fgj",
						"max": 0,
						"min": 0,
						"name": "slug",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "axcok2ww",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"convertURLs": false,
						"hidden": false,
						"id": "khr7f9me",
						"maxSize": 0,
						"name": "description",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "editor"
					},
					{
						"cascadeDelete": false,
						"collectionId": "zpr3heo6mae3h1w",
						"hidden": false,
						"id": "oxs4pmme",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "format",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "anl7vmmb",
						"maxSelect": 1,
						"maxSize": 20971520,
						"mimeTypes": null,
						"name": "cover",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": [
							"80x120"
						],
						"type": "file"
					},
					{
						"cascadeDelete": false,
						"collectionId": "8msmj3ci8k33wbe",
						"hidden": false,
						"id": "tseu7q2w",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "demographic",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "walac4l9hx6i63v",
						"hidden": false,
						"id": "r8to7vei",
						"maxSelect": 2147483647,
						"minSelect": 0,
						"name": "genres",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "ptmy3urf",
						"maxSize": 2000000,
						"name": "metadata",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"cascadeDelete": false,
						"collectionId": "3j32s2l7fdos1e4",
						"hidden": false,
						"id": "nu1kgjh8",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "defaultRelease",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "s91oidzeo1xm4m7",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_gFgrqNg` + "`" + ` ON ` + "`" + `titles` + "`" + ` (` + "`" + `slugGroup` + "`" + `)",
					"CREATE INDEX ` + "`" + `idx_05tXG2O` + "`" + ` ON ` + "`" + `titles` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"name": "titles",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "e5b8x7mo",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "djiyotvx",
						"maxSelect": 1,
						"maxSize": 5242880,
						"mimeTypes": [
							"image/jpeg",
							"image/png",
							"image/svg+xml",
							"image/gif",
							"image/webp"
						],
						"name": "logo",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": [
							"24x24"
						],
						"type": "file"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "atfsttrk",
						"max": 0,
						"min": 0,
						"name": "slug",
						"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "w8uj8pzd",
						"max": 7,
						"min": 4,
						"name": "color",
						"pattern": "^#(?:[0-9a-fA-F]{3}){1,2}$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "2lrfiedkzjul4s1",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_RmOvURr` + "`" + ` ON ` + "`" + `publishers` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"name": "publishers",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 0,
						"min": 0,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "_clone_Vmhe",
						"max": "",
						"min": "",
						"name": "publishDate",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "date"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "_clone_vxcM",
						"max": 0,
						"min": 0,
						"name": "edition",
						"pattern": "",
						"presentable": true,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "_clone_WVfI",
						"max": null,
						"min": null,
						"name": "price",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"convertURLs": false,
						"hidden": false,
						"id": "_clone_y0rV",
						"maxSize": 0,
						"name": "note",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "editor"
					},
					{
						"hidden": false,
						"id": "json1326724116",
						"maxSize": 1,
						"name": "metadata",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"cascadeDelete": false,
						"collectionId": "guv9vnyfu5pdz9t",
						"hidden": false,
						"id": "relation2939971449",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "publication",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "3j32s2l7fdos1e4",
						"hidden": false,
						"id": "_clone_e7Na",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "release",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "json264759313",
						"maxSize": 1,
						"name": "parentCollection",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"cascadeDelete": false,
						"collectionId": "mu2u4hp0vc4dle5",
						"hidden": false,
						"id": "relation284052718",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "parentId",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "_clone_12hT",
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
						"name": "covers",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": null,
						"type": "file"
					},
					{
						"hidden": false,
						"id": "_clone_MPZ1",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "_clone_q36f",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "0l473ttmx8o31i9",
				"indexes": [],
				"listRule": "",
				"name": "bookDetails",
				"system": false,
				"type": "view",
				"updateRule": null,
				"viewQuery": "SELECT \n  books.id,\n  books.publishDate,\n  books.edition,\n  books.price,\n  books.note,\n  \"{}\" as metadata,\n  publications.id as publication,\n  publications.release,\n  \"publications\" as parentCollection,\n  publications.id as parentId,\n  publications.covers,\n  publications.created,\n  publications.updated\nFROM books\nLEFT JOIN publications ON books.publication = publications.id\nWHERE books.covers = \"[]\"\nUNION ALL\nSELECT\n  books.id,\n  books.publishDate,\n  books.edition,\n  books.price,\n  books.note,\n  \"{}\" as metadata,\n  publications.id as publication,\n  publications.release,\n  \"books\" as parentCollection,\n  books.id as parentId,\n  books.covers,\n  books.created,\n  books.updated\nFROM books\nLEFT JOIN publications ON books.publication = publications.id\nWHERE books.covers != \"[]\"",
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "_pb_users_auth_",
						"hidden": false,
						"id": "nfiyam1n",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "owner",
						"presentable": true,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "vq5nrwbv",
						"maxSelect": 1,
						"name": "visibility",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"PRIVATE",
							"UNLISTED",
							"PUBLIC"
						]
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "9ctkkqxa",
						"max": 0,
						"min": 1,
						"name": "name",
						"pattern": "",
						"presentable": true,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "mbjirlsm",
						"name": "default",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "bool"
					},
					{
						"convertURLs": false,
						"hidden": false,
						"id": "mvj1yde4",
						"maxSize": 0,
						"name": "description",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "editor"
					},
					{
						"hidden": false,
						"id": "c66s8lzq",
						"max": null,
						"min": 0,
						"name": "order",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "ldgikhnt12bt4a6",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_h3R3CeQ` + "`" + ` ON ` + "`" + `collections` + "`" + ` (` + "`" + `owner` + "`" + `)"
				],
				"listRule": null,
				"name": "collections",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": null
			},
			{
				"createRule": "user = @request.auth.id && @request.auth.verified = true",
				"deleteRule": "user = @request.auth.id",
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "3j32s2l7fdos1e4",
						"hidden": false,
						"id": "fwtyy7lx",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "release",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "_pb_users_auth_",
						"hidden": false,
						"id": "xboghilk",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "user",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "wup1bczp",
						"max": 0,
						"min": 0,
						"name": "header",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"convertURLs": false,
						"hidden": false,
						"id": "j68gxbbk",
						"maxSize": 0,
						"name": "content",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "editor"
					},
					{
						"hidden": false,
						"id": "xcwy9z37",
						"max": 10,
						"min": 0,
						"name": "score",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "wtd1x8mugo9liuw",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_ojmVl7c` + "`" + ` ON ` + "`" + `reviews` + "`" + ` (\n  ` + "`" + `release` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)",
					"CREATE INDEX ` + "`" + `idx_KazuTq0` + "`" + ` ON ` + "`" + `reviews` + "`" + ` (` + "`" + `release` + "`" + `)"
				],
				"listRule": "",
				"name": "reviews",
				"system": false,
				"type": "base",
				"updateRule": "user = @request.auth.id",
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "s91oidzeo1xm4m7",
						"hidden": false,
						"id": "ctbkfe27",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "title",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "kdwverajgytgjpe",
						"hidden": false,
						"id": "zxp2x5pq",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "staff",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "c5ta1rqb",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "pj5teplb",
						"max": null,
						"min": null,
						"name": "priority",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "6uk141b1jx0dkhu",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_RXT00L7` + "`" + ` ON ` + "`" + `works` + "`" + ` (` + "`" + `title` + "`" + `)"
				],
				"listRule": "",
				"name": "works",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "s91oidzeo1xm4m7",
						"hidden": false,
						"id": "vesku0uz",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "title",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "0jd7tc1qu0m84u8",
						"hidden": false,
						"id": "qztytmal",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "source",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"exceptDomains": null,
						"hidden": false,
						"id": "mkjeereu",
						"name": "url",
						"onlyDomains": null,
						"presentable": false,
						"required": true,
						"system": false,
						"type": "url"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "aq9t2qxia36sz22",
				"indexes": [],
				"listRule": "",
				"name": "links",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "ldgikhnt12bt4a6",
						"hidden": false,
						"id": "9louzrtd",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "collection",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "_pb_users_auth_",
						"hidden": false,
						"id": "c8hj1jzf",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "user",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "cbsgvhrm",
						"maxSelect": 1,
						"name": "role",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "select",
						"values": [
							"EDITOR",
							"MEMBER"
						]
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "380dcd8bylxiw0w",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_RCz3SdE` + "`" + ` ON ` + "`" + `collectionMembers` + "`" + ` (` + "`" + `collection` + "`" + `)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_7sz2GU1` + "`" + ` ON ` + "`" + `collectionMembers` + "`" + ` (\n  ` + "`" + `collection` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)"
				],
				"listRule": null,
				"name": "collectionMembers",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": null
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "ldgikhnt12bt4a6",
						"hidden": false,
						"id": "mptqkias",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "collection",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": true,
						"collectionId": "mu2u4hp0vc4dle5",
						"hidden": false,
						"id": "6wxoqtgr",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "book",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "3k6cwj6s",
						"max": null,
						"min": 1,
						"name": "quantity",
						"onlyInt": false,
						"presentable": false,
						"required": true,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "8huiodky",
						"maxSelect": 1,
						"name": "status",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "select",
						"values": [
							"PLANNING",
							"COMPLETED"
						]
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "barh8omf",
						"max": 0,
						"min": 0,
						"name": "notes",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "nkw6spljgatiiyx",
				"indexes": [],
				"listRule": null,
				"name": "collectionBooks",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": null
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "smxo8l5p",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "e5pu0xos",
						"max": 0,
						"min": 0,
						"name": "color",
						"pattern": "^#(?:[0-9a-fA-F]{3}){1,2}$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "x105n2lt",
						"maxSelect": 1,
						"maxSize": 5242880,
						"mimeTypes": [
							"image/png",
							"image/jpeg",
							"image/svg+xml",
							"image/webp"
						],
						"name": "icon",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": [
							"50x50"
						],
						"type": "file"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "0jd7tc1qu0m84u8",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_JpwCQ4l` + "`" + ` ON ` + "`" + `linkSources` + "`" + ` (` + "`" + `name` + "`" + `)"
				],
				"listRule": "",
				"name": "linkSources",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "9mkj3bxh",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "yxa6wanz",
						"max": 0,
						"min": 0,
						"name": "slug",
						"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "walac4l9hx6i63v",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_CtQKNT9` + "`" + ` ON ` + "`" + `genres` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"name": "genres",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 0,
						"min": 0,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "json4035834290",
						"maxSize": 1,
						"name": "covers",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"hidden": false,
						"id": "json724990059",
						"maxSize": 1,
						"name": "title",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"hidden": false,
						"id": "json264759313",
						"maxSize": 1,
						"name": "parentCollection",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"hidden": false,
						"id": "json3113930206",
						"maxSize": 1,
						"name": "volume",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"hidden": false,
						"id": "json1326724116",
						"maxSize": 1,
						"name": "metadata",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					}
				],
				"id": "tnyytazu9gvdxse",
				"indexes": [],
				"listRule": "",
				"name": "titleCovers",
				"system": false,
				"type": "view",
				"updateRule": null,
				"viewQuery": "SELECT \n  id,\n  covers, \n  title, \n  parentCollection,\n  volume,\n  \"{}\" as metadata\nFROM \n  (\n    SELECT \n      books.id as id, \n      books.covers as covers, \n      publications.volume as volume,\n      releases.title as title,\n      \"books\" as parentCollection \n    FROM \n      books \n      RIGHT JOIN publications ON publications.id = books.publication \n      RIGHT JOIN releases ON releases.id = publications.release \n    WHERE \n      books.covers != \"[]\" \n    UNION \n    SELECT \n      publications.id as id, \n      publications.covers as covers, \n      publications.volume as volume,\n      releases.title as title,\n      \"publications\" as parentCollection \n    FROM \n      publications \n      RIGHT JOIN releases ON releases.id = publications.release \n    WHERE \n      publications.covers != \"[]\"\n  )\nORDER BY title ASC, volume ASC, (CASE parentCollection\n  WHEN 'publications' THEN 0\n  ELSE 1\nEND) ASC;",
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "pnwz4ww0",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "dfd7yiaa",
						"max": 0,
						"min": 0,
						"name": "slug",
						"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "8msmj3ci8k33wbe",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_TRi2rPa` + "`" + ` ON ` + "`" + `demographics` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"name": "demographics",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 0,
						"min": 0,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "_clone_9Kbv",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "s91oidzeo1xm4m7",
						"hidden": false,
						"id": "_clone_TS0G",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "title",
						"presentable": true,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "2lrfiedkzjul4s1",
						"hidden": false,
						"id": "_clone_fxU0",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "publisher",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "_clone_LvDy",
						"maxSelect": 1,
						"name": "status",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"WAITING_FOR_APPROVAL",
							"REGISTERED",
							"LICENSED",
							"ON_GOING",
							"COMPLETED",
							"HIATUS",
							"CANCELLED"
						]
					},
					{
						"hidden": false,
						"id": "_clone_80Ei",
						"maxSelect": 1,
						"maxSize": 20971520,
						"mimeTypes": null,
						"name": "cover",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": [
							"80x120"
						],
						"type": "file"
					},
					{
						"hidden": false,
						"id": "_clone_kEEb",
						"maxSize": 2000000,
						"name": "metadata",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"hidden": false,
						"id": "_clone_sLTM",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "_clone_MkZ0",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "uw1kal8bb55bf1e",
				"indexes": [],
				"listRule": "",
				"name": "releaseDetails",
				"system": false,
				"type": "view",
				"updateRule": null,
				"viewQuery": "SELECT releases.id, releases.name, releases.title, releases.publisher, releases.status, titles.cover, titles.metadata, releases.created, releases.updated\nFROM releases, titles\nWHERE releases.title = titles.id",
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": true,
						"collectionId": "s91oidzeo1xm4m7",
						"hidden": false,
						"id": "yozmq8fz",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "title",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "k7s4a6td",
						"max": 0,
						"min": 0,
						"name": "language",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "uedowwct",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "wzjok6uyx3y1qiz",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_3Fo1G1g` + "`" + ` ON ` + "`" + `additionalTitleNames` + "`" + ` (` + "`" + `title` + "`" + `)"
				],
				"listRule": "",
				"name": "additionalTitleNames",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": true,
						"collectionId": "mu2u4hp0vc4dle5",
						"hidden": false,
						"id": "ttvregrx",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "book",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "zyecoebo",
						"max": 0,
						"min": 0,
						"name": "isbn",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "h0sq0kvc",
						"max": 0,
						"min": 0,
						"name": "fahasaSKU",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "iz7fd2s1",
						"max": null,
						"min": 0,
						"name": "sizeX",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "7mdnwwrb",
						"max": null,
						"min": 0,
						"name": "sizeY",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "orewr0zw",
						"max": null,
						"min": 0,
						"name": "sizeZ",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "jpp8a5f8",
						"max": null,
						"min": 1,
						"name": "pageCount",
						"onlyInt": true,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "rlbe2tdi",
						"max": null,
						"min": 0,
						"name": "weight",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "k3r456599a14nlb",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_sctUHKm` + "`" + ` ON ` + "`" + `bookMetadata` + "`" + ` (` + "`" + `book` + "`" + `)"
				],
				"listRule": "",
				"name": "bookMetadata",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "isshyiso",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "d6x2asie8rsmmqm",
				"indexes": [],
				"listRule": "",
				"name": "assetTypes",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "mu2u4hp0vc4dle5",
						"hidden": false,
						"id": "xzrmlltm",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "book",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "d6x2asie8rsmmqm",
						"hidden": false,
						"id": "lu21hg1m",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "type",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "rehsspul",
						"maxSelect": 1,
						"maxSize": 15728640,
						"mimeTypes": null,
						"name": "image",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": null,
						"type": "file"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "yw37p5ut",
						"max": 0,
						"min": 0,
						"name": "description",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "wp4crwfz",
						"maxSize": 2000000,
						"name": "resizedImage",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"hidden": false,
						"id": "ff51lttd",
						"max": null,
						"min": 0,
						"name": "priority",
						"onlyInt": true,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "fygd9716lhpgmur",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_uo7NyeJ` + "`" + ` ON ` + "`" + `assets` + "`" + ` (` + "`" + `book` + "`" + `)",
					"CREATE INDEX ` + "`" + `idx_zlyTsDZ` + "`" + ` ON ` + "`" + `assets` + "`" + ` (\n  ` + "`" + `book` + "`" + `,\n  ` + "`" + `type` + "`" + `\n)"
				],
				"listRule": "",
				"name": "assets",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": ""
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "hbqimlup",
						"max": 0,
						"min": 0,
						"name": "value",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "tflqwa7u8grbgqp",
				"indexes": [],
				"listRule": null,
				"name": "states",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": null
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 0,
						"min": 0,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "_clone_hZwq",
						"max": "",
						"min": "",
						"name": "publishDate",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "date"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "_clone_xbW7",
						"max": 0,
						"min": 0,
						"name": "edition",
						"pattern": "",
						"presentable": true,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "_clone_Xxka",
						"max": null,
						"min": null,
						"name": "price",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"convertURLs": false,
						"hidden": false,
						"id": "_clone_xc49",
						"maxSize": 0,
						"name": "note",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "editor"
					},
					{
						"hidden": false,
						"id": "_clone_etm8",
						"max": null,
						"min": null,
						"name": "volume",
						"onlyInt": false,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"hidden": false,
						"id": "_clone_mz9o",
						"name": "digital",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "bool"
					},
					{
						"cascadeDelete": false,
						"collectionId": "fygd9716lhpgmur",
						"hidden": false,
						"id": "relation2043772302",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "assets",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "k3r456599a14nlb",
						"hidden": false,
						"id": "relation1326724116",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "metadata",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "guv9vnyfu5pdz9t",
						"hidden": false,
						"id": "relation2939971449",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "publication",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "3j32s2l7fdos1e4",
						"hidden": false,
						"id": "_clone_Zp1b",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "release",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "s91oidzeo1xm4m7",
						"hidden": false,
						"id": "_clone_QsPj",
						"maxSelect": 1,
						"minSelect": 0,
						"name": "title",
						"presentable": true,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "_clone_h1Lm",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "_clone_0TRg",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "d2ak05ggk8sdscc",
				"indexes": [],
				"listRule": null,
				"name": "booksComplete",
				"system": false,
				"type": "view",
				"updateRule": null,
				"viewQuery": "SELECT \n  -- details\n  books.id,\n  books.publishDate,\n  books.edition,\n  books.price,\n  books.note,\n  publications.volume,\n  releases.digital as digital,\n  -- relations\n  assets.id as assets,\n  bookMetadata.id as metadata,\n  publications.id as publication,\n  publications.release,\n  releases.title,\n  -- time\n  books.created,\n  books.updated\nFROM books\nLEFT JOIN bookMetadata ON books.id = bookMetadata.book\nLEFT JOIN publications ON books.publication = publications.id\nLEFT JOIN assets ON ((publications.defaultBook = assets.book AND assets.priority = 1) OR (books.id = assets.book AND assets.priority = 1))\nLEFT JOIN releases ON publications.release = releases.id",
				"viewRule": null
			},
			{
				"authAlert": {
					"emailTemplate": {
						"body": "<p>Hello,</p>\n<p>We noticed a login to your {APP_NAME} account from a new location.</p>\n<p>If this was you, you may disregard this email.</p>\n<p><strong>If this wasn't you, you should immediately change your {APP_NAME} account password to revoke access from all other locations.</strong></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
						"subject": "Login from a new location"
					},
					"enabled": true
				},
				"authRule": "",
				"authToken": {
					"duration": 1209600
				},
				"confirmEmailChangeTemplate": {
					"body": "<p>Hello,</p>\n<p>Click on the button below to confirm your new email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-email-change/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Confirm new email</a>\n</p>\n<p><i>If you didn't ask to change your email address, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Confirm your {APP_NAME} new email address"
				},
				"createRule": "",
				"deleteRule": "id = @request.auth.id",
				"emailChangeToken": {
					"duration": 1800
				},
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cost": 10,
						"hidden": true,
						"id": "password901924565",
						"max": 0,
						"min": 8,
						"name": "password",
						"pattern": "",
						"presentable": false,
						"required": true,
						"system": true,
						"type": "password"
					},
					{
						"autogeneratePattern": "[a-zA-Z0-9_]{50}",
						"hidden": true,
						"id": "text2504183744",
						"max": 60,
						"min": 30,
						"name": "tokenKey",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"exceptDomains": null,
						"hidden": false,
						"id": "email3885137012",
						"name": "email",
						"onlyDomains": null,
						"presentable": false,
						"required": false,
						"system": true,
						"type": "email"
					},
					{
						"hidden": false,
						"id": "bool1547992806",
						"name": "emailVisibility",
						"presentable": false,
						"required": false,
						"system": true,
						"type": "bool"
					},
					{
						"hidden": false,
						"id": "bool256245529",
						"name": "verified",
						"presentable": false,
						"required": false,
						"system": true,
						"type": "bool"
					},
					{
						"autogeneratePattern": "users[0-9]{6}",
						"hidden": false,
						"id": "text4166911607",
						"max": 150,
						"min": 3,
						"name": "username",
						"pattern": "^[\\w][\\w\\.\\-]*$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "users_name",
						"max": 0,
						"min": 0,
						"name": "displayName",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "duuqo7tx",
						"max": 280,
						"min": 0,
						"name": "bio",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "users_avatar",
						"maxSelect": 1,
						"maxSize": 1048576,
						"mimeTypes": [
							"image/jpeg",
							"image/png",
							"image/webp"
						],
						"name": "avatar",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": [
							"32x32",
							"128x128"
						],
						"type": "file"
					},
					{
						"hidden": false,
						"id": "nlmcbsrx",
						"maxSelect": 1,
						"maxSize": 2097152,
						"mimeTypes": [
							"image/jpeg",
							"image/png",
							"image/webp"
						],
						"name": "banner",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": [
							"420x180"
						],
						"type": "file"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"fileToken": {
					"duration": 120
				},
				"id": "_pb_users_auth_",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `__pb_users_auth__username_idx` + "`" + ` ON ` + "`" + `users` + "`" + ` (username COLLATE NOCASE)",
					"CREATE UNIQUE INDEX ` + "`" + `__pb_users_auth__email_idx` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `email` + "`" + `) WHERE ` + "`" + `email` + "`" + ` != ''",
					"CREATE UNIQUE INDEX ` + "`" + `__pb_users_auth__tokenKey_idx` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `tokenKey` + "`" + `)"
				],
				"listRule": "",
				"manageRule": null,
				"mfa": {
					"duration": 1800,
					"enabled": false,
					"rule": ""
				},
				"name": "users",
				"oauth2": {
					"enabled": false,
					"mappedFields": {
						"avatarURL": "",
						"id": "",
						"name": "",
						"username": "username"
					}
				},
				"otp": {
					"duration": 180,
					"emailTemplate": {
						"body": "<p>Hello,</p>\n<p>Your one-time password is: <strong>{OTP}</strong></p>\n<p><i>If you didn't ask for the one-time password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
						"subject": "OTP for {APP_NAME}"
					},
					"enabled": false,
					"length": 8
				},
				"passwordAuth": {
					"enabled": true,
					"identityFields": [
						"email",
						"username"
					]
				},
				"passwordResetToken": {
					"duration": 1800
				},
				"resetPasswordTemplate": {
					"body": "<p>Hello,</p>\n<p>Click on the button below to reset your password.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-password-reset/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Reset password</a>\n</p>\n<p><i>If you didn't ask to reset your password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Reset your {APP_NAME} password"
				},
				"system": false,
				"type": "auth",
				"updateRule": "id = @request.auth.id",
				"verificationTemplate": {
					"body": "<p>Hello,</p>\n<p>Thank you for joining us at {APP_NAME}.</p>\n<p>Click on the button below to verify your email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-verification/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Verify</a>\n</p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Verify your {APP_NAME} email"
				},
				"verificationToken": {
					"duration": 604800
				},
				"viewRule": ""
			},
			{
				"authAlert": {
					"emailTemplate": {
						"body": "<p>Hello,</p>\n<p>We noticed a login to your {APP_NAME} account from a new location.</p>\n<p>If this was you, you may disregard this email.</p>\n<p><strong>If this wasn't you, you should immediately change your {APP_NAME} account password to revoke access from all other locations.</strong></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
						"subject": "Login from a new location"
					},
					"enabled": true
				},
				"authRule": "",
				"authToken": {
					"duration": 1209600
				},
				"confirmEmailChangeTemplate": {
					"body": "<p>Hello,</p>\n<p>Click on the button below to confirm your new email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-email-change/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Confirm new email</a>\n</p>\n<p><i>If you didn't ask to change your email address, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Confirm your {APP_NAME} new email address"
				},
				"createRule": null,
				"deleteRule": null,
				"emailChangeToken": {
					"duration": 1800
				},
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cost": 0,
						"hidden": true,
						"id": "password901924565",
						"max": 0,
						"min": 8,
						"name": "password",
						"pattern": "",
						"presentable": false,
						"required": true,
						"system": true,
						"type": "password"
					},
					{
						"autogeneratePattern": "[a-zA-Z0-9]{50}",
						"hidden": true,
						"id": "text2504183744",
						"max": 60,
						"min": 30,
						"name": "tokenKey",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"exceptDomains": null,
						"hidden": false,
						"id": "email3885137012",
						"name": "email",
						"onlyDomains": null,
						"presentable": false,
						"required": true,
						"system": true,
						"type": "email"
					},
					{
						"hidden": false,
						"id": "bool1547992806",
						"name": "emailVisibility",
						"presentable": false,
						"required": false,
						"system": true,
						"type": "bool"
					},
					{
						"hidden": false,
						"id": "bool256245529",
						"name": "verified",
						"presentable": false,
						"required": false,
						"system": true,
						"type": "bool"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"fileToken": {
					"duration": 120
				},
				"id": "pbc_3142635823",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_tokenKey_pbc_3142635823` + "`" + ` ON ` + "`" + `_superusers` + "`" + ` (` + "`" + `tokenKey` + "`" + `)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_email_pbc_3142635823` + "`" + ` ON ` + "`" + `_superusers` + "`" + ` (` + "`" + `email` + "`" + `) WHERE ` + "`" + `email` + "`" + ` != ''"
				],
				"listRule": null,
				"manageRule": null,
				"mfa": {
					"duration": 1800,
					"enabled": false,
					"rule": ""
				},
				"name": "_superusers",
				"oauth2": {
					"enabled": false,
					"mappedFields": {
						"avatarURL": "",
						"id": "",
						"name": "",
						"username": ""
					}
				},
				"otp": {
					"duration": 180,
					"emailTemplate": {
						"body": "<p>Hello,</p>\n<p>Your one-time password is: <strong>{OTP}</strong></p>\n<p><i>If you didn't ask for the one-time password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
						"subject": "OTP for {APP_NAME}"
					},
					"enabled": false,
					"length": 8
				},
				"passwordAuth": {
					"enabled": true,
					"identityFields": [
						"email"
					]
				},
				"passwordResetToken": {
					"duration": 1800
				},
				"resetPasswordTemplate": {
					"body": "<p>Hello,</p>\n<p>Click on the button below to reset your password.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-password-reset/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Reset password</a>\n</p>\n<p><i>If you didn't ask to reset your password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Reset your {APP_NAME} password"
				},
				"system": true,
				"type": "auth",
				"updateRule": null,
				"verificationTemplate": {
					"body": "<p>Hello,</p>\n<p>Thank you for joining us at {APP_NAME}.</p>\n<p>Click on the button below to verify your email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-verification/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Verify</a>\n</p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Verify your {APP_NAME} email"
				},
				"verificationToken": {
					"duration": 259200
				},
				"viewRule": null
			},
			{
				"createRule": null,
				"deleteRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text455797646",
						"max": 0,
						"min": 0,
						"name": "collectionRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text127846527",
						"max": 0,
						"min": 0,
						"name": "recordRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text2462348188",
						"max": 0,
						"min": 0,
						"name": "provider",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text1044722854",
						"max": 0,
						"min": 0,
						"name": "providerId",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"id": "pbc_2281828961",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_externalAuths_record_provider` + "`" + ` ON ` + "`" + `_externalAuths` + "`" + ` (collectionRef, recordRef, provider)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_externalAuths_collection_provider` + "`" + ` ON ` + "`" + `_externalAuths` + "`" + ` (collectionRef, provider, providerId)"
				],
				"listRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"name": "_externalAuths",
				"system": true,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId"
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text455797646",
						"max": 0,
						"min": 0,
						"name": "collectionRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text127846527",
						"max": 0,
						"min": 0,
						"name": "recordRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text1582905952",
						"max": 0,
						"min": 0,
						"name": "method",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"id": "pbc_2279338944",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_mfas_collectionRef_recordRef` + "`" + ` ON ` + "`" + `_mfas` + "`" + ` (collectionRef,recordRef)"
				],
				"listRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"name": "_mfas",
				"system": true,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId"
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text455797646",
						"max": 0,
						"min": 0,
						"name": "collectionRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text127846527",
						"max": 0,
						"min": 0,
						"name": "recordRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cost": 8,
						"hidden": true,
						"id": "password901924565",
						"max": 0,
						"min": 0,
						"name": "password",
						"pattern": "",
						"presentable": false,
						"required": true,
						"system": true,
						"type": "password"
					},
					{
						"autogeneratePattern": "",
						"hidden": true,
						"id": "text3866985172",
						"max": 0,
						"min": 0,
						"name": "sentTo",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"id": "pbc_1638494021",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_otps_collectionRef_recordRef` + "`" + ` ON ` + "`" + `_otps` + "`" + ` (collectionRef, recordRef)"
				],
				"listRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"name": "_otps",
				"system": true,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId"
			},
			{
				"createRule": null,
				"deleteRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text455797646",
						"max": 0,
						"min": 0,
						"name": "collectionRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text127846527",
						"max": 0,
						"min": 0,
						"name": "recordRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"hidden": false,
						"id": "text4228609354",
						"max": 0,
						"min": 0,
						"name": "fingerprint",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"id": "pbc_4275539003",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_authOrigins_unique_pairs` + "`" + ` ON ` + "`" + `_authOrigins` + "`" + ` (collectionRef, recordRef, fingerprint)"
				],
				"listRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"name": "_authOrigins",
				"system": true,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId"
			}
		]`

		return app.ImportCollectionsByMarshaledJSON([]byte(jsonData), false)
	}, func(app core.App) error {
		return nil
	})
}
