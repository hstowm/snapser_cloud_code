package inventory

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"net/http"
	"snapser/cloudecode/configs"
	"snapser/cloudecode/economy"
	"snapser/cloudecode/response"
	proto "snapser/cloudecode/snapserpb/storage"
	"strconv"
	"strings"
)

type EquipmentData struct {
	UserEquipmentUID string `json:"equipmentUserUID"`
	EquipmentUID     string `json:"EquipmentUID"`
	Level            int    `json:"level"`
	Rarity           int    `json:"equipmentRarity"`
	NumReRoll        int    `json:"numReRoll"`
	NumModifier      int    `json:"numModifier"`
	NumEnhance       int    `json:"timeEnhance"`
}

// CalculateEquipmentReRollFee Return fee is a negative number
func CalculateEquipmentReRollFee(data EquipmentData) map[string]int32 {
	fees := make(map[string]int32)
	//fees[configs.Gold] = -int32(float64(data.Level) * float64(data.Rarity+1) / 2 * math.Pow(configs.EquipmentModifierReRollScale, float64(data.NumReRoll)))
	fees[configs.Diamond] = -int32(data.Level * (data.Rarity + 1))
	return fees
}

// CalculateEquipmentUpgradeFee Return fee is a negative number
func CalculateEquipmentUpgradeFee(data EquipmentData, levelUp int) map[string]int32 {
	fees := make(map[string]int32)
	if data.NumEnhance < 3 {
		fees[configs.Gold] = -int32(100 + data.Level*250*levelUp*(data.Rarity+1))
	} else if data.NumEnhance == 3 {
		fees[configs.Diamond] = -int32(data.Level)
	} else if data.NumEnhance == 4 {
		fees[configs.Diamond] = -int32(data.Level * (data.Rarity + 1))
	}
	return fees
}

func CalculateUpgradeEquipmentFee(level int, rarity int) int32 {
	if level > 0 {
		fmt.Println("Fee to upgrade equipment is: " + strconv.Itoa(int(100+int32((rarity+1)*level*250))))
		return 100 + int32((rarity+1)*level*250)
	} else {
		return 0
	}
}
func CalculateEquipmentSellFee(data EquipmentData) map[string]int32 {
	fees := make(map[string]int32)
	var sellCost int32
	sellCost = int32(data.Level * 10 * (data.Rarity + 1))
	for i := 1; i <= data.NumEnhance; i++ {
		sellCost += int32(0.25 * (float64(CalculateUpgradeEquipmentFee(data.Level-i*2, data.Rarity))))
	}
	fees[configs.Gold] = sellCost
	return fees
}
func (pgs *Router) GetPlayerAllEquipment(userID string, c context.Context) ([]EquipmentData, error) {
	blob, err := pgs.Route.StorageClient.GetBlob(c, &proto.GetBlobRequest{AccessType: "private", Key: configs.EquipmentDataBlob, OwnerId: userID})
	if err != nil {
		return nil, errors.New("error when call api to snapser server " + err.Error())
	}
	var b map[string][]EquipmentData
	jsonData := blob.GetValue()
	err = json.Unmarshal([]byte(jsonData), &b)
	if err != nil {
		return nil, errors.New("can't parse data from json to list equipment")
	}
	var equipments = b["data"]
	return equipments, err
}

func (pgs *Router) GetEquipmentWithUID(userEquipmentID string, userId string, c context.Context) (*EquipmentData, error) {

	equipments, err := pgs.GetPlayerAllEquipment(userId, c)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(equipments); i++ {
		if equipments[i].UserEquipmentUID == userEquipmentID {
			return &equipments[i], nil
		}
	}
	return nil, errors.New("Can not found equipment with Id: " + userEquipmentID + "in list " + strconv.Itoa(len(equipments)))
}

