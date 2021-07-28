package drinks

import (
	"github.com/ioartigiano/ioartigiano-be/internal/test"
	"net/http"
	"testing"
)

func TestInsertProductsAPI(t *testing.T) {
	router := test.MockRouter()
	header := test.MockHeader()
	db := test.DB(t)

	RegisterHandlers(router, NewService(NewRepository(db), test.MockValidator()))

	testsInsertProducts := []test.APITestCase{
		{"Insert product with only one PNG image",
			http.MethodPost,
			"/product",
			`{
					"sku": "test",
					"name": "entity with only one image PNG",
					"price": 22.5,
					"description": "test",
					"mainImg": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
					"productAsset": []
				}`,
			header,
			http.StatusCreated,
			"*entity with only one image PNG*",
			0,
		},
		{"Insert product with more PNG images",
			http.MethodPost,
			"/product",
			`{
					"sku": "test",
					"name": "entity with more images PNG",
					"price": 22.5,
					"description": "test",
					"mainImg": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII=",
					"productAssets": [
						{
							"path": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII="
						},
						{
							"path": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII="
						}
					]
				}`,
			header,
			http.StatusCreated,
			"*entity with more images PNG*",
			0,
		},
		{"Insert product with only one JPEG image",
			http.MethodPost,
			"/product",
			`{
					"sku": "test",
					"name": "entity with only one image JPEG",
					"price": 22.5,
					"description": "test",
					"mainImg": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEASABIAAD//gATQ3JlYXRlZCB3aXRoIEdJTVD/2wBDAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAABAAEDAREAAhEBAxEB/8QAFAABAAAAAAAAAAAAAAAAAAAAC//EABQQAQAAAAAAAAAAAAAAAAAAAAD/xAAUAQEAAAAAAAAAAAAAAAAAAAAA/8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAwDAQACEQMRAD8AP/B//9k="
				}`,
			header,
			http.StatusCreated,
			"*entity with only one image JPEG*",
			0,
		},
		{"Insert product with more JPEG images",
			http.MethodPost,
			"/product",
			`{
					"sku": "test",
					"name": "entity with more images JPEG",
					"price": 22.5,
					"description": "test",
					"mainImg": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEASABIAAD//gATQ3JlYXRlZCB3aXRoIEdJTVD/2wBDAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAABAAEDAREAAhEBAxEB/8QAFAABAAAAAAAAAAAAAAAAAAAAC//EABQQAQAAAAAAAAAAAAAAAAAAAAD/xAAUAQEAAAAAAAAAAAAAAAAAAAAA/8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAwDAQACEQMRAD8AP/B//9k=",
					"productAssets": [
						{
							"path": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAABAAEDAREAAhEBAxEB/8QAFAABAAAAAAAAAAAAAAAAAAAACv/EABQQAQAAAAAAAAAAAAAAAAAAAAD/xAAUAQEAAAAAAAAAAAAAAAAAAAAA/8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAwDAQACEQMRAD8AfwD/2Q=="
						},
						{
							"path": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAABAAEDAREAAhEBAxEB/8QAFAABAAAAAAAAAAAAAAAAAAAACv/EABQQAQAAAAAAAAAAAAAAAAAAAAD/xAAUAQEAAAAAAAAAAAAAAAAAAAAA/8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAwDAQACEQMRAD8AfwD/2Q=="
						}
					]
				}`,
			header,
			http.StatusCreated,
			"*entity with more images JPEG*",
			0,
		},
	}

	for _, tc := range testsInsertProducts {
		test.Endpoint(t, router, tc)
	}
	test.ResetTables(t, db, "drinks", "product_assets")
}

