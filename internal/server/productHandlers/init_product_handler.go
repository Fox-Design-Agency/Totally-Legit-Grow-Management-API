// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package producthandlers

import (
	productcontrol "totally-legit-grow-management/v1/internal/logic/productsControl"
)

func InitProductHandler(productLogic productcontrol.IProductsLogic) *Product {
	return &Product{
		Logic: productLogic,
	}
}
