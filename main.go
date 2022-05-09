// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"totally-legit-grow-management/v1/pkg/server"
	"totally-legit-grow-management/v1/resources/config"

	"github.com/gorilla/mux"
)

func main() {
	/**********************************************************************
	/
	/	Configuration
	/
	/**********************************************************************/
	ctx, cancel := context.WithCancel(context.Background())

	cfg := config.LoadConfig()

	svr := server.NewServer(cfg, true)

	/**********************************************************************
	/
	/	Initialize router and middlesware
	/
	/**********************************************************************/

	r := mux.NewRouter()

	/**********************************************************************
	/
	/	Health Check & Container routes
	/
	/**********************************************************************/

	r.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) { rw.WriteHeader(http.StatusOK) })

	/**********************************************************************
	/
	/	Device Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/device", svr.CreateDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/action", svr.CreateDeviceAction).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/growing-group", svr.CreateGrowingGroupDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/growing-location", svr.CreateGrowingLocationDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/growing-level", svr.CreateGrowingLevelDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/growing-spot", svr.CreateGrowingSpotDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device", svr.DeleteDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device/growing-group", svr.DeleteGrowingGroupDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device/growing-location", svr.DeleteGrowingLocationDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device/growing-level", svr.DeleteGrowingLevelDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device/growing-spot", svr.DeleteGrowingSpotDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device", svr.EditDevice).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/device", svr.GetDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/actions", svr.GetDeviceActions).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices", svr.GetAllDevices).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/growing-group", svr.GetGrowingGroupDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/growing-location", svr.GetGrowingLocationDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/growing-level", svr.GetGrowingLevelDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/growing-spot", svr.GetGrowingSpotDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices/growing-group", svr.GetAllGrowingGroupDevices).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices/growing-location", svr.GetAllGrowingLocationDevices).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices/growing-level", svr.GetAllGrowingLevelDevices).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices/growing-spot", svr.GetAllGrowingSpotDevices).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Grow Spot Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/growing-spot", svr.CreateGrowSpot).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-spot", svr.DeleteGrowSpot).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-spot", svr.EditGrowSpot).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-spot", svr.GetGrowSpot).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-spots", svr.GetAllGrowSpots).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Grow Spot Plants Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/growing-spot-plant", svr.CreateGrowSpotPlant).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-spot-plant", svr.DeleteGrowSpotPlant).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-spot-plant", svr.EditGrowSpotPlant).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-spot-plant", svr.GetGrowSpotPlant).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-spot-plants", svr.GetAllGrowSpotPlants).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Growing Group Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/growing-group", svr.CreateGrowingGroup).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-group", svr.DeleteGrowingGroup).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-group", svr.EditGrowingGroup).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-group", svr.GetGrowingGroup).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-groups", svr.GetAllGrowingGroups).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Growing Level Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/growing-level", svr.CreateGrowingLevel).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-level", svr.DeleteGrowingLevel).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-level", svr.EditGrowingLevel).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-level", svr.GetGrowingLevel).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-levels", svr.GetAllGrowingLevels).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Grow Locations Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/growing-location", svr.CreateGrowingLocation).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-location", svr.DeleteGrowingLocation).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-location", svr.EditGrowingLocation).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-location", svr.GetGrowingLocation).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-locations", svr.GetAllGrowingLocations).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Growing Medium Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/growing-medium", svr.CreateGrowingMedium).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-medium", svr.DeleteGrowingMedium).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-medium", svr.EditGrowingMedium).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-medium", svr.GetGrowingMedium).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-mediums", svr.GetAllGrowingMediums).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Harvest Routes
	/
	/**********************************************************************/

	// r.HandleFunc("/api/v1/harvest", svr.CreateHarvest).Methods(http.MethodPost)
	// r.HandleFunc("/api/v1/harvest", svr.DeleteHarvest).Methods(http.MethodDelete)
	// r.HandleFunc("/api/v1/harvest", svr.EditHarvest).Methods(http.MethodPut)
	// r.HandleFunc("/api/v1/harvest", svr.GetHarvest).Methods(http.MethodGet)
	// r.HandleFunc("/api/v1/harvests", svr.GetAllHarvests).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Inventory Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/inventory", svr.CreateInventory).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/inventory", svr.DeleteInventory).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/inventory", svr.EditInventory).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/inventory", svr.GetInventory).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/inventories", svr.GetAllInventories).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Members Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/member", svr.CreateMember).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/member", svr.DeleteMember).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/member", svr.EditMember).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/member", svr.GetMember).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/members", svr.GetAllMembers).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Nutrients Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/nutrient", svr.CreateNutrient).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/nutrient", svr.DeleteNutrient).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/nutrient", svr.EditNutrient).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/nutrient", svr.GetNutrient).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/nutrients", svr.GetAllNutrients).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Organization Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/organization", svr.CreateOrganization).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/organization", svr.DeleteOrganization).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/organization", svr.EditOrganization).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/organization", svr.GetOrganization).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/organizations", svr.GetAllOrganizations).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Plant Categories Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/plant-category", svr.CreatePlantCategory).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-category", svr.DeletePlantCategory).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/plant-category", svr.EditPlantCategory).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/plant-category", svr.GetPlantCategory).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-categories", svr.GetAllPlantCategories).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Plant Life Cycles Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/plant-life-cycle", svr.CreatePlantLifeCycle).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-life-cycle", svr.DeletePlantLifeCycle).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/plant-life-cycle", svr.EditPlantLifeCycle).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/plant-life-cycle", svr.GetPlantLifeCycleByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-life-cycles", svr.GetAllPlantLifeCycles).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Plant Stages Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/plant-stage", svr.CreatePlantStage).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/care", svr.CreatePlantStageCare).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/nutrient", svr.CreatePlantStageNutrients).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/connect", svr.ConnectPlantStage).Methods(http.MethodPost)
	//
	r.HandleFunc("/api/v1/plant-stage", svr.DeletePlantStage).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/plant-stage", svr.EditPlantStage).Methods(http.MethodPut)
	//
	r.HandleFunc("/api/v1/plant-stage", svr.GetPlantStageByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stage/care", svr.GetPlantStageCareByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stage/nutrient", svr.GetPlantStageNutrient).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stages", svr.GetAllPlantStages).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Plant Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/plant", svr.CreatePlant).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant", svr.DeletePlant).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/plant", svr.EditPlant).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/plant", svr.GetPlant).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plants", svr.GetAllPlants).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Products Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/product", svr.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/product", svr.DeleteProduct).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/product", svr.EditProduct).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/product", svr.GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/products", svr.GetAllProducts).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Seeds Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/seed", svr.CreateSeed).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/seed", svr.DeleteSeed).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/seed", svr.EditSeed).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/seed", svr.GetSeed).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/seeds", svr.GetAllSeeds).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Task Routes
	/
	/**********************************************************************/

	r.HandleFunc("/api/v1/task", svr.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/task", svr.DeleteTask).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/task", svr.EditTask).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/task", svr.GetTask).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/tasks", svr.GetAllTasks).Methods(http.MethodGet)

	/**********************************************************************
	/
	/	Start Server
	/
	/**********************************************************************/
	httpServer := &http.Server{
		Addr:        ":" + cfg.Port,
		Handler:     r,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	log.Println("Running local on port: ", cfg.Port)
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	/**********************************************************************
	/
	/	Run Shutdown Sequence
	/
	/**********************************************************************/
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}
