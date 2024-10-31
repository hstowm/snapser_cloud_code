package economy

import (
	"fmt"
	"math"
	"net/http"
	"snapser/cloudecode/configs"
	"snapser/cloudecode/internal_connector"
	"snapser/cloudecode/response"
	proto "snapser/cloudecode/snapserpb/inventory"
	proto2 "snapser/cloudecode/snapserpb/storage"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type Router struct {
	Route *internal_connector.SnapserServiceConnector
}

func GetEnergy(userID string, c context.Context, service *internal_connector.SnapserServiceConnector) (map[string]int32, error) {

	blob, err := service.StorageClient.GetBlob(c, &proto2.GetBlobRequest{
		AccessType: "private",
		Key:        configs.EnergyUpdateKey,
		OwnerId:    userID,
	})
	lastUpdatedTime := ""
	currentTime := time.Now()
	energyAdded := 0
	if err != nil {
		fmt.Println("update energy to max")
		lastUpdatedTime = currentTime.Format(configs.TimeFormatTimeDay)
		energyAdded = configs.MaxEnergy
		_, err = service.StorageClient.InsertBlob(c, &proto2.InsertBlobRequest{
			Key:        configs.EnergyUpdateKey,
			AccessType: "private",
			OwnerId:    userID,
			Value:      lastUpdatedTime,
		})
		if err != nil {
			fmt.Println("Cannot update energy update time")
		}
		return UpdateEnergy(userID, int32(energyAdded), false, c, service)
	} else {
		fmt.Println("update energy to by time")
		_, err = service.StorageClient.ReplaceBlob(c, &proto2.ReplaceBlobRequest{
			Key:        configs.EnergyUpdateKey,
			AccessType: "private",
			OwnerId:    userID,
			Value:      currentTime.Format(configs.TimeFormatTimeDay),
			Create:     true,
			Cas:        blob.Cas,
			Ttl:        0,
		})
		if err != nil {
			fmt.Println("Cannot update energy update time with cas: " + blob.Cas + " error " + err.Error())
		}
		lastUpdatedTime = blob.GetValue()
		fmt.Println("last update time : " + lastUpdatedTime)
	}
	energy, err := GetUserCurrencies(userID, c, service)
	if err != nil {
		return nil, err
	}
	if energy[configs.Energy] >= configs.MaxEnergy {
		return energy, nil
	}
	lastUpdateEnergyTime, err := time.Parse(configs.TimeFormatTimeDay, lastUpdatedTime)
	if err != nil {
		lastUpdateEnergyTime = currentTime
		fmt.Println("error" + err.Error())
	}
	currentDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	fmt.Print("day time: " + currentDay.Format(configs.TimeFormatTimeDay))
	timeDistance1 := int(math.Floor(lastUpdateEnergyTime.Sub(currentDay).Minutes() / 5))
	timeDistance2 := int(math.Floor(currentTime.Sub(currentDay).Minutes() / 5))
	fmt.Print("time update 1: " + strconv.Itoa(timeDistance1) + " time update 2 : " + strconv.Itoa(timeDistance2))
	energyAdded = timeDistance2 - timeDistance1
	if energyAdded != 0 {
		// get currencies update time
		fmt.Println("Update currencies amount: " + strconv.Itoa(energyAdded))
		return UpdateEnergy(userID, int32(energyAdded), false, c, service)
	} else {
		fmt.Println("energy not changed")
		return GetUserCurrencies(userID, c, service)
	}
}
func UpdateEnergy(userID string, amountUpdate int32, isOver bool, c context.Context, service *internal_connector.SnapserServiceConnector) (map[string]int32, error) {
	currenciesUpdate := make(map[string]int32)
	currenciesUpdate[configs.Energy] = amountUpdate
	currencies, err := GetUserCurrencies(userID, c, service)
	if err != nil || currencies == nil {
		fmt.Println("Can not get currency data : " + err.Error())
		currencies = make(map[string]int32)
	}
	if !isOver {
		energy, isExist := currencies[configs.Energy]
		if !isExist {
			energy = 0
		}
		amountUpdate = int32(math.Min(float64(amountUpdate), float64(configs.MaxEnergy-(energy))))
	}
	if amountUpdate == 0 {
		return currencies, nil
	}
	fmt.Println("Update energy amount: " + strconv.Itoa(int(amountUpdate)))
	_, err = service.InventoryClient.UpdateUserVirtualCurrency(c, &proto.UpdateUserVirtualCurrencyRequest{
		UserId:       userID,
		CurrencyName: configs.Energy,
		Amount:       amountUpdate,
	})
	currencies[configs.Energy] += amountUpdate
	fmt.Println("energy amount: " + strconv.Itoa(int(currencies[configs.Energy])))
	return currencies, err
}

// @Summary      Get energy
// @Description  Get energy
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	query	string	true	"userId when they login"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/buy-energy [get]
func (r *Router) GetEnergy(ctx *gin.Context) {
	userId, _ := ctx.GetQuery(configs.UserId)

	if userId == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	fmt.Println("Get energy data")
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	energy, err := GetEnergy(userId, c, r.Route)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{
			Code:       http.StatusInternalServerError,
			Message:    err.Error(),
			Currencies: energy,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "Get Energy success",
		Currencies: energy,
	})
}

// @Summary      Buy energy by gold
// @Description  Buy energy by gold
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/buy-energy [post]
func (r *Router) BuyEnergy(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)

	if userId == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	_, err := GetEnergy(userId, c, r.Route)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	// Sub Money
	currenciesUpdate := make(map[string]int32)
	currenciesUpdate[configs.Diamond] = -25
	currencies, err := UpdateCurrencies(currenciesUpdate, userId, c, r.Route)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code:       http.StatusBadRequest,
			Message:    err.Error(),
			Currencies: currencies,
		})
		return
	}
	energy, err := UpdateEnergy(userId, configs.MaxEnergy, false, c, r.Route)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResMessage{
			Code:       http.StatusInternalServerError,
			Message:    "Buy energy fail with error: " + err.Error(),
			Currencies: currencies,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "Get Energy success",
		Currencies: energy,
	})
}
