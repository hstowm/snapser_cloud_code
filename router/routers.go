package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"snapser/cloudecode/economy"
	"snapser/cloudecode/internal_connector"
	"snapser/cloudecode/inventory"
	"snapser/cloudecode/profile"
	inventorypb "snapser/cloudecode/snapserpb/inventory"
	profilepb "snapser/cloudecode/snapserpb/profiles"
	statspb "snapser/cloudecode/snapserpb/statistics"
	storagepb "snapser/cloudecode/snapserpb/storage"
	"snapser/cloudecode/storage"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type ROICloudServer struct {
	connector      *internal_connector.SnapserServiceConnector
	MatchingRoute  *storage.Router
	InventoryRoute *inventory.Router
	EconomyRouter  *economy.Router
	ProfileRouter  *profile.Router
}
type Routes []Route

func (pgs *ROICloudServer) NewRouter() *gin.Engine {
	var router = gin.Default()
	{
		router.POST("/v1/byosnap-skybull-cloudcode/level-up-champion", pgs.InventoryRoute.LevelUpChampion)
		router.POST("/v1/byosnap-skybull-cloudcode/level-up-equipment", pgs.InventoryRoute.LevelUpEquipment)
		router.POST("/v1/byosnap-skybull-cloudcode/sell-equipment", pgs.InventoryRoute.SellEquipment)
		router.POST("/v1/byosnap-skybull-cloudcode/re-roll-equipment-modifier", pgs.InventoryRoute.ReRollEquipmentModifier)
		router.POST("/v1/byosnap-skybull-cloudcode/sell-champion", pgs.InventoryRoute.SellEquipment)
		router.POST("/v1/byosnap-skybull-cloudcode/win-pvp", pgs.MatchingRoute.WinPvPGame)
		router.POST("/v1/byosnap-skybull-cloudcode/win-campaign", pgs.MatchingRoute.WinCampaign)
		router.POST("/v1/byosnap-skybull-cloudcode/create-pvp-game", pgs.MatchingRoute.CreatePvpGame)
		router.POST("/v1/byosnap-skybull-cloudcode/create-campaign-game", pgs.MatchingRoute.CreateCampaignGame)
		router.GET("/v1/byosnap-skybull-cloudcode/get-energy", pgs.EconomyRouter.GetEnergy)
		router.GET("/v1/byosnap-skybull-cloudcode/account-info", pgs.ProfileRouter.GetAccountInfo)
		router.GET("/v1/byosnap-skybull-cloudcode/account-info-detail", pgs.ProfileRouter.GetAccountInfo2)
		router.POST("/v1/byosnap-skybull-cloudcode/buy-energy", pgs.EconomyRouter.BuyEnergy)
		router.POST("/v1/byosnap-skybull-cloudcode/recycle-equipment", pgs.InventoryRoute.RecycleEquipment)
		router.POST("/v1/byosnap-skybull-cloudcode/recycle-champion", pgs.InventoryRoute.RecycleChampion)
		router.POST("/v1/byosnap-skybull-cloudcode/recycle-card-skill", pgs.InventoryRoute.RecycleCardSkill)
		router.GET("/v1/byosnap-skybull-cloudcode/daily-login-info", pgs.ProfileRouter.GetDailyRewardCheckingStatus)
		router.GET("/v1/byosnap-skybull-cloudcode/newbie-login-info", pgs.ProfileRouter.GetNewbieRewardCheckingStatus)
		router.POST("/v1/byosnap-skybull-cloudcode/claim-daily", pgs.ProfileRouter.CheckinDaily)
		router.POST("/v1/byosnap-skybull-cloudcode/claim-newbie", pgs.ProfileRouter.CheckinNewbie)
	}
	return router
}
func DeleteAccount() {

}
func New() (*ROICloudServer, error) {
	pgs := &ROICloudServer{}
	connector, err := internal_connector.New()
	pgs.connector = connector
	pgs.CreateServiceRouter()
	return pgs, err
}

// We use GRPC internally to provide fast requests between services.  Users are authenticated at our API Gateway and then trusted once internal;
// The API Gateway will add/override a header called "User-Id" that is the authenticated trusted user and can be relied on internally.
func createInventoryClient() (inventorypb.InventoryServiceClient, error) {
	conn, err := grpc.Dial("service-inventory:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := inventorypb.NewInventoryServiceClient(conn)
	return client, nil
}

func createStatisticsClient() (statspb.StatisticsServiceClient, error) {
	conn, err := grpc.Dial("service-statistics:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := statspb.NewStatisticsServiceClient(conn)
	return client, nil
}
func createStorageClient() (storagepb.StorageServiceClient, error) {
	conn, err := grpc.Dial("service-storage:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := storagepb.NewStorageServiceClient(conn)
	return client, nil
}
func createProfiledClient() (profilepb.ProfilesServiceClient, error) {
	conn, err := grpc.Dial("service-profiles:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := profilepb.NewProfilesServiceClient(conn)
	return client, nil
}
func (pgs *ROICloudServer) CreateServiceRouter() {
	pgs.MatchingRoute = new(storage.Router)
	pgs.InventoryRoute = new(inventory.Router)
	pgs.EconomyRouter = new(economy.Router)
	pgs.ProfileRouter = new(profile.Router)
	pgs.InventoryRoute.Route = pgs.connector
	pgs.MatchingRoute.Route = pgs.connector
	pgs.EconomyRouter.Route = pgs.connector
	pgs.ProfileRouter.Route = pgs.connector
}
