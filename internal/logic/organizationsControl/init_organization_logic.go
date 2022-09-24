// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package organizationcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence/organizations"
)

func InitOrganizationLogic(organizationDB organizations.IOrganizationsDB) *OrganizationControl {
	return &OrganizationControl{
		Persistence: organizationDB,
	}
}
