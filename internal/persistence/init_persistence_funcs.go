// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

import (
	"log"
	"totally-legit-grow-management/v1/internal/persistence/devices"
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
	_ "github.com/lib/pq" // should be here
)

func NewPersistence(dialect, connectionInfo string) (*Persistence, error) {
	db, err := sqlx.Open(dialect, connectionInfo)
	if err != nil {
		log.Println("Error connecting to DB")
		log.Println(err)
		return nil, err
	}

	//devices
	devicePersistence := devices.InitDevicePersistence(db)
	//growing groups
	growingGroupPersistence := growinggroups.InitGrowingGroupPersistence(db)
	//growing levels
	growingLevelPersistence := growinglevels.InitGrowingLevelPersistence(db)
	//growing locations
	growingLocationPersistence := growinglocations.InitGrowingLocationPersistence(db)
	// growing medium
	growingMediumPersistence := growingmedium.InitGrowingMediumPersistence(db)
	// grow spot plants
	growSpotPlantsPersistence := growspotplants.InitGrowSpotPlantsPersistence(db)
	// grow spots
	growSpotsPersistence := growspots.InitDevicePersistence(db)
	// harvests
	harvestsPersistence := harvests.InitHarvestPersistence(db)
	// inventories
	inventoriesPersistence := inventories.InitInventoryPersistence(db)
	// members
	membersPersistence := members.InitMemberPersistence(db)
	// nutrients
	nutrientsPersistence := nutrients.InitNutrientPersistence(db)
	// organizations
	organizationsPersistence := organizations.InitOrganizationPersistence(db)
	// plant categories
	plantCategoriesPersistence := plantcategories.InitPlantCategoryPersistence(db)
	// plant life cycles
	plantLifecyclesPersistence := plantlifecycle.InitPlantLifecyclePersistence(db)
	// plants
	plantPersistence := plants.InitPlantPersistence(db)
	// plant stages
	plantstagePersistence := plantstages.InitPlantStagePersistence(db)
	// products
	productsPersistence := products.InitProductPersistence(db)
	// seeds
	seedsPersistence := seeds.InitSeedPersistence(db)
	// tasks
	tasksPersistence := tasks.InitTaskPersistence(db)

	return &Persistence{
		Devices:          devicePersistence,
		GrowingGroup:     growingGroupPersistence,
		GrowingLevel:     growingLevelPersistence,
		GrowingLocations: growingLocationPersistence,
		GrowingMedium:    growingMediumPersistence,
		GrowSpotPlants:   growSpotPlantsPersistence,
		GrowSpots:        growSpotsPersistence,
		Harvests:         harvestsPersistence,
		Inventories:      inventoriesPersistence,
		Members:          membersPersistence,
		Nutrients:        nutrientsPersistence,
		Organizations:    organizationsPersistence,
		PlantCategories:  plantCategoriesPersistence,
		PlantLifecycles:  plantLifecyclesPersistence,
		Plants:           plantPersistence,
		Plantstages:      plantstagePersistence,
		Products:         productsPersistence,
		Seeds:            seedsPersistence,
		Tasks:            tasksPersistence,
	}, nil
}
