package persistence

import (
	routemodels "smart-grow-management-api/v1/pkg/internal/route-models"

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
type INutrientsDB interface {
	CreateNutrient(routemodels.CreateNutrientRequest) (*routemodels.CreateNutrientResponse, error)
	DeleteNutrient(routemodels.DeleteNutrientRequest) error
	EditNutrient(routemodels.EditNutrientRequest) (*routemodels.EditNutrientResponse, error)
	GetNutrient(routemodels.GetNutrientRequest) (*routemodels.GetNutrientResponse, error)
	GetAllNutrients(routemodels.GetAllNutrientsRequest) (*routemodels.GetAllNutrientsResponse, error)
}

//
type IOrganizationsDB interface {
	CreateOrganization(routemodels.CreateOrganizationRequest) (*routemodels.CreateOrganizationResponse, error)
	DeleteOrganization(routemodels.DeleteOrganizationRequest) error
	EditOrganization(routemodels.EditOrganizationRequest) (*routemodels.EditOrganizationResponse, error)
	GetOrganization(routemodels.GetOrganizationRequest) (*routemodels.GetOrganizationResponse, error)
	GetAllOrganizations(routemodels.GetAllOrganizationsRequest) (*routemodels.GetAllOrganizationsResponse, error)
}

//
type IPlantCategoriesDB interface {
	CreatePlantCategory(routemodels.CreatePlantCategoryRequest) (*routemodels.CreatePlantCategoryResponse, error)
	DeletePlantCategory(routemodels.DeletePlantCategoryRequest) error
	EditPlantCategory(routemodels.EditPlantCategoryRequest) (*routemodels.EditPlantCategoryResponse, error)
	GetPlantCategory(routemodels.GetPlantCategoryRequest) (*routemodels.GetPlantCategoryResponse, error)
	GetAllPlantCategories(routemodels.GetAllPlantCategoriesRequest) (*routemodels.GetAllPlantCategoriesResponse, error)
}

//
type IPlantLifeCyclesDB interface{}

//
type IPlantStagesDB interface{}

//
type IPlantsDB interface{}

//
type ITasksDB interface{}
