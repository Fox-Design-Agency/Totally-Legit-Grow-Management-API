// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateProductRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateProductResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteProductRequest struct {
	ID string `json:"id"`
}

type EditProductRequest struct {
	DisplayName string `json:"displayName"`
}

type EditProductResponse struct {
	Product
}

type GetProductRequest struct {
	ID string `json:"id"`
}

type GetProductResponse struct {
	Product
}

type GetAllProductsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllProductsResponse struct {
	Products []Product
}
