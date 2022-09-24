// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logic

import (
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
	plantcategorycontrol "totally-legit-grow-management/v1/internal/logic/plantCategoriesControl"
	plantcontrol "totally-legit-grow-management/v1/internal/logic/plantControl"
	plantlifecyclecontrol "totally-legit-grow-management/v1/internal/logic/plantLifecyclesControl"
	plantstagecontrol "totally-legit-grow-management/v1/internal/logic/plantStagesControl"
	productcontrol "totally-legit-grow-management/v1/internal/logic/productsControl"
	seedcontrol "totally-legit-grow-management/v1/internal/logic/seedsControl"
	taskcontrol "totally-legit-grow-management/v1/internal/logic/taskControl"
	"totally-legit-grow-management/v1/internal/persistence"
)

type Logic struct {
	DeviceLogic          devicecontrol.IDeviceLogic
	Persistence          persistence.IPersistence
	GrowingGroupLogic    growinggroupcontrol.IGrowingGroupsLogic
	GrowingLevelLogic    growinglevelcontrol.IGrowingLevelsLogic
	GrowingLocationLogic growinglocationcontrol.IGrowingLocationsLogic
	GrowingMediumLogic   growingmediumcontrol.IGrowingMediumsLogic
	GrowSpotPlantsLogic  growspotplantscontrol.IGrowSpotPlantsLogic
	GrowSpotLogic        growspotcontrol.IGrowSpotsLogic
	HarvestLogic         harvestcontrol.IHarvestsLogic
	InventoryLogic       inventorycontrol.IInventoriesLogic
	MembersLogic         membercontrol.IMembersLogic
	NutrientLogic        nutrientcontrol.INutrientsLogic
	OrganizationLogic    organizationcontrol.IOrganizationsLogic
	PlantCategoriesLogic plantcategorycontrol.IPlantCategoriesLogic
	PlantLogic           plantcontrol.IPlantsLogic
	PlantLifecycleLogic  plantlifecyclecontrol.IPlantLifeCyclesLogic
	PlantStagesLogic     plantstagecontrol.IPLantStagesLogic
	ProductLogic         productcontrol.IProductsLogic
	SeedLogic            seedcontrol.ISeedsLogic
	TaskLogic            taskcontrol.ITasksLogic
}
