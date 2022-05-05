package logic

import (
	"totally-legit-grow-management/v1/pkg/internal/persistence"
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"
)

type Logic struct {
	Persistence persistence.IPersistence
	ILogic
}

type ILogic interface {
	IDeviceLogic
	IGrowSpotPlantsLogic
	IGrowSpotsLogic
	IGrowingGroupsLogic
	IGrowingLevelsLogic
	IGrowingLocationsLogic
	IGrowingMediumsLogic
	IHarvestsLogic
	IInventoriesLogic
	IMembersLogic
	INutrientsLogic
	IOrganizationsLogic
	IPlantCategoriesLogic
	IPlantLifeCyclesLogic
	IPLantStagesLogic
	IPlantsLogic
	IProductsLogic
	ISeedsLogic
	ITasksLogic
}

//
type IDeviceLogic interface {
	CreateDevice(*routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error)
	CreateDeviceAction(*routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error)
	CreateGrowingGroupDevice(*routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error)
	CreateGrowingLocationDevice(*routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error)
	CreateGrowingLevelDevice(*routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error)
	CreateGrowingSpotDevice(*routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error)
	DeleteDevice(*routemodels.DeleteDeviceRequest) error
	DeleteGrowingGroupDevice(*routemodels.DeleteGrowingGroupDeviceRequest) error
	DeleteGrowingLocationDevice(*routemodels.DeleteGrowingLocationDeviceRequest) error
	DeleteGrowingLevelDevice(*routemodels.DeleteGrowingLevelDeviceRequest) error
	DeleteGrowingSpotDevice(*routemodels.DeleteGrowingSpotDeviceRequest) error
	EditDevice(*routemodels.EditDeviceRequest) (*routemodels.EditDeviceResponse, error)
	GetAllDevices(*routemodels.GetAllDevicesRequest) (*routemodels.GetAllDevicesResponse, error)
	GetAllGrowingGroupDevices(*routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error)
	GetAllGrowingLocationDevices(*routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error)
	GetAllGrowingLevelDevices(*routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error)
	GetAllGrowingSpotDevices(*routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error)
	GetDevice(*routemodels.GetDeviceRequest) (*routemodels.GetDeviceResponse, error)
	GetDeviceActions(*routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error)
	GetGrowingGroupDevice(*routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error)
	GetGrowingLocationDevice(*routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error)
	GetGrowingLevelDevice(*routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error)
	GetGrowingSpotDevice(*routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error)
}

//
type IGrowSpotPlantsLogic interface {
	CreateGrowSpotPlant(*routemodels.CreateGrowSpotPlantRequest) (*routemodels.CreateGrowSpotPlantResponse, error)
	DeleteGrowSpotPlant(*routemodels.DeleteGrowSpotPlantRequest) error
	EditGrowSpotPlant(*routemodels.EditGrowSpotPlantRequest) (*routemodels.EditGrowSpotPlantResponse, error)
	GetGrowSpotPlant(*routemodels.GetGrowSpotPlantRequest) (*routemodels.GetGrowSpotPlantResponse, error)
	GetAllGrowSpotPlants(*routemodels.GetAllGrowSpotPlantsRequest) (*routemodels.GetAllGrowSpotPlantsResponse, error)
}

//
type IGrowSpotsLogic interface {
	CreateGrowSpot(*routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error)
	DeleteGrowSpot(*routemodels.DeleteGrowSpotRequest) error
	EditGrowSpot(*routemodels.EditGrowSpotRequest) (*routemodels.EditGrowSpotResponse, error)
	GetGrowSpotByID(*routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error)
	GetAllGrowSpotsByGrowLevelID(*routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error)
}

//
type IGrowingGroupsLogic interface {
	CreateGrowingGroup(*routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error)
	DeleteGrowingGroup(*routemodels.DeleteGrowingGroupRequest) error
	EditGrowingGroup(*routemodels.EditGrowingGroupRequest) (*routemodels.EditGrowingGroupResponse, error)
	GetGrowingGroupByID(*routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error)
	GetAllGrowingGroupsByOrganizationID(*routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error)
}

