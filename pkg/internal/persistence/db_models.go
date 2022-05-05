package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

//
type Persistence struct {
	Postgres *sqlx.DB
	IPersistence
}

//
type IPersistence interface {
	MigrateDBUP() error
	MigrateDBDown() error
	StartTransaction() (*sqlx.Tx, error)
	IDevicesDB
	IGrowSpotsDB
	IGrowSpotPlantsDB
	IGrowingGroupsDB
	IGrowingLevelsDB
	IGrowingLocationsDB
	IGrowingMediumsDB
	IHarvestsDB
	IInventoriesDB
	IMembersDB
	INutrientsDB
	IOrganizationsDB
	IPlantCategoriesDB
	IPlantLifeCyclesDB
	IPlantStagesDB
	IPlantsDB
	IProductsDB
	ISeedsDB
	ITasksDB
}

//
type IDevicesDB interface {
	CreateDevice(*routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error)
	CreateDeviceWithTransaction(*sqlx.Tx, *routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error)
	CreateDeviceAction(*routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error)
	CreateDeviceActionWithTransaction(*sqlx.Tx, *routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error)
	CreateGrowingGroupDevice(*routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error)
	CreateGrowingGroupDeviceWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error)
	CreateGrowingLocationDevice(*routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error)
	CreateGrowingLocationDeviceWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error)
	CreateGrowingLevelDevice(*routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error)
	CreateGrowingLevelDeviceWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error)
	CreateGrowingSpotDevice(*routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error)
	CreateGrowingSpotDeviceWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error)
	DeleteDevice(*routemodels.DeleteDeviceRequest) error
	DeleteGrowingGroupDevice(*routemodels.DeleteGrowingGroupDeviceRequest) error
	DeleteGrowingLocationDevice(*routemodels.DeleteGrowingLocationDeviceRequest) error
	DeleteGrowingLevelDevice(*routemodels.DeleteGrowingLevelDeviceRequest) error
	DeleteGrowingSpotDevice(*routemodels.DeleteGrowingSpotDeviceRequest) error
	EditDevice(*routemodels.EditDeviceRequest) (*routemodels.EditDeviceResponse, error)
	//
	GetAllDevices(*routemodels.GetAllDevicesRequest) (*routemodels.GetAllDevicesResponse, error)
	GetAllGrowingGroupDevices(*routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error)
	GetAllGrowingGroupDeviceWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error)
	GetAllGrowingLocationDevices(*routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error)
	GetAllGrowingLocationDeviceWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error)
	GetAllGrowingLevelDevices(*routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error)
	GetAllGrowingLevelDeviceWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error)
	GetAllGrowingSpotDevices(*routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error)
	GetAllGrowingSpotDeviceWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error)
	//
	GetDevice(*routemodels.GetDeviceRequest) (*routemodels.GetDeviceResponse, error)
	GetDeviceActions(*routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error)
	GetDeviceActionsWithTransaction(*sqlx.Tx, *routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error)
	GetGrowingGroupDevice(*routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error)
	GetGrowingGroupDeviceWithTransaction(*sqlx.Tx, *routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error)
	GetGrowingLocationDevice(*routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error)
	GetGrowingLocationDeviceWithTransaction(*sqlx.Tx, *routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error)
	GetGrowingLevelDevice(*routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error)
	GetGrowingLevelDeviceWithTransaction(*sqlx.Tx, *routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error)
	GetGrowingSpotDevice(*routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error)
	GetGrowingSpotDeviceWithTransaction(*sqlx.Tx, *routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error)
}

