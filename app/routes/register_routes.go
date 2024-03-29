// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, svr *server.Server) error {

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

	deviceRoutes(r, svr)

	/**********************************************************************
	/
	/	Grow Spot Routes
	/
	/**********************************************************************/

	growSpotRoutes(r, svr)

	/**********************************************************************
	/
	/	Grow Spot Plants Routes
	/
	/**********************************************************************/

	growSpotPlantsRoutes(r, svr)

	/**********************************************************************
	/
	/	Growing Group Routes
	/
	/**********************************************************************/

	growingGroupRoutes(r, svr)

	/**********************************************************************
	/
	/	Growing Level Routes
	/
	/**********************************************************************/

	growingLevelRoutes(r, svr)

	/**********************************************************************
	/
	/	Grow Locations Routes
	/
	/**********************************************************************/

	growLocationRoutes(r, svr)

	/**********************************************************************
	/
	/	Growing Medium Routes
	/
	/**********************************************************************/

	growingMediumRoutes(r, svr)

	/**********************************************************************
	/
	/	Harvest Routes
	/
	/**********************************************************************/

	harvestRoutes(r, svr)

	/**********************************************************************
	/
	/	Inventory Routes
	/
	/**********************************************************************/

	inventoryRoutes(r, svr)

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

	nutrientRoutes(r, svr)

	/**********************************************************************
	/
	/	Organization Routes
	/
	/**********************************************************************/

	organizationRoutes(r, svr)

	/**********************************************************************
	/
	/	Plant Categories Routes
	/
	/**********************************************************************/

	plantCategoriesRoutes(r, svr)

	/**********************************************************************
	/
	/	Plant Life Cycles Routes
	/
	/**********************************************************************/

	plantLifeCycleRoutes(r, svr)

	/**********************************************************************
	/
	/	Plant Stages Routes
	/
	/**********************************************************************/

	plantStageRoutes(r, svr)

	/**********************************************************************
	/
	/	Plant Routes
	/
	/**********************************************************************/

	plantRoutes(r, svr)

	/**********************************************************************
	/
	/	Products Routes
	/
	/**********************************************************************/

	productRoutes(r, svr)

	/**********************************************************************
	/
	/	Seeds Routes
	/
	/**********************************************************************/

	seedRoutes(r, svr)

	/**********************************************************************
	/
	/	Task Routes
	/
	/**********************************************************************/

	taskRoutes(r, svr)

	return nil
}
