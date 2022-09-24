// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package products

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IProductsDB = &ProductPersistence{}

//
type IProductsDB interface {
	CreateProduct(*routemodels.CreateProductRequest) (*routemodels.CreateProductResponse, error)
	DeleteProduct(*routemodels.DeleteProductRequest) error
	EditProduct(*routemodels.EditProductRequest) (*routemodels.EditProductResponse, error)
	GetProduct(*routemodels.GetProductRequest) (*routemodels.GetProductResponse, error)
	GetAllProducts(*routemodels.GetAllProductsRequest) (*routemodels.GetAllProductsResponse, error)
}

type ProductPersistence struct {
	Postgres *sqlx.DB
}