//
type IGrowingLevelsLogic interface {
	CreateGrowingLevel(*routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error)
	DeleteGrowingLevel(*routemodels.DeleteGrowingLevelRequest) error
	EditGrowingLevel(*routemodels.EditGrowingLevelRequest) (*routemodels.EditGrowingLevelResponse, error)
	GetGrowingLevelByID(*routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error)
	GetAllGrowingLevelsByGrowingLocationID(*routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error)
}

//
type IGrowingLocationsLogic interface {
	CreateGrowingLocation(*routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error)
	DeleteGrowingLocation(*routemodels.DeleteGrowingLocationRequest) error
	EditGrowingLocation(*routemodels.EditGrowingLocationRequest) (*routemodels.EditGrowingLocationResponse, error)
	GetGrowingLocationByID(*routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error)
	GetAllGrowingLocationsByGrowingGroupID(*routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error)
}

//
type IGrowingMediumsLogic interface {
	CreateGrowingMedium(*routemodels.CreateGrowingMediumRequest) (*routemodels.CreateGrowingMediumResponse, error)
	DeleteGrowingMedium(*routemodels.DeleteGrowingMediumRequest) error
	EditGrowingMedium(*routemodels.EditGrowingMediumRequest) (*routemodels.EditGrowingMediumResponse, error)
	GetGrowingMedium(*routemodels.GetGrowingMediumRequest) (*routemodels.GetGrowingMediumResponse, error)
	GetAllGrowingMediums(*routemodels.GetAllGrowingMediumsRequest) (*routemodels.GetAllGrowingMediumsResponse, error)
}

//
type IHarvestsLogic interface {
	CreateHarvest(*routemodels.CreateHarvestRequest) (*routemodels.CreateHarvestResponse, error)
	DeleteHarvest(*routemodels.DeleteHarvestRequest) error
	EditHarvest(*routemodels.EditHarvestRequest) (*routemodels.EditHarvestResponse, error)
	GetHarvest(*routemodels.GetHarvestRequest) (*routemodels.GetHarvestResponse, error)
	GetAllHarvests(*routemodels.GetAllHarvestsRequest) (*routemodels.GetAllHarvestsResponse, error)
}

//
type IInventoriesLogic interface {
	CreateInventory(*routemodels.CreateInventoryRequest) (*routemodels.CreateInventoryResponse, error)
	DeleteInventory(*routemodels.DeleteInventoryRequest) error
	EditInventory(*routemodels.EditInventoryRequest) (*routemodels.EditInventoryResponse, error)
	GetInventory(*routemodels.GetInventoryRequest) (*routemodels.GetInventoryResponse, error)
	GetAllInventories(*routemodels.GetAllInventoriesRequest) (*routemodels.GetAllInventoriesResponse, error)
}

//
type IMembersLogic interface {
	CreateMember(*routemodels.CreateMemberRequest) (*routemodels.CreateMemberResponse, error)
	DeleteMember(*routemodels.DeleteMemberRequest) error
	EditMember(*routemodels.EditMemberRequest) (*routemodels.EditMemberResponse, error)
	GetMember(*routemodels.GetMemberRequest) (*routemodels.GetMemberResponse, error)
	GetAllMembers(*routemodels.GetAllMembersRequest) (*routemodels.GetAllMembersResponse, error)
}

//
type INutrientsLogic interface {
	CreateNutrient(*routemodels.CreateNutrientRequest) (*routemodels.CreateNutrientResponse, error)
	DeleteNutrient(*routemodels.DeleteNutrientRequest) error
	EditNutrient(*routemodels.EditNutrientRequest) (*routemodels.EditNutrientResponse, error)
	GetNutrient(*routemodels.GetNutrientRequest) (*routemodels.GetNutrientResponse, error)
	GetAllNutrients(*routemodels.GetAllNutrientsRequest) (*routemodels.GetAllNutrientsResponse, error)
}

//
type IOrganizationsLogic interface {
	CreateOrganization(*routemodels.CreateOrganizationRequest) (*routemodels.CreateOrganizationResponse, error)
	DeleteOrganization(*routemodels.DeleteOrganizationRequest) error
	EditOrganization(*routemodels.EditOrganizationRequest) (*routemodels.EditOrganizationResponse, error)
	GetOrganization(*routemodels.GetOrganizationRequest) (*routemodels.GetOrganizationResponse, error)
	GetAllOrganizations(*routemodels.GetAllOrganizationsRequest) (*routemodels.GetAllOrganizationsResponse, error)
}

