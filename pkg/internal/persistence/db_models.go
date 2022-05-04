package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

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
type IDevicesDB interface {
	CreateDevice(routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error)
	CreateDeviceWithTransaction(*sqlx.Tx, routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error)
	CreateDeviceAction(routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error)
	CreateDeviceActionWithTransaction(*sqlx.Tx, routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error)
	CreateGrowingGroupDevice(routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error)
	CreateGrowingGroupDeviceWithTransaction(*sqlx.Tx, routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error)
	CreateGrowingLocationDevice(routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error)
	CreateGrowingLocationDeviceWithTransaction(*sqlx.Tx, routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error)
	CreateGrowingLevelDevice(routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error)
	CreateGrowingLevelDeviceWithTransaction(*sqlx.Tx, routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error)
	CreateGrowingSpotDevice(routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error)
	CreateGrowingSpotDeviceWithTransaction(*sqlx.Tx, routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error)
	DeleteDevice(routemodels.DeleteDeviceRequest) error
	DeleteGrowingGroupDevice(routemodels.DeleteGrowingGroupDeviceRequest) error
	DeleteGrowingLocationDevice(routemodels.DeleteGrowingLocationDeviceRequest) error
	DeleteGrowingLevelDevice(routemodels.DeleteGrowingLevelDeviceRequest) error
	DeleteGrowingSpotDevice(routemodels.DeleteGrowingSpotDeviceRequest) error
	EditDevice(routemodels.EditDeviceRequest) (*routemodels.EditDeviceResponse, error)
	//
	GetAllDevices(routemodels.GetAllDevicesRequest) (*routemodels.GetAllDevicesResponse, error)
	GetAllGrowingGroupDevices(routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error)
	GetAllGrowingGroupDeviceWithTransaction(*sqlx.Tx, routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error)
	GetAllGrowingLocationDevices(routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error)
	GetAllGrowingLocationDeviceWithTransaction(*sqlx.Tx, routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error)
	GetAllGrowingLevelDevices(routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error)
	GetAllGrowingLevelDeviceWithTransaction(*sqlx.Tx, routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error)
	GetAllGrowingSpotDevices(routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error)
	GetAllGrowingSpotDeviceWithTransaction(*sqlx.Tx, routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error)
	//
	GetDevice(routemodels.GetDeviceRequest) (*routemodels.GetDeviceResponse, error)
	GetDeviceActions(routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error)
	GetDeviceActionsWithTransaction(*sqlx.Tx, routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error)
	GetGrowingGroupDevice(routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error)
	GetGrowingGroupDeviceWithTransaction(*sqlx.Tx, routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error)
	GetGrowingLocationDevice(routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error)
	GetGrowingLocationDeviceWithTransaction(*sqlx.Tx, routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error)
	GetGrowingLevelDevice(routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error)
	GetGrowingLevelDeviceWithTransaction(*sqlx.Tx, routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error)
	GetGrowingSpotDevice(routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error)
	GetGrowingSpotDeviceWithTransaction(*sqlx.Tx, routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error)
}

//
type IGrowSpotsDB interface{}

//
type IGrowingGroupsDB interface {
	CreateGrowingGroup(routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error)
	CreateGrowingGroupWithTransaction(*sqlx.Tx, routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error)
	DeleteGrowingGroup(routemodels.DeleteGrowingGroupRequest) error
	EditGrowingGroup(routemodels.EditGrowingGroupRequest) (*routemodels.EditGrowingGroupResponse, error)
	GetGrowingGroup(routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error)
	GetGrowingGroupWithTransaction(*sqlx.Tx, routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error)
	GetAllGrowingGroups(routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error)
}

//
type IGrowingLevelsDB interface {
	CreateGrowingLevel(routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error)
	CreateGrowingLevelWithTransaction(*sqlx.Tx, routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error)
	DeleteGrowingLevel(routemodels.DeleteGrowingLevelRequest) error
	EditGrowingLevel(routemodels.EditGrowingLevelRequest) (*routemodels.EditGrowingLevelResponse, error)
	GetGrowingLevel(routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error)
	GetGrowingLevelWithTransaction(*sqlx.Tx, routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error)
	GetAllGrowingLevels(routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error)
}

//
type IGrowingLocationsDB interface {
	CreateGrowingLocation(routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error)
	CreateGrowingLocationWithTransaction(*sqlx.Tx, routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error)
	DeleteGrowingLocation(routemodels.DeleteGrowingLocationRequest) error
	EditGrowingLocation(routemodels.EditGrowingLocationRequest) (*routemodels.EditGrowingLocationResponse, error)
	GetGrowingLocation(routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error)
	GetGrowingLocationWithTransaction(*sqlx.Tx, routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error)
	GetAllGrowingLocations(routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error)
}

//
type IMembersDB interface {
	CreateMember(routemodels.CreateMemberRequest) (*routemodels.CreateMemberResponse, error)
	DeleteMember(routemodels.DeleteMemberRequest) error
	EditMember(routemodels.EditMemberRequest) (*routemodels.EditMemberResponse, error)
	GetMember(routemodels.GetMemberRequest) (*routemodels.GetMemberResponse, error)
	GetAllMembers(routemodels.GetAllMembersRequest) (*routemodels.GetAllMembersResponse, error)
}

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