//
type IGrowSpotPlantsDB interface {
	CreateGrowSpotPlant(*routemodels.CreateGrowSpotPlantRequest) (*routemodels.CreateGrowSpotPlantResponse, error)
	DeleteGrowSpotPlant(*routemodels.DeleteGrowSpotPlantRequest) error
	EditGrowSpotPlant(*routemodels.EditGrowSpotPlantRequest) (*routemodels.EditGrowSpotPlantResponse, error)
	GetGrowSpotPlant(*routemodels.GetGrowSpotPlantRequest) (*routemodels.GetGrowSpotPlantResponse, error)
	GetAllGrowSpotPlants(*routemodels.GetAllGrowSpotPlantsRequest) (*routemodels.GetAllGrowSpotPlantsResponse, error)
}

//
type IGrowSpotsDB interface {
	CreateGrowSpot(*routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error)
	CreateGrowSpotWithTransaction(*sqlx.Tx, *routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error)
	DeleteGrowSpot(*routemodels.DeleteGrowSpotRequest) error
	EditGrowSpot(*routemodels.EditGrowSpotRequest) (*routemodels.EditGrowSpotResponse, error)
	GetGrowSpotByID(*routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error)
	GetGrowSpotByIDWithTransaction(*sqlx.Tx, *routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error)
	GetAllGrowSpotsByGrowLevelID(*routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error)
	GetAllGrowSpotsByGrowLevelIDWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error)
}

//
type IGrowingGroupsDB interface {
	CreateGrowingGroup(*routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error)
	CreateGrowingGroupWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error)
	DeleteGrowingGroup(*routemodels.DeleteGrowingGroupRequest) error
	EditGrowingGroup(*routemodels.EditGrowingGroupRequest) (*routemodels.EditGrowingGroupResponse, error)
	GetGrowingGroupByID(*routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error)
	GetGrowingGroupByIDWithTransaction(*sqlx.Tx, *routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error)
	GetAllGrowingGroupsByOrganizationID(*routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error)
	GetAllGrowingGroupsByOrganizationIDWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error)
}

//
type IGrowingLevelsDB interface {
	CreateGrowingLevel(*routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error)
	CreateGrowingLevelWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error)
	DeleteGrowingLevel(*routemodels.DeleteGrowingLevelRequest) error
	EditGrowingLevel(*routemodels.EditGrowingLevelRequest) (*routemodels.EditGrowingLevelResponse, error)
	GetGrowingLevelByID(*routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error)
	GetGrowingLevelByIDWithTransaction(*sqlx.Tx, *routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error)
	GetAllGrowingLevelsByGrowingLocationID(*routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error)
	GetAllGrowingLevelsByGrowingLocationIDWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error)
}

//
type IGrowingLocationsDB interface {
	CreateGrowingLocation(*routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error)
	CreateGrowingLocationWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error)
	DeleteGrowingLocation(*routemodels.DeleteGrowingLocationRequest) error
	EditGrowingLocation(*routemodels.EditGrowingLocationRequest) (*routemodels.EditGrowingLocationResponse, error)
	GetGrowingLocationByID(*routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error)
	GetGrowingLocationByIDWithTransaction(*sqlx.Tx, *routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error)
	GetAllGrowingLocationsByGrowingGroupID(*routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error)
	GetAllGrowingLocationsByGrowingGroupIDWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error)
}

//
type IGrowingMediumsDB interface {
	CreateGrowingMedium(*routemodels.CreateGrowingMediumRequest) (*routemodels.CreateGrowingMediumResponse, error)
	DeleteGrowingMedium(*routemodels.DeleteGrowingMediumRequest) error
	EditGrowingMedium(*routemodels.EditGrowingMediumRequest) (*routemodels.EditGrowingMediumResponse, error)
	GetGrowingMedium(*routemodels.GetGrowingMediumRequest) (*routemodels.GetGrowingMediumResponse, error)
	GetAllGrowingMediums(*routemodels.GetAllGrowingMediumsRequest) (*routemodels.GetAllGrowingMediumsResponse, error)
}

//
type IHarvestsDB interface {
	CreateHarvest(*routemodels.CreateHarvestRequest) (*routemodels.CreateHarvestResponse, error)
	DeleteHarvest(*routemodels.DeleteHarvestRequest) error
	EditHarvest(*routemodels.EditHarvestRequest) (*routemodels.EditHarvestResponse, error)
	GetHarvest(*routemodels.GetHarvestRequest) (*routemodels.GetHarvestResponse, error)
	GetAllHarvests(*routemodels.GetAllHarvestsRequest) (*routemodels.GetAllHarvestsResponse, error)
}

//
type IInventoriesDB interface {
	CreateInventory(*routemodels.CreateInventoryRequest) (*routemodels.CreateInventoryResponse, error)
	DeleteInventory(*routemodels.DeleteInventoryRequest) error
	EditPlant(*routemodels.EditInventoryRequest) (*routemodels.EditInventoryResponse, error)
	GetInventory(*routemodels.GetInventoryRequest) (*routemodels.GetInventoryResponse, error)
	GetAllInventories(*routemodels.GetAllInventoriesRequest) (*routemodels.GetAllInventoriesResponse, error)
}

//
type IMembersDB interface {
	CreateMember(*routemodels.CreateMemberRequest) (*routemodels.CreateMemberResponse, error)
	DeleteMember(*routemodels.DeleteMemberRequest) error
	EditMember(*routemodels.EditMemberRequest) (*routemodels.EditMemberResponse, error)
	GetMember(*routemodels.GetMemberRequest) (*routemodels.GetMemberResponse, error)
	GetAllMembers(*routemodels.GetAllMembersRequest) (*routemodels.GetAllMembersResponse, error)
}

//
type INutrientsDB interface {
	CreateNutrient(*routemodels.CreateNutrientRequest) (*routemodels.CreateNutrientResponse, error)
	DeleteNutrient(*routemodels.DeleteNutrientRequest) error
	EditNutrient(*routemodels.EditNutrientRequest) (*routemodels.EditNutrientResponse, error)
	GetNutrient(*routemodels.GetNutrientRequest) (*routemodels.GetNutrientResponse, error)
	GetAllNutrients(*routemodels.GetAllNutrientsRequest) (*routemodels.GetAllNutrientsResponse, error)
}

//
type IOrganizationsDB interface {
	CreateOrganization(*routemodels.CreateOrganizationRequest) (*routemodels.CreateOrganizationResponse, error)
	DeleteOrganization(*routemodels.DeleteOrganizationRequest) error
	EditOrganization(*routemodels.EditOrganizationRequest) (*routemodels.EditOrganizationResponse, error)
	GetOrganization(*routemodels.GetOrganizationRequest) (*routemodels.GetOrganizationResponse, error)
	GetAllOrganizations(*routemodels.GetAllOrganizationsRequest) (*routemodels.GetAllOrganizationsResponse, error)
}

//
type IPlantCategoriesDB interface {
	CreatePlantCategory(*routemodels.CreatePlantCategoryRequest) (*routemodels.CreatePlantCategoryResponse, error)
	DeletePlantCategory(*routemodels.DeletePlantCategoryRequest) error
	EditPlantCategory(*routemodels.EditPlantCategoryRequest) (*routemodels.EditPlantCategoryResponse, error)
	GetPlantCategory(*routemodels.GetPlantCategoryRequest) (*routemodels.GetPlantCategoryResponse, error)
	GetAllPlantCategories(*routemodels.GetAllPlantCategoriesRequest) (*routemodels.GetAllPlantCategoriesResponse, error)
}

//
type IPlantLifeCyclesDB interface {
	CreatePlantLifeCycle(*routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error)
	CreatePlantLifeCycleWithTransaction(*sqlx.Tx, *routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error)
	DeletePlantLifeCycle(*routemodels.DeletePlantLifeCycleRequest) error
	EditPlantLifeCycle(*routemodels.EditPlantLifeCycleRequest) (*routemodels.EditPlantLifeCycleResponse, error)
	GetPlantLifeCycleByID(*routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error)
	GetPlantLifeCycleByIDWithTransaction(*sqlx.Tx, *routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error)
	GetAllPlantLifeCycles(*routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error)
	GetAllPlantLifeCyclesWithTransaction(*sqlx.Tx, *routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error)
}