//
type IPlantCategoriesLogic interface {
	CreatePlantCategory(*routemodels.CreatePlantCategoryRequest) (*routemodels.CreatePlantCategoryResponse, error)
	DeletePlantCategory(*routemodels.DeletePlantCategoryRequest) error
	EditPlantCategory(*routemodels.EditPlantCategoryRequest) (*routemodels.EditPlantCategoryResponse, error)
	GetPlantCategory(*routemodels.GetPlantCategoryRequest) (*routemodels.GetPlantCategoryResponse, error)
	GetAllPlantCategories(*routemodels.GetAllPlantCategoriesRequest) (*routemodels.GetAllPlantCategoriesResponse, error)
}

//
type IPlantLifeCyclesLogic interface {
	CreatePlantLifeCycle(*routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error)
	DeletePlantLifeCycle(*routemodels.DeletePlantLifeCycleRequest) error
	EditPlantLifeCycle(*routemodels.EditPlantLifeCycleRequest) (*routemodels.EditPlantLifeCycleResponse, error)
	GetPlantLifeCycleByID(*routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error)
	GetAllPlantLifeCycles(*routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error)
}

//
type IPLantStagesLogic interface {
	CreatePlantStage(*routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error)
	CreatePlantStageCare(*routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error)
	CreatePlantStageNutrients(*routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error)
	ConnectPlantStage(*routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error)
	DeletePlantStage(*routemodels.DeletePlantStageRequest) error
	EditPlantStage(*routemodels.EditPlantStageRequest) (*routemodels.EditPlantStageResponse, error)
	//
	GetPlantStageByID(*routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error)
	GetPlantStageCareByID(*routemodels.GetPlantStageCareByIDRequest) (*routemodels.GetPlantStageCareByIDResponse, error)
	GetPlantStageNutrient(*routemodels.GetPlantStageNutrientRequest) (*routemodels.GetPlantStageNutrientResponse, error)
	GetAllPlantStages(*routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error)
}

//
type IPlantsLogic interface {
	CreatePlant(*routemodels.CreatePlantRequest) (*routemodels.CreatePlantResponse, error)
	DeletePlant(*routemodels.DeletePlantRequest) error
	EditPlant(*routemodels.EditPlantRequest) (*routemodels.EditPlantResponse, error)
	GetPlant(*routemodels.GetPlantRequest) (*routemodels.GetPlantResponse, error)
	GetAllPlants(*routemodels.GetAllPlantsRequest) (*routemodels.GetAllPlantsResponse, error)
}

//
type IProductsLogic interface {
	CreateProduct(*routemodels.CreateProductRequest) (*routemodels.CreateProductResponse, error)
	DeleteProduct(*routemodels.DeleteProductRequest) error
	EditProduct(*routemodels.EditProductRequest) (*routemodels.EditProductResponse, error)
	GetProduct(*routemodels.GetProductRequest) (*routemodels.GetProductResponse, error)
	GetAllProducts(*routemodels.GetAllProductsRequest) (*routemodels.GetAllProductsResponse, error)
}

//
type ISeedsLogic interface {
	CreateSeed(*routemodels.CreateSeedRequest) (*routemodels.CreateSeedResponse, error)
	DeleteSeed(*routemodels.DeleteSeedRequest) error
	EditSeed(*routemodels.EditSeedRequest) (*routemodels.EditSeedResponse, error)
	GetSeed(*routemodels.GetSeedRequest) (*routemodels.GetSeedResponse, error)
	GetAllSeeds(*routemodels.GetAllSeedsRequest) (*routemodels.GetAllSeedsResponse, error)
}

//
type ITasksLogic interface {
	CreateTask(*routemodels.CreateTaskRequest) (*routemodels.CreateTaskResponse, error)
	DeleteTask(*routemodels.DeleteTaskRequest) error
	EditTask(*routemodels.EditTaskRequest) (*routemodels.EditTaskResponse, error)
	GetTask(*routemodels.GetTaskRequest) (*routemodels.GetTaskResponse, error)
	GetAllTasks(*routemodels.GetAllTasksRequest) (*routemodels.GetAllTasksResponse, error)
}
