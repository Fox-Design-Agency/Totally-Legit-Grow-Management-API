package main

import (
	"fmt"
	"log"
	"net/http"

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

	cfg := config.LoadConfig()

	// define server
	svr := server.NewServer(cfg)

	err := svr.Persistence.MigrateDBUP()
	if err != nil {
		// migrations couldn't happen
		log.Println(err)
	}

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

	/**********************************************************************
	/
	/	Growing Group Routes
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Growing Level Routes
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Grow Locations Routes
	/
	/**********************************************************************/

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

	/**********************************************************************
	/
	/	Plant Stages Routes
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Plant Routes
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Start Server
	/
	/**********************************************************************/
	log.Println("Running local on port: ", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil))
}
