package inventory

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"math"
	"net/http"
	"snapser/cloudecode/configs"
	"snapser/cloudecode/economy"
	"snapser/cloudecode/response"
	proto "snapser/cloudecode/snapserpb/storage"
	"strconv"
	"strings"
)

type ChampionData struct {
	UserChampionUID   string   `json:"UserChampionUID"`
	ChampionUID       string   `json:"championUID"`
	TraitUIDs         []string `json:"traitUIDs"`
	Level             int      `json:"level"`
	Star              int      `json:"star"`
	UnlockedEquipSlot int      `json:"unlockedEquipSlot"`
	UnlockedCardSlot  int      `json:"unlockedCardSlot"`
}

// CalculateChampionUpgradeFee Return fee is a negative number
func CalculateChampionUpgradeFee(data ChampionData) map[string]int32 {
	fees := make(map[string]int32)
	levelUpChampionFee := 0
	if data.Level >= 30 {
		levelUpChampionFee = int(configs.AdvanceLevelUpChampionCost * math.Pow(configs.AdvanceLevelUpChampionScale, float64(data.Level-1)))
	} else {
		levelUpChampionFee = int(configs.BaseLevelUpChampionCost * math.Pow(configs.BaseLevelUpChampionScale, float64(data.Level-1)))
	}
	fees[configs.Gold] = -int32(levelUpChampionFee)
	return fees
}

func (pgs *Router) GetUserChampion(userID string, c context.Context) ([]ChampionData, error) {
	blob, err := pgs.Route.StorageClient.GetBlob(c, &proto.GetBlobRequest{AccessType: "private", Key: configs.ChampionDataBlob, OwnerId: userID})
	if err != nil {
		return nil, errors.New("error when call api to snapser server")
	}
	var b map[string][]ChampionData
	jsonData := blob.GetValue()
	err = json.Unmarshal([]byte(jsonData), &b)
	if err != nil {
		return nil, errors.New("Cannot parse data from json to list champion ")
	}
	champions, isExist := b["data"]
	if !isExist {
		return nil, errors.New("data champion is empty")
	}
	return champions, err
}

func (pgs *Router) GetChampionWithUID(userID string, UserChampionUID string, c context.Context) (*ChampionData, error) {

	champions, err := pgs.GetUserChampion(userID, c)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(champions); i++ {
		if champions[i].UserChampionUID == UserChampionUID {
			return &champions[i], nil
		}
	}
	return nil, errors.New("Can not found equipment with Id: " + UserChampionUID)
}

// @Summary      Level up champion
// @Description  Level up champion
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	userChampionUid	formData	string	true	"champion userID"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/level-up-champion [post]
func (pgs *Router) LevelUpChampion(ctx *gin.Context) {
	userChampionUID := ctx.PostForm(configs.UserChampionUid)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)

	if userChampionUID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	champion, err := pgs.GetChampionWithUID(userId, userChampionUID, c)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	if champion.Level >= configs.ChampionMaxLevel {
		ctx.String(http.StatusBadRequest, "champion is max level")
		return
	}
	fees := CalculateChampionUpgradeFee(*champion)
	// Update diamond amount
	currencies, err := economy.UpdateCurrencies(fees, userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code: http.StatusBadRequest, Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "upgrade champion success" + strconv.Itoa(int(currencies[configs.Gold])),
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
// @Param	champions	formData	string	true	"list champion want to recycle"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/recycle-champion [post]
func (pgs *Router) RecycleChampion(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	champions := ctx.PostForm("champions")
	if userId == "" || sessionToken == "" || champions == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	userChampionsRecycle := strings.Split(champions, ",")
	if len(userChampionsRecycle) < configs.NumChampionToRecycle {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: "Not enough number champion"})
		return
	}
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	userChampions, err := pgs.GetUserChampion(userId, c)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	if len(userChampions) < configs.NumChampionToRecycle {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: "Player don't have enough champion"})
	}
	for _, championID := range userChampionsRecycle {
		index := -1
		for i, champion := range userChampions {
			if champion.UserChampionUID == championID {
				index = i
			}
		}
		if index >= 0 {
			copy(userChampions[index:], userChampions[index+1:])
			userChampions[len(userChampions)-1] = userChampions[0] // or the zero value of T
			userChampions = userChampions[:len(userChampions)-1]
		} else {
			ctx.JSON(http.StatusOK, response.ResMessage{
				Code:    http.StatusBadRequest,
				Message: "Player don't own champion with id: " + championID,
			})
		}
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:    http.StatusOK,
		Message: "can merger",
	})
	return
}
