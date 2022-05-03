package persistence

import (
	"github.com/jmoiron/sqlx"
)

//
type Persistence struct {
	Postgres         *sqlx.DB
	GrowSpots        IGrowSpotsDB
	GrowingGroups    IGrowingGroupsDB
	GrowingLevels    IGrowingLevelsDB
	GrowingLocations IGrowingLocationsDB
	Organizations    IOrganizationsDB
	PlantLifeCycles  IPlantLifeCyclesDB
	Plants           IPlantsDB
	Users            IUsersDB
}

//
type IGrowSpotsDB interface{}

//
type IGrowingGroupsDB interface{}

//
type IGrowingLevelsDB interface{}

//
type IGrowingLocationsDB interface{}

//
type IOrganizationsDB interface{}

//
type IPlantLifeCyclesDB interface{}

//
type IPlantsDB interface{}

//
type IUsersDB interface{}