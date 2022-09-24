// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logic

import (
	"log"
	devicecontrol "totally-legit-grow-management/v1/internal/logic/deviceControl"
	growspotplantscontrol "totally-legit-grow-management/v1/internal/logic/growSpotPlantsControl"
	growspotcontrol "totally-legit-grow-management/v1/internal/logic/growSpotsControl"
	growinggroupcontrol "totally-legit-grow-management/v1/internal/logic/growingGroupsControl"
	growinglevelcontrol "totally-legit-grow-management/v1/internal/logic/growingLevelsControl"
	growinglocationcontrol "totally-legit-grow-management/v1/internal/logic/growingLocationsControl"
	growingmediumcontrol "totally-legit-grow-management/v1/internal/logic/growingMediumControl"
	harvestcontrol "totally-legit-grow-management/v1/internal/logic/harvestsControl"
	inventorycontrol "totally-legit-grow-management/v1/internal/logic/inventoryControl"
	membercontrol "totally-legit-grow-management/v1/internal/logic/membersControl"
	nutrientcontrol "totally-legit-grow-management/v1/internal/logic/nutrientControl"
	organizationcontrol "totally-legit-grow-management/v1/internal/logic/organizationsControl"
	plantcategoriescontrol "totally-legit-grow-management/v1/internal/logic/plantCategoriesControl"
	plantcontrol "totally-legit-grow-management/v1/internal/logic/plantControl"
	plantlifecyclecontrol "totally-legit-grow-management/v1/internal/logic/plantLifecyclesControl"
	plantstagescontrol "totally-legit-grow-management/v1/internal/logic/plantStagesControl"
	productcontrol "totally-legit-grow-management/v1/internal/logic/productsControl"
	seedcontrol "totally-legit-grow-management/v1/internal/logic/seedsControl"
	taskcontrol "totally-legit-grow-management/v1/internal/logic/taskControl"
	"totally-legit-grow-management/v1/internal/persistence"
	"totally-legit-grow-management/v1/resources/config"
)

func NewLogicControl(cfg config.Config) (*Logic, error) {
	persist, err := persistence.NewPersistence(cfg.Database.Dialect(), cfg.Database.Connection())
	if err != nil {
		log.Fatal("Couldn't Make DB Connection")
	}
	// device logic
	deviceLogic := devicecontrol.InitDeviceLogic(persist.Devices)

	// growing groups
	growingGroupLogic := growinggroupcontrol.InitGrowingGroupLogic(persist.GrowingGroup)

	//  growing levels
	growingLevelsLogic := growinglevelcontrol.InitGrowingLevelLogic(persist.GrowingLevel)

	// growing locations
	growingLocationsLogic := growinglocationcontrol.InitGrowingLocationLogic(persist.GrowingLocations)

	// growing medium
	growingMediumLogic := growingmediumcontrol.InitGrowingMediumLogic(persist.GrowingMedium)

	// grow spot plants
	growSpotPlantsLogic := growspotplantscontrol.InitGrowSpotPlantLogic(persist.GrowSpotPlants)

	// grow spots
	growSpotLogic := growspotcontrol.InitGrowSpotLogic(persist.GrowSpots)

	// harvests
	harvestsLogic := harvestcontrol.InitHarvestLogic(persist.Harvests)

	// inventory
	inventoryLogic := inventorycontrol.InitInventoryLogic(persist.Inventories)

	// members
	membersLogic := membercontrol.InitMembersLogic(persist.Members)

	// nutrients
	nutrientLogic := nutrientcontrol.InitNutrientLogic(persist.Nutrients)

	// organizations
	organizationLogic := organizationcontrol.InitOrganizationLogic(persist.Organizations)

	// plant categories
	plantCategoryLogic := plantcategoriescontrol.InitPlantCategoriesLogic(persist.PlantCategories)

	// plant
	plantLogic := plantcontrol.InitPlantLogic(persist.Plants)

	// plant life cycle
	plantLifecycleLogic := plantlifecyclecontrol.InitPlantLifecycleLogic(persist.PlantLifecycles)

	// plant stages
	plantStagesLogic := plantstagescontrol.InitPlantStageLogic(persist.Plantstages)

	// products
	productLogic := productcontrol.InitProductLogic(persist.Products)

	// seeds
	seedsLogic := seedcontrol.InitSeedLogic(persist.Seeds)

	//task
	taskLogic := taskcontrol.InitTaskLogic(persist.Tasks)

	return &Logic{
		DeviceLogic:          deviceLogic,
		GrowingGroupLogic:    growingGroupLogic,
		GrowingLevelLogic:    growingLevelsLogic,
		GrowingLocationLogic: growingLocationsLogic,
		GrowingMediumLogic:   growingMediumLogic,
		GrowSpotPlantsLogic:  growSpotPlantsLogic,
		GrowSpotLogic:        growSpotLogic,
		HarvestLogic:         harvestsLogic,
		InventoryLogic:       inventoryLogic,
		MembersLogic:         membersLogic,
		NutrientLogic:        nutrientLogic,
		OrganizationLogic:    organizationLogic, Persistence: persist.DBControl,
		PlantCategoriesLogic: plantCategoryLogic,
		PlantLogic:           plantLogic,
		PlantLifecycleLogic:  plantLifecycleLogic,
		PlantStagesLogic:     plantStagesLogic,
		ProductLogic:         productLogic,
		SeedLogic:            seedsLogic,
		TaskLogic:            taskLogic,
	}, nil
}
