// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	devicehandlers "totally-legit-grow-management/v1/internal/server/deviceHandlers"
	growspothandlers "totally-legit-grow-management/v1/internal/server/growSpotHandlers"
	growspotplantshandlers "totally-legit-grow-management/v1/internal/server/growSpotPlantsHandlers"
	growinggrouphandlers "totally-legit-grow-management/v1/internal/server/growingGroupHandlers"
	growinglevelhandlers "totally-legit-grow-management/v1/internal/server/growingLevelHandlers"
	growinglocationhandlers "totally-legit-grow-management/v1/internal/server/growingLocationHandlers"
	growingmediumhandlers "totally-legit-grow-management/v1/internal/server/growingMediumHandlers"
	harvesthandlers "totally-legit-grow-management/v1/internal/server/harvestHandlers"
	inventoryhandlers "totally-legit-grow-management/v1/internal/server/inventoryHandlers"
	memberhandlers "totally-legit-grow-management/v1/internal/server/memberHandlers"
	nutrienthandlers "totally-legit-grow-management/v1/internal/server/nutrientHandlers"
	organizationhandlers "totally-legit-grow-management/v1/internal/server/organizationHandlers"
	plantcategoryhandlers "totally-legit-grow-management/v1/internal/server/plantCategoryHandlers"
	planthandlers "totally-legit-grow-management/v1/internal/server/plantHandlers"
	plantlifecyclehandlers "totally-legit-grow-management/v1/internal/server/plantLifeCycleHandlers"
	plantstagehandlers "totally-legit-grow-management/v1/internal/server/plantStageHandlers"
	producthandlers "totally-legit-grow-management/v1/internal/server/productHandlers"
	seedhandlers "totally-legit-grow-management/v1/internal/server/seedHandlers"
	taskhandlers "totally-legit-grow-management/v1/internal/server/taskHandlers"
)

type Server struct {
	DeviceHandlers  devicehandlers.IDeviceHandler
	GrowingGroup    growinggrouphandlers.IGrowingGroupHandler
	GrowingLevel    growinglevelhandlers.IGrowingLevelHandler
	GrowingLocation growinglocationhandlers.IGrowingLocationHandlers
	GrowingMedium   growingmediumhandlers.IGrowingMediumHandler
	GrowSpot        growspothandlers.IGrowSpotHandler
	GrowSpotPlants  growspotplantshandlers.IGrowSpotPlantsHandlers
	Harvest         harvesthandlers.IHarvestHandler
	Inventory       inventoryhandlers.IInventoryHandler
	Member          memberhandlers.IMemberHandler
	Nutrient        nutrienthandlers.INutrientHandler
	Organization    organizationhandlers.IOrganizationHandler
	PlantCategory   plantcategoryhandlers.IPlantCategoryHandler
	Plant           planthandlers.IPlantHandlers
	PlantLifeCycle  plantlifecyclehandlers.IPlantLifeCycleHandler
	PlantStage      plantstagehandlers.IPlantStageHandler
	Product         producthandlers.IProductHandler
	Seed            seedhandlers.ISeedHandler
	Task            taskhandlers.ITaskHandler
}
