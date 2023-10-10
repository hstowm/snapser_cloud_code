package economy

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"snapser"
	"snapser/cloudecode/configs"
	"strconv"
)

type EquipmentData struct {
	UserEquipmentUID string
	EquipmentUID     string
	level            int
	rarity           int
	numReRoll        int
	numModifier      int
}

func GetUserCurrencies(ctx *gin.Context) *snapser.InventoryGetUserCurrenciesResponse {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())
	data, _, _ := apiClient.InventoryServiceApi.GetUserCurrencies(context.Background(), userId).Token(sessionToken).Execute()
	return data
}

func UpdateCurrencies(ctx *gin.Context, amountUpdate map[string]int32, currentCurrencies map[string]int32) bool {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())

	for currencyKey, value := range amountUpdate {
		if value+currentCurrencies[currencyKey] < 0 {
			ctx.String(http.StatusBadRequest, "Not enough money to use")
			return false
		}
	}
	for currencyKey, value := range amountUpdate {
		gold := currentCurrencies[currencyKey] + value
		_, res, err := apiClient.InventoryServiceApi.UpdateUserVirtualCurrency(context.Background(), userId, currencyKey).Body(snapser.UpdateUserVirtualCurrencyRequest{
			Amount: &gold,
		}).Token(sessionToken).Execute()
		if err != nil {
			body, _ := io.ReadAll(res.Body)
			ctx.String(res.StatusCode, string(body))
			return false
		}
	}
	return true
}
func GetEquipmentWithUID(ctx *gin.Context) *EquipmentData {
	userEquipmentID := ctx.PostForm(configs.UserEquipmentUID)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userEquipmentID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return nil
	}
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())

	// Get all player champion data
	execute, h, err := apiClient.StorageServiceApi.GetBlob(context.Background(), userId, "private", configs.EquipmentDataBlob).Token(sessionToken).Execute()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return nil
	}
	if h.StatusCode != http.StatusOK {
		ctn, _ := io.ReadAll(h.Body)
		ctx.String(h.StatusCode, string(ctn))
		return nil
	}
	var b map[string][]EquipmentData
	jsonData := execute.GetValue()
	err = json.Unmarshal([]byte(jsonData), &b)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Cannot parse data from json to list champion")
		return nil
	}
	var equipments = b["data"]
	for i := 0; i < len(equipments); i++ {
		if equipments[i].UserEquipmentUID == userEquipmentID {
			return &equipments[i]
		}
	}
	ctx.String(http.StatusNotFound, "Can not found equipment with Id: "+userEquipmentID)
	return nil
}
func GetChampionWithUID(ctx *gin.Context) *ChampionData {
	UserChampionUID := ctx.PostForm(configs.UserChampionUid)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if UserChampionUID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return nil
	}
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())

	// Get all player champion data
	execute, h, err := apiClient.StorageServiceApi.GetBlob(context.Background(), userId, "private", configs.ChampionDataBlob).Token(sessionToken).Execute()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return nil
	}
	if h.StatusCode != http.StatusOK {
		ctn, _ := io.ReadAll(h.Body)
		ctx.String(h.StatusCode, string(ctn))
		return nil
	}
	var b map[string][]ChampionData
	jsonData := execute.GetValue()
	err = json.Unmarshal([]byte(jsonData), &b)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Cannot parse data from json to list champion")
		return nil
	}
	var equipments = b["data"]
	for i := 0; i < len(equipments); i++ {
		if equipments[i].UserChampionUID == UserChampionUID {
			return &equipments[i]
		}
	}
	ctx.String(http.StatusNotFound, "Can not found equipment with Id: "+UserChampionUID)
	return nil
}

// CalculateEquipmentUpgradeFee Return fee is a negative number
func CalculateEquipmentUpgradeFee(data EquipmentData, levelUp int) map[string]int32 {
	fees := make(map[string]int32)
	return fees
}
func CalculateEquipmentSellFee(data EquipmentData) map[string]int32 {
	fees := make(map[string]int32)
	return fees
}

// CalculateChampionUpgradeFee Return fee is a negative number
func CalculateChampionUpgradeFee(data ChampionData) map[string]int32 {
	fees := make(map[string]int32)
	return fees
}
func CalculateCampaignFee(power int, star int) map[string]int32 {
	fees := make(map[string]int32)
	return fees
}
func CalculatePvPFee(enemyElo int) map[string]int32 {
	fees := make(map[string]int32)
	return fees
}

// CalculateEquipmentReRollFee Return fee is a negative number
func CalculateEquipmentReRollFee(data EquipmentData) map[string]int32 {
	fees := make(map[string]int32)
	return fees
}

// @Summary      Level up champion
// @Description  Level up champion
// @Accept       mpfd
// @Produce      mpfd
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	championUserId	formData	string	true	"champion userID"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v2/byosnap-skybull/level-up-champion [post]
func LevelUpChampionFee(ctx *gin.Context) {
	userChampionUID := ctx.PostForm(configs.UserChampionUid)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)

	if userChampionUID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	champion := GetChampionWithUID(ctx)
	if champion == nil {
		return
	}
	if champion.Level >= configs.ChampionMaxLevel {
		ctx.String(http.StatusBadRequest, "champion is max level")
		return
	}
	data := GetUserCurrencies(ctx)
	// Validate gold
	if data == nil {
		ctx.String(http.StatusInternalServerError, "Cannot get currencies data")
		return
	}
	fees := CalculateChampionUpgradeFee(*champion)
	// Update diamond amount
	if !UpdateCurrencies(ctx, fees, data.GetCurrencies()) {
		return
	}
	ctx.String(http.StatusOK, "Update champion success")
	return
}