func TestGetProductsAPI(t *testing.T) {
	router := test.MockRouter()
	header := test.MockHeader()

	db := test.DB(t)

	query := `INSERT INTO "drinks" ("id", "sku", "name", "price", "description", "main_img") 
			VALUES ('967417e2-07e8-443b-b2a8-41efa912d3ed',	'test',	'test',	55.36,	'test',	'967417e2-07e8-443b-b2a8-41efa912d3ed.png')`
	test.PopulateTable(t, db, query)

	RegisterHandlers(router, NewService(NewRepository(db), test.MockValidator()))

	testsGetProducts := []test.APITestCase{
		{"Get product with id ",
			http.MethodGet,
			"/product/967417e2-07e8-443b-b2a8-41efa912d3ed",
			"",
			header,
			http.StatusOK,
			`{
					"ID": "967417e2-07e8-443b-b2a8-41efa912d3ed",
					"SKU": "test",
					"Name": "test",
					"Price": 55.36,
					"Description": "test",
					"MainImg": "967417e2-07e8-443b-b2a8-41efa912d3ed.png",
					"ProductAssets": null}`,
			0,
		},
		{"Get product with a invalid id ",
			http.MethodGet,
			"/product/item-not-exists",
			"",
			header,
			http.StatusBadRequest,
			"",
			0,
		},
		{"Get all product",
			http.MethodGet,
			"/product",
			"",
			header,
			http.StatusOK,
			`[{
					"ID": "967417e2-07e8-443b-b2a8-41efa912d3ed",
					"SKU": "test",
					"Name": "test",
					"Price": 55.36,
					"Description": "test",
					"MainImg": "967417e2-07e8-443b-b2a8-41efa912d3ed.png",
					"ProductAssets": []}]`,
			0,
		},
		{"Get a product that not exists",
			http.MethodGet,
			"/product/967417e2-07e8-443b-b2a8-41efa912d3ea",
			"",
			header,
			http.StatusNotFound,
			"",
			0,
		},
	}

	for _, tc := range testsGetProducts {
		test.Endpoint(t, router, tc)
	}
	test.ResetTables(t, db, "drinks", "product_assets")
}

func TestDeleteProductsAPI(t *testing.T) {
	router := test.MockRouter()
	header := test.MockHeader()
	db := test.DB(t)

	query := `INSERT INTO "drinks" ("id", "sku", "name", "price", "description", "main_img") 
			VALUES ('967417e2-07e8-443b-b2a8-41efa912d3ed',	'test',	'test',	55.36,	'test',	'967417e2-07e8-443b-b2a8-41efa912d3ed.png')`
	test.PopulateTable(t, db, query)

	RegisterHandlers(router, NewService(NewRepository(db), test.MockValidator()))

	testsDeleteProduct := []test.APITestCase{
		{"Delete a product",
			http.MethodDelete,
			"/product/967417e2-07e8-443b-b2a8-41efa912d3ed",
			"",
			header,
			http.StatusOK,
			"",
			0,
		},
		{"Delete a product that not exists",
			http.MethodDelete,
			"/product/967417e2-07e8-443b-b2a8-41efa912d3ee",
			"",
			header,
			http.StatusNotFound,
			"",
			0,
		},
	}

	for _, tc := range testsDeleteProduct {
		test.Endpoint(t, router, tc)
	}

	test.ResetTables(t, db, "drinks", "product_assets")
}

func TestUpdateProductsAPI(t *testing.T) {
	router := test.MockRouter()
	header := test.MockHeader()
	db := test.DB(t)

	query := `INSERT INTO "drinks" ("id", "sku", "name", "price", "description", "main_img")
			VALUES ('967417e2-07e8-443b-b2a8-41efa912d3ed',	'test',	'test',	55.36,	'test',	'967417e2-07e8-443b-b2a8-41efa912d3ed.png')`
	test.PopulateTable(t, db, query)
	RegisterHandlers(router, NewService(NewRepository(db), test.MockValidator()))

	testsUpdateProduct := []test.APITestCase{
		{"Update a product",
			http.MethodPut,
			"/product",
			`{
					"id": "967417e2-07e8-443b-b2a8-41efa912d3ed",
					"sku": "new_test",
					"name": "new_entity",
					"price": 252.5,
					"description": "new_test",
					"mainImg": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
					"productAssets": [
						{
							"path": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII="
						},
						{
							"path": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII="
						}
					]
				}`,
			header,
			http.StatusOK,
			"*new_test*",
			0,
		},
	}

	for _, tc := range testsUpdateProduct {
		test.Endpoint(t, router, tc)
	}

	test.ResetTables(t, db, "drinks", "product_assets")
}
