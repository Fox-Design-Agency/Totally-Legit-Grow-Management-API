// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"log"

	"totally-legit-grow-management/v1/internal/logic"
	devicehandlers "totally-legit-grow-management/v1/internal/server/deviceHandlers"
	growspothandlers "totally-legit-grow-management/v1/internal/server/growSpotHandlers"
	growinggrouphandlers "totally-legit-grow-management/v1/internal/server/growingGroupHandlers"
	growinglocationshandlers "totally-legit-grow-management/v1/internal/server/growingLocationHandlers"
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
	"totally-legit-grow-management/v1/resources/config"
)

func NewServer(cfg config.Config, shouldMigrate bool) *Server {
	control, err := logic.NewLogicControl(cfg)
	if err != nil {
		log.Fatal("Logic Layer Not Setup!")
	}

	if shouldMigrate {
		err := control.Persistence.MigrateDBUP()
		if err != nil {
			// migrations couldn't happen
			log.Println(err)
		}
	}

	// Device Handlers
	deviceHandler := devicehandlers.InitDeviceHandler(control.DeviceLogic)
	// Growing Group Handlers
	growingGroupHandler := growinggrouphandlers.InitGrowingGroupHandler(control.GrowingGroupLogic)
	// GrowingLocation Handlers
	growingLocationHandler := growinglocationshandlers.InitGrowingLocationHandler(control.GrowingLocationLogic)
	// Growing Medium Handlers
	growingMediumHandler := growingmediumhandlers.InitGrowingMediumHandler(control.GrowingMediumLogic)
	// Grow Spot Handlers
	growSpotHandler := growspothandlers.InitGrowSpotHandler(control.GrowSpotLogic)
	// Harvest Handlers
	harvestHandler := harvesthandlers.InitHarvestHandler(control.HarvestLogic)
	// Inventory Handlers
	inventoryHandler := inventoryhandlers.InitInventoryHandler(control.InventoryLogic)
	// Member Handlers
	memberHandler := memberhandlers.InitMemberHandler(control.MembersLogic)
	// Nutrient Handlers
	nutrienthandlers := nutrienthandlers.InitNutrientHandler(control.NutrientLogic)
	// Organization Handlers
	organizationHandler := organizationhandlers.InitOrganizationHandler(control.OrganizationLogic)
	// Plant Category Handlers
	plantCategoryHandler := plantcategoryhandlers.InitPlantCategoryHandler(control.PlantCategoriesLogic)
	// Plant Handlers
	plantHandler := planthandlers.InitPlantHandler(control.PlantLogic)
	// Plant Life Cycle Handlers
	plantLifecycleHandler := plantlifecyclehandlers.InitPlantLifecycleHandler(control.PlantLifecycleLogic)
	// Plant Stage Handlers
	plantStageHandler := plantstagehandlers.InitPlantStageHandler(control.PlantStagesLogic)
	// Product Handlers
	productHandler := producthandlers.InitProductHandler(control.ProductLogic)
	// Seed Handlers
	seedHandler := seedhandlers.InitSeedHandler(control.SeedLogic)
	// Task Handlers
	taskHandler := taskhandlers.InitTaskHandler(control.TaskLogic)

	return &Server{
		DeviceHandlers:  deviceHandler,
		GrowingGroup:    growingGroupHandler,
		GrowingLocation: growingLocationHandler,
		GrowingMedium:   growingMediumHandler,
		GrowSpot:        growSpotHandler,
		Harvest:         harvestHandler,
		Inventory:       inventoryHandler,
		Member:          memberHandler,
		Nutrient:        nutrienthandlers,
		Organization:    organizationHandler,
		PlantCategory:   plantCategoryHandler,
		Plant:           plantHandler,
		PlantLifeCycle:  plantLifecycleHandler,
		PlantStage:      plantStageHandler,
		Product:         productHandler,
		Seed:            seedHandler,
		Task:            taskHandler,
	}
}