// @Summary      Level up equipment
// @Description  Level up equipment
// @Accept       mpfd
// @Produce      mpfd
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	UserEquipmentUID	formData	string	true	"id of equipment player want to level up"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v2/byosnap-skybull/level-up-equipment [post]
func LevelUpEquipmentFee(ctx *gin.Context) {
	userEquipmentID := ctx.PostForm(configs.UserEquipmentUID)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	levelUpgrade, err := strconv.Atoi(ctx.PostForm("levelUpgrade"))
	if userEquipmentID == "" || userId == "" || sessionToken == "" || err != nil {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	equipment := GetEquipmentWithUID(ctx)
	if equipment == nil {
		return
	}
	if equipment.level >= configs.EquipmentMaxLevel {
		ctx.String(http.StatusBadRequest, "champion is max level")
		return
	}
	data := GetUserCurrencies(ctx)
	// Validate gold
	if data == nil {
		ctx.String(http.StatusInternalServerError, "Cannot get currencies data")
		return
	}
	fees := CalculateEquipmentUpgradeFee(*equipment, levelUpgrade)
	// Update diamond amount
	if !UpdateCurrencies(ctx, fees, data.GetCurrencies()) {
		return
	}
	ctx.String(http.StatusOK, "Update equipment success")
	return
}

// @Summary      Sell equipment
// @Description  Sell player equipment
// @Accept       mpfd
// @Produce      mpfd
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	UserEquipmentUID	formData	string	true	"id of equipment player want to sell"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v2/byosnap-skybull/sell-equipment [post]
func SellEquipment(ctx *gin.Context) {
	userEquipmentID := ctx.PostForm(configs.UserEquipmentUID)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userEquipmentID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	equipment := GetEquipmentWithUID(ctx)
	if equipment == nil {
		return
	}
	data := GetUserCurrencies(ctx)
	// Validate gold
	if data == nil {
		ctx.String(http.StatusInternalServerError, "Cannot get currencies data")
		return
	}
	fees := CalculateEquipmentSellFee(*equipment)
	if !UpdateCurrencies(ctx, fees, data.GetCurrencies()) {
		return
	}
	ctx.String(http.StatusOK, "Sell equipment success")
	return
}

// @Summary      Re-Roll equipment modifier
// @Description  Re-roll player equipment modifier
// @Accept       mpfd
// @Produce      mpfd
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	UserEquipmentUID	formData	string	true	"id of equipment player want to re-roll"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v2/byosnap-skybull/re-roll-equipment-modifier [post]
func ReRollEquipmentModifier(ctx *gin.Context) {
	userEquipmentID := ctx.PostForm(configs.UserEquipmentUID)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userEquipmentID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	equipment := GetEquipmentWithUID(ctx)
	if equipment == nil {
		return
	}
	data := GetUserCurrencies(ctx)
	// Validate gold
	if data == nil {
		ctx.String(http.StatusInternalServerError, "Cannot get currencies data")
		return
	}
	fees := CalculateEquipmentReRollFee(*equipment)
	if !UpdateCurrencies(ctx, fees, data.GetCurrencies()) {
		return
	}
	ctx.String(http.StatusOK, "Sell equipment success")
	return
}

// @Summary      Call when player win a pvp game
// @Description  Call when player win a pvp game
// @Accept       mpfd
// @Produce      mpfd
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	opponentID	formData	string	true	"opponent snapser id "
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v2/byosnap-skybull/win-pvp [post]
func WinPvPGame(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	opponentId := ctx.PostForm("opponentID")
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userId == "" || sessionToken == "" || opponentId == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())
	execute, _, err := apiClient.ProfilesServiceApi.GetProfile(context.Background(), opponentId).Token(sessionToken).Execute()
	if err != nil {
		return
	}
	data := GetUserCurrencies(ctx)
	// Validate gold
	if data == nil {
		ctx.String(http.StatusInternalServerError, "Cannot get currencies data")
		return
	}
	elo, _ := strconv.Atoi(execute.Profile[configs.EloPoint].(string))
	currenciesEarn := CalculatePvPFee(elo)
	if !UpdateCurrencies(ctx, currenciesEarn, data.GetCurrencies()) {
		return
	}
	ctx.String(http.StatusOK, "Add currency success")
}

// @Summary      Call when win a campaign
// @Description   Call when win a campaign
// @Accept       mpfd
// @Produce      mpfd
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	stagePower	formData	string	true	"power stage player win"
// @Param	star	formData	string	true	"the start they reach"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v2/byosnap-skybull/win-campaign [post]
func WinCampaign(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	stagePower, _ := strconv.Atoi(ctx.PostForm("stagePower"))
	star, _ := strconv.Atoi(ctx.PostForm("star"))
	if userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	currencies := CalculateCampaignFee(stagePower, star)
	data := GetUserCurrencies(ctx)
	// Validate gold
	if data == nil {
		ctx.String(http.StatusInternalServerError, "Cannot get currencies data")
		return
	}
	if !UpdateCurrencies(ctx, currencies, data.GetCurrencies()) {
		return
	}
	ctx.String(http.StatusOK, "Add currency success")
}
