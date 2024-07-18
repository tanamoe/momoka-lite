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
				"created": "2023-06-27 13:23:10.160Z",
				"updated": "2024-07-15 17:00:44.289Z",
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
						"presentable": false,
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
						"presentable": false,
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
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/webp"
							],
							"thumbs": [
								"32x32",
								"128x128"
							],
							"maxSelect": 1,
							"maxSize": 1048576,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "nlmcbsrx",
						"name": "banner",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/webp"
							],
							"thumbs": [
								"420x180"
							],
							"maxSelect": 1,
							"maxSize": 2097152,
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
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
					"onlyVerified": false,
					"requireEmail": false
				}
			},
			{
				"id": "zpr3heo6mae3h1w",
				"created": "2023-07-01 02:59:19.734Z",
				"updated": "2024-07-15 17:00:44.290Z",
				"name": "formats",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "y19nx09k",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
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
						"presentable": false,
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
						"presentable": false,
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
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "cwcaxuub",
						"name": "thumbnail",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
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
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_YSULLEh` + "`" + ` ON ` + "`" + `formats` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "s91oidzeo1xm4m7",
				"created": "2023-07-01 03:02:58.825Z",
				"updated": "2024-07-15 17:00:44.290Z",
				"name": "titles",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "cqxzavfw",
						"name": "slugGroup",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "tlb30fgj",
						"name": "slug",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "axcok2ww",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
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
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "oxs4pmme",
						"name": "format",
						"type": "relation",
						"required": true,
						"presentable": false,
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
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [
								"80x120"
							],
							"maxSelect": 1,
							"maxSize": 20971520,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "tseu7q2w",
						"name": "demographic",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "8msmj3ci8k33wbe",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "r8to7vei",
						"name": "genres",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "walac4l9hx6i63v",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "ptmy3urf",
						"name": "metadata",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2000000
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_gFgrqNg` + "`" + ` ON ` + "`" + `titles` + "`" + ` (` + "`" + `slugGroup` + "`" + `)",
					"CREATE INDEX ` + "`" + `idx_05tXG2O` + "`" + ` ON ` + "`" + `titles` + "`" + ` (` + "`" + `slug` + "`" + `)"
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
				"created": "2023-07-01 03:06:45.488Z",
				"updated": "2024-07-15 17:00:44.290Z",
				"name": "publications",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "g4g08sqp",
						"name": "release",
						"type": "relation",
						"required": true,
						"presentable": false,
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
						"presentable": false,
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
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "h0okjh8g",
						"name": "covers",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
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
							"maxSelect": 99,
							"maxSize": 20971520,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "wgzhppl8",
						"name": "metadata",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2000000
						}
					},
					{
						"system": false,
						"id": "vr9ftnmg",
						"name": "old_id",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_jj5RsfT` + "`" + ` ON ` + "`" + `publications` + "`" + ` (` + "`" + `release` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "mu2u4hp0vc4dle5",
				"created": "2023-07-01 03:56:44.672Z",
				"updated": "2024-07-15 17:00:44.290Z",
				"name": "books",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wx5t7htt",
						"name": "publication",
						"type": "relation",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"collectionId": "guv9vnyfu5pdz9t",
							"cascadeDelete": true,
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
						"presentable": true,
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
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "n99n0fa3",
						"name": "covers",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
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
							"maxSelect": 99,
							"maxSize": 20971520,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "6m7pzsej",
						"name": "price",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "inz6maav",
						"name": "note",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "nudhir82",
						"name": "metadata",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2000000
						}
					},
					{
						"system": false,
						"id": "ifejwbve",
						"name": "old_id",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_gZH4WB5` + "`" + ` ON ` + "`" + `books` + "`" + ` (` + "`" + `publication` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "2lrfiedkzjul4s1",
				"created": "2023-07-01 07:10:59.758Z",
				"updated": "2024-07-15 17:00:44.290Z",
				"name": "publishers",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "e5b8x7mo",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
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
						"presentable": false,
						"unique": false,
						"options": {
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
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "atfsttrk",
						"name": "slug",
						"type": "text",
						"required": true,
						"presentable": false,
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
						"presentable": false,
						"unique": false,
						"options": {
							"min": 4,
							"max": 7,
							"pattern": "^#(?:[0-9a-fA-F]{3}){1,2}$"
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_RmOvURr` + "`" + ` ON ` + "`" + `publishers` + "`" + ` (` + "`" + `slug` + "`" + `)"
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
				"created": "2023-07-01 07:23:12.193Z",
				"updated": "2024-07-15 17:00:44.291Z",
				"name": "releases",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "4tlee6c6",
						"name": "title",
						"type": "relation",
						"required": true,
						"presentable": true,
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
						"id": "lzdfyn1k",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "8nglstcz",
						"name": "type",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "gcb2iw3u",
						"name": "digital",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "awceclka",
						"name": "disambiguation",
						"type": "text",
						"required": false,
						"presentable": false,
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
						"presentable": false,
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
						"id": "w9p6emac",
						"name": "partner",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "2lrfiedkzjul4s1",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "1t7lpcuz",
						"name": "status",
						"type": "select",
						"required": false,
						"presentable": false,
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
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
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
				"id": "kdwverajgytgjpe",
				"created": "2023-08-01 16:51:59.175Z",
				"updated": "2024-07-15 17:00:44.291Z",
				"name": "staffs",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "aewdsjta",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
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
				"id": "0l473ttmx8o31i9",
				"created": "2023-08-11 03:05:01.810Z",
				"updated": "2024-07-15 17:00:44.321Z",
				"name": "bookDetails",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "5bioi2df",
						"name": "publishDate",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "lno01kvo",
						"name": "edition",
						"type": "text",
						"required": false,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "rizq0u8o",
						"name": "price",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "qd8njjtw",
						"name": "note",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "jbxzwsmo",
						"name": "metadata",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "tvhvxilz",
						"name": "publication",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "guv9vnyfu5pdz9t",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "ykllyrgb",
						"name": "release",
						"type": "relation",
						"required": true,
						"presentable": false,
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
						"id": "frx5oo1q",
						"name": "parentCollection",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "4lwu2y9b",
						"name": "parentId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "mu2u4hp0vc4dle5",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "0ikseh5n",
						"name": "covers",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
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
							"maxSelect": 99,
							"maxSize": 20971520,
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
					"query": "SELECT \n  books.id,\n  books.publishDate,\n  books.edition,\n  books.price,\n  books.note,\n  \"{}\" as metadata,\n  publications.id as publication,\n  publications.release,\n  \"publications\" as parentCollection,\n  publications.id as parentId,\n  publications.covers,\n  publications.created,\n  publications.updated\nFROM books\nLEFT JOIN publications ON books.publication = publications.id\nWHERE books.covers = \"[]\"\nUNION ALL\nSELECT\n  books.id,\n  books.publishDate,\n  books.edition,\n  books.price,\n  books.note,\n  \"{}\" as metadata,\n  publications.id as publication,\n  publications.release,\n  \"books\" as parentCollection,\n  books.id as parentId,\n  books.covers,\n  books.created,\n  books.updated\nFROM books\nLEFT JOIN publications ON books.publication = publications.id\nWHERE books.covers != \"[]\""
				}
			},
			{
				"id": "6uk141b1jx0dkhu",
				"created": "2023-08-16 16:47:11.110Z",
				"updated": "2024-07-15 17:00:44.291Z",
				"name": "works",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ctbkfe27",
						"name": "title",
						"type": "relation",
						"required": true,
						"presentable": false,
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
						"id": "zxp2x5pq",
						"name": "staff",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "kdwverajgytgjpe",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "c5ta1rqb",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pj5teplb",
						"name": "priority",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_RXT00L7` + "`" + ` ON ` + "`" + `works` + "`" + ` (` + "`" + `title` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "wtd1x8mugo9liuw",
				"created": "2023-08-21 19:11:48.813Z",
				"updated": "2024-07-15 17:00:44.291Z",
				"name": "reviews",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "fwtyy7lx",
						"name": "release",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "3j32s2l7fdos1e4",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"title",
								"name"
							]
						}
					},
					{
						"system": false,
						"id": "xboghilk",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "wup1bczp",
						"name": "header",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "j68gxbbk",
						"name": "content",
						"type": "editor",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "xcwy9z37",
						"name": "score",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 0,
							"max": 10,
							"noDecimal": false
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_ojmVl7c` + "`" + ` ON ` + "`" + `reviews` + "`" + ` (\n  ` + "`" + `release` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)",
					"CREATE INDEX ` + "`" + `idx_KazuTq0` + "`" + ` ON ` + "`" + `reviews` + "`" + ` (` + "`" + `release` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "user = @request.auth.id && @request.auth.verified = true",
				"updateRule": "user = @request.auth.id",
				"deleteRule": "user = @request.auth.id",
				"options": {}
			},
			{
				"id": "ldgikhnt12bt4a6",
				"created": "2023-09-10 18:17:23.763Z",
				"updated": "2024-07-15 17:00:44.291Z",
				"name": "collections",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "nfiyam1n",
						"name": "owner",
						"type": "relation",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"username"
							]
						}
					},
					{
						"system": false,
						"id": "vq5nrwbv",
						"name": "visibility",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"PRIVATE",
								"UNLISTED",
								"PUBLIC"
							]
						}
					},
					{
						"system": false,
						"id": "9ctkkqxa",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": 1,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mbjirlsm",
						"name": "default",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "mvj1yde4",
						"name": "description",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "c66s8lzq",
						"name": "order",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 0,
							"max": null,
							"noDecimal": false
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_h3R3CeQ` + "`" + ` ON ` + "`" + `collections` + "`" + ` (` + "`" + `owner` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "380dcd8bylxiw0w",
				"created": "2023-09-10 18:22:02.787Z",
				"updated": "2024-07-15 17:00:44.292Z",
				"name": "collectionMembers",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "9louzrtd",
						"name": "collection",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "ldgikhnt12bt4a6",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"id",
								"owner",
								"name"
							]
						}
					},
					{
						"system": false,
						"id": "c8hj1jzf",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"username"
							]
						}
					},
					{
						"system": false,
						"id": "cbsgvhrm",
						"name": "role",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"EDITOR",
								"MEMBER"
							]
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_RCz3SdE` + "`" + ` ON ` + "`" + `collectionMembers` + "`" + ` (` + "`" + `collection` + "`" + `)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_7sz2GU1` + "`" + ` ON ` + "`" + `collectionMembers` + "`" + ` (\n  ` + "`" + `collection` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)"
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "nkw6spljgatiiyx",
				"created": "2023-09-10 18:29:22.941Z",
				"updated": "2024-07-15 17:00:44.292Z",
				"name": "collectionBooks",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "mptqkias",
						"name": "collection",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "ldgikhnt12bt4a6",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"id",
								"owner",
								"name"
							]
						}
					},
					{
						"system": false,
						"id": "6wxoqtgr",
						"name": "book",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "mu2u4hp0vc4dle5",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"id",
								"publication",
								"edition"
							]
						}
					},
					{
						"system": false,
						"id": "3k6cwj6s",
						"name": "quantity",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "8huiodky",
						"name": "status",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"PLANNING",
								"COMPLETED"
							]
						}
					},
					{
						"system": false,
						"id": "barh8omf",
						"name": "notes",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
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
				"id": "0jd7tc1qu0m84u8",
				"created": "2023-10-10 09:43:12.590Z",
				"updated": "2024-07-15 17:00:44.292Z",
				"name": "linkSources",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "smxo8l5p",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "e5pu0xos",
						"name": "color",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^#(?:[0-9a-fA-F]{3}){1,2}$"
						}
					},
					{
						"system": false,
						"id": "x105n2lt",
						"name": "icon",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/png",
								"image/jpeg",
								"image/svg+xml",
								"image/webp"
							],
							"thumbs": [
								"50x50"
							],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_JpwCQ4l` + "`" + ` ON ` + "`" + `linkSources` + "`" + ` (` + "`" + `name` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "aq9t2qxia36sz22",
				"created": "2023-10-10 09:43:34.829Z",
				"updated": "2024-07-15 17:00:44.292Z",
				"name": "links",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "vesku0uz",
						"name": "title",
						"type": "relation",
						"required": true,
						"presentable": false,
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
						"id": "qztytmal",
						"name": "source",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "0jd7tc1qu0m84u8",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "mkjeereu",
						"name": "url",
						"type": "url",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
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
				"id": "walac4l9hx6i63v",
				"created": "2023-10-15 02:39:47.497Z",
				"updated": "2024-07-15 17:00:44.292Z",
				"name": "genres",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "9mkj3bxh",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "yxa6wanz",
						"name": "slug",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_CtQKNT9` + "`" + ` ON ` + "`" + `genres` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "tnyytazu9gvdxse",
				"created": "2023-10-17 11:39:22.520Z",
				"updated": "2024-07-15 17:00:44.323Z",
				"name": "titleCovers",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "dtsgoewt",
						"name": "covers",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "hv6vysgb",
						"name": "title",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "pkz9pabs",
						"name": "parentCollection",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "nw2it7jg",
						"name": "volume",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "36yjjk6c",
						"name": "metadata",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
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
					"query": "SELECT \n  id,\n  covers, \n  title, \n  parentCollection,\n  volume,\n  \"{}\" as metadata\nFROM \n  (\n    SELECT \n      books.id as id, \n      books.covers as covers, \n      publications.volume as volume,\n      releases.title as title,\n      \"books\" as parentCollection \n    FROM \n      books \n      RIGHT JOIN publications ON publications.id = books.publication \n      RIGHT JOIN releases ON releases.id = publications.release \n    WHERE \n      books.covers != \"[]\" \n    UNION \n    SELECT \n      publications.id as id, \n      publications.covers as covers, \n      publications.volume as volume,\n      releases.title as title,\n      \"publications\" as parentCollection \n    FROM \n      publications \n      RIGHT JOIN releases ON releases.id = publications.release \n    WHERE \n      publications.covers != \"[]\"\n  )\nORDER BY title ASC, volume ASC, (CASE parentCollection\n  WHEN 'publications' THEN 0\n  ELSE 1\nEND) ASC;"
				}
			},
			{
				"id": "8msmj3ci8k33wbe",
				"created": "2023-10-28 06:03:39.444Z",
				"updated": "2024-07-15 17:00:44.292Z",
				"name": "demographics",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "pnwz4ww0",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "dfd7yiaa",
						"name": "slug",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_TRi2rPa` + "`" + ` ON ` + "`" + `demographics` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "uw1kal8bb55bf1e",
				"created": "2023-10-28 17:48:09.861Z",
				"updated": "2024-07-15 17:00:44.325Z",
				"name": "releaseDetails",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "vyj65qrm",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "2ranrkeu",
						"name": "title",
						"type": "relation",
						"required": true,
						"presentable": true,
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
						"id": "t41q2nq2",
						"name": "publisher",
						"type": "relation",
						"required": true,
						"presentable": false,
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
						"id": "eogcq8rv",
						"name": "status",
						"type": "select",
						"required": false,
						"presentable": false,
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
						"id": "bkm7he3h",
						"name": "cover",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [
								"80x120"
							],
							"maxSelect": 1,
							"maxSize": 20971520,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "s7ls8yq7",
						"name": "metadata",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2000000
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
					"query": "SELECT releases.id, releases.name, releases.title, releases.publisher, releases.status, titles.cover, titles.metadata, releases.created, releases.updated\nFROM releases, titles\nWHERE releases.title = titles.id"
				}
			},
			{
				"id": "wzjok6uyx3y1qiz",
				"created": "2024-07-18 14:29:52.500Z",
				"updated": "2024-07-18 14:29:52.500Z",
				"name": "additionalTitleNames",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "yozmq8fz",
						"name": "title",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "s91oidzeo1xm4m7",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "k7s4a6td",
						"name": "language",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "uedowwct",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_3Fo1G1g` + "`" + ` ON ` + "`" + `additionalTitleNames` + "`" + ` (` + "`" + `title` + "`" + `)"
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
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