//
type IPlantStagesDB interface {
	CreatePlantStage(*routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error)
	CreatePlantStageWithTransaction(*sqlx.Tx, *routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error)
	CreatePlantStageCare(*routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error)
	CreatePlantStageCareWithTransaction(*sqlx.Tx, *routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error)
	CreatePlantStageNutrients(*routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error)
	CreatePlantStageNutrientsWithTransaction(*sqlx.Tx, *routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error)
	ConnectPlantStage(*routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error)
	ConnectPlantStageWithTransaction(*sqlx.Tx, *routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error)
	//
	DeletePlantStage(*routemodels.DeletePlantStageRequest) error
	EditPlantStage(*routemodels.EditPlantStageRequest) (*routemodels.EditPlantStageResponse, error)
	//
	GetPlantStageByID(*routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error)
	GetPlantStageByIDWithTransaction(*sqlx.Tx, *routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error)
	GetPlantStageCareByID(*routemodels.GetPlantStageCareByIDRequest) (*routemodels.GetPlantStageCareByIDResponse, error)
	GetPlantStageCareByPlantStageIDWithTransaction(*sqlx.Tx, *routemodels.GetPlantStageCareByPlantStageIDRequest) (*routemodels.GetPlantStageCareByPlantStageIDResponse, error)
	GetPlantStageNutrient(*routemodels.GetPlantStageNutrientRequest) (*routemodels.GetPlantStageNutrientResponse, error)
	//
	GetAllPlantStages(*routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error)
	GetAllPlantStagesWithTransaction(*sqlx.Tx, *routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error)
}

//
type IPlantsDB interface {
	CreatePlant(*routemodels.CreatePlantRequest) (*routemodels.CreatePlantResponse, error)
	DeletePlant(*routemodels.DeletePlantRequest) error
	EditPlant(*routemodels.EditPlantRequest) (*routemodels.EditPlantResponse, error)
	GetPlant(*routemodels.GetPlantRequest) (*routemodels.GetPlantResponse, error)
	GetAllPlants(*routemodels.GetAllPlantsRequest) (*routemodels.GetAllPlantsResponse, error)
}

//
type IProductsDB interface {
	CreateProduct(*routemodels.CreateProductRequest) (*routemodels.CreateProductResponse, error)
	DeleteProduct(*routemodels.DeleteProductRequest) error
	EditPlant(*routemodels.EditProductRequest) (*routemodels.EditProductResponse, error)
	GetProduct(*routemodels.GetProductRequest) (*routemodels.GetProductResponse, error)
	GetAllProducts(*routemodels.GetAllProductsRequest) (*routemodels.GetAllProductsResponse, error)
}

//
type ISeedsDB interface {
	CreateSeed(*routemodels.CreateSeedRequest) (*routemodels.CreateSeedResponse, error)
	DeleteSeed(*routemodels.DeleteSeedRequest) error
	EditSeed(*routemodels.EditSeedRequest) (*routemodels.EditSeedResponse, error)
	GetSeed(*routemodels.GetSeedRequest) (*routemodels.GetSeedResponse, error)
	GetAllSeeds(*routemodels.GetAllSeedsRequest) (*routemodels.GetAllSeedsResponse, error)
}

//
type ITasksDB interface {
	CreateTask(*routemodels.CreateTaskRequest) (*routemodels.CreateTaskResponse, error)
	DeleteTask(*routemodels.DeleteTaskRequest) error
	EditTask(*routemodels.EditTaskRequest) (*routemodels.EditTaskResponse, error)
	GetTask(*routemodels.GetTaskRequest) (*routemodels.GetTaskResponse, error)
	GetAllTasks(*routemodels.GetAllTasksRequest) (*routemodels.GetAllTasksResponse, error)
}
