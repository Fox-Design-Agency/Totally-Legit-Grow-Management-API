package main

import (
	"fmt"
	"log"
	"net/http"

	"smart-grow-management-api/v1/pkg/server"
	"smart-grow-management-api/v1/resources/config"

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