// @Summary      Sell equipment
// @Description  Sell player equipment
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	UserEquipmentUID	formData	string	true	"id of equipment player want to sell"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/sell-equipment [post]
func (pgs *Router) SellEquipment(ctx *gin.Context) {
	userEquipmentID := ctx.PostForm(configs.UserEquipmentUID)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userEquipmentID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	equipment, err := pgs.GetEquipmentWithUID(userEquipmentID, userId, c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	fees := CalculateEquipmentSellFee(*equipment)
	currencies, err := economy.UpdateCurrencies(fees, userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code: http.StatusInternalServerError, Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "Sell equipment success " + strconv.Itoa(int(currencies[configs.Gold])),
		Currencies: currencies,
	})
	return
}

// @Summary      Re-Roll equipment modifier
// @Description  Re-roll player equipment modifier
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	UserEquipmentUID	formData	string	true	"id of equipment player want to re-roll"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/re-roll-equipment-modifier [post]
func (pgs *Router) ReRollEquipmentModifier(ctx *gin.Context) {
	userEquipmentID := ctx.PostForm(configs.UserEquipmentUID)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userEquipmentID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	equipment, err := pgs.GetEquipmentWithUID(userEquipmentID, userId, c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	fees := CalculateEquipmentReRollFee(*equipment)
	currencies, err := economy.UpdateCurrencies(fees, userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code: http.StatusInternalServerError, Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "Re roll modifier success" + strconv.Itoa(int(currencies[configs.Gold])),
		Currencies: currencies,
	})
	return
}

// @Summary      Level up equipment
// @Description  Level up equipment
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	userEquipmentUID	formData	string	true	"id of equipment player want to level up"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/level-up-equipment [post]
func (pgs *Router) LevelUpEquipment(ctx *gin.Context) {
	userEquipmentID := ctx.PostForm(configs.UserEquipmentUID)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	levelUpgrade, err := strconv.Atoi(ctx.PostForm("levelUpgrade"))
	if userEquipmentID == "" || userId == "" || sessionToken == "" || err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code: http.StatusBadRequest, Message: "not enough data",
		})
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	equipment, err := pgs.GetEquipmentWithUID(userEquipmentID, userId, c)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	if equipment.NumEnhance >= configs.MaxUpgradeTime {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code: http.StatusBadRequest, Message: "Equipment upgrade time is max",
		})
		return
	}
	fees := CalculateEquipmentUpgradeFee(*equipment, levelUpgrade)
	currencies, err := economy.UpdateCurrencies(fees, userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code: http.StatusInternalServerError, Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "upgrade equipment success" + strconv.Itoa(int(currencies[configs.Gold])),
		Currencies: currencies,
	})
	return
}

// @Summary      Recycle champion
// @Description  Recycle champion
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	equipment	formData	string	true	"list equipment want to recycle"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/recycle-equipment [post]
func (pgs *Router) RecycleEquipment(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	equipments := ctx.PostForm("equipments")
	if userId == "" || sessionToken == "" || equipments == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	userEquipmentsRecycle := strings.Split(equipments, ",")
	if len(userEquipmentsRecycle) < configs.NumEquipmentToRecycle {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: "Not enough number equipment"})
		return
	}
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	userEquipments, err := pgs.GetPlayerAllEquipment(userId, c)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	if len(userEquipments) < configs.NumEquipmentToRecycle {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: "Player don't have enough equipment"})
	}
	for _, equipmentID := range userEquipmentsRecycle {
		index := -1
		for i, equipment := range userEquipments {
			if equipment.UserEquipmentUID == equipmentID {
				index = i
			}
		}
		if index >= 0 {
			copy(userEquipments[index:], userEquipments[index+1:])
			userEquipments[len(userEquipments)-1] = userEquipments[0] // or the zero value of T
			userEquipments = userEquipments[:len(userEquipments)-1]
		} else {
			ctx.JSON(http.StatusOK, response.ResMessage{
				Code:    http.StatusBadRequest,
				Message: "Player don't own equipment with id: " + equipmentID,
			})
		}
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:    http.StatusOK,
		Message: "can merger",
	})
	return
}
