// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

import (
	devicepersistence "totally-legit-grow-management/v1/internal/persistence/devices"
	growspotplants "totally-legit-grow-management/v1/internal/persistence/growSpotPlants"
	growspots "totally-legit-grow-management/v1/internal/persistence/growSpots"
	growinggroups "totally-legit-grow-management/v1/internal/persistence/growingGroups"
	growinglevels "totally-legit-grow-management/v1/internal/persistence/growingLevels"
	growinglocations "totally-legit-grow-management/v1/internal/persistence/growingLocations"
	growingmedium "totally-legit-grow-management/v1/internal/persistence/growingMedium"
	"totally-legit-grow-management/v1/internal/persistence/harvests"
	"totally-legit-grow-management/v1/internal/persistence/inventories"
	"totally-legit-grow-management/v1/internal/persistence/members"
	"totally-legit-grow-management/v1/internal/persistence/nutrients"
	"totally-legit-grow-management/v1/internal/persistence/organizations"
	plantcategories "totally-legit-grow-management/v1/internal/persistence/plantCategories"
	plantlifecycle "totally-legit-grow-management/v1/internal/persistence/plantLifeCycles"
	plantstages "totally-legit-grow-management/v1/internal/persistence/plantStages"
	"totally-legit-grow-management/v1/internal/persistence/plants"
	"totally-legit-grow-management/v1/internal/persistence/products"
	"totally-legit-grow-management/v1/internal/persistence/seeds"
	"totally-legit-grow-management/v1/internal/persistence/tasks"

	"github.com/jmoiron/sqlx"
)

var _ IPersistence = &DBControl{}

type DBControl struct {
	Postgres *sqlx.DB
}

//
type Persistence struct {
	DBControl        IPersistence
	Devices          devicepersistence.IDevicesDB
	GrowingGroup     growinggroups.IGrowingGroupsDB
	GrowingLevel     growinglevels.IGrowingLevelsDB
	GrowingLocations growinglocations.IGrowingLocationsDB
	GrowingMedium    growingmedium.IGrowingMediumsDB
	GrowSpotPlants   growspotplants.IGrowSpotPlantsDB
	GrowSpots        growspots.IGrowSpotsDB
	Harvests         harvests.IHarvestsDB
	Inventories      inventories.IInventoriesDB
	Members          members.IMembersDB
	Nutrients        nutrients.INutrientsDB
	Organizations    organizations.IOrganizationsDB
	PlantCategories  plantcategories.IPlantCategoriesDB
	PlantLifecycles  plantlifecycle.IPlantLifeCyclesDB
	Plants           plants.IPlantsDB
	Plantstages      plantstages.IPlantStagesDB
	Products         products.IProductsDB
	Seeds            seeds.ISeedsDB
	Tasks            tasks.ITasksDB
}

//
type IPersistence interface {
	MigrateDBUP() error
	MigrateDBDown() error
	StartTransaction() (*sqlx.Tx, error)
}
