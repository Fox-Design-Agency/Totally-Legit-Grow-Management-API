package persistence

import (
	"github.com/jmoiron/sqlx"
)

//
type Persistence struct {
	Postgres         *sqlx.DB
	Devices          IDevicesDB
	GrowSpots        IGrowSpotsDB
	GrowingGroups    IGrowingGroupsDB
	GrowingLevels    IGrowingLevelsDB
	GrowingLocations IGrowingLocationsDB
	Members          IMembersDB
	Nutrients        INutrientsDB
	Organizations    IOrganizationsDB
	PlantCategories  IPlantCategoriesDB
	PlantLifeCycles  IPlantLifeCyclesDB
	PlantStages      IPlantStagesDB
	Plants           IPlantsDB
	Tasks            ITasksDB
}

//
type IDevicesDB interface{}

//
type IGrowSpotsDB interface{}

//
type IGrowingGroupsDB interface{}

//
type IGrowingLevelsDB interface{}

//
type IGrowingLocationsDB interface{}

//
type IMembersDB interface{}

//
type INutrientsDB interface{}

//
type IOrganizationsDB interface{}

//
type IPlantCategoriesDB interface{}

//
type IPlantLifeCyclesDB interface{}

//
type IPlantStagesDB interface{}

//
type IPlantsDB interface{}

//
type ITasksDB interface{}
