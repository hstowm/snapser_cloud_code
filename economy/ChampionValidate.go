package economy

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"snapser"
	"snapser/cloudecode/configs"
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

func LevelUpChampion(ctx *gin.Context) {
	userChampionUID := ctx.PostForm(configs.UserChampionUid)
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)

	if userChampionUID == "" || userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())

	execute, h, err := apiClient.StorageServiceApi.GetBlob(context.Background(), userId, "private", configs.ChampionDataBlob).Token(sessionToken).Execute()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if h.StatusCode != http.StatusOK {
		ctn, _ := io.ReadAll(h.Body)
		ctx.String(h.StatusCode, string(ctn))
		return
	}
	var b map[string][]ChampionData
	jsonData := execute.GetValue()
	err = json.Unmarshal([]byte(jsonData), &b)
	//err = json.Unmarshal([]byte(b["data"]), &champions)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Cannot parse data from json to list champion")
		return
	}
	var champions = b["data"]
	for i := 0; i < len(champions); i++ {
		if champions[i].UserChampionUID == userChampionUID {
			if champions[i].Level >= configs.ChampionMaxLevel {
				ctx.String(http.StatusBadRequest, "champion is max level")
				return
			}
			data, _, _ := apiClient.InventoryServiceApi.GetUserCurrencies(context.Background(), userId).Token(sessionToken).Execute()
			// Validate gold
			if data == nil {
				ctx.String(http.StatusInternalServerError, "Cannot get currencies data")
				return
			}
			var goldUse int32 = 0
			var diamondUse int32 = 50
			// Update diamond amount
			if diamondUse > 0 {
				if data.GetCurrencies()[configs.Diamond] < diamondUse {
					ctx.String(http.StatusBadRequest, "Not enough diamond to upgrade champion")
					return
				}
				diamond := data.GetCurrencies()[configs.Diamond] - diamondUse
				_, _, err := apiClient.InventoryServiceApi.UpdateUserVirtualCurrency(context.Background(), userId, configs.Diamond).Body(snapser.UpdateUserVirtualCurrencyRequest{
					Amount: &diamond,
				}).Token(sessionToken).Execute()
				if err != nil {
					ctx.String(http.StatusBadRequest, err.Error())
					return
				}
			}
			// Update gold amount
			if goldUse > 0 {
				if data.GetCurrencies()[configs.Gold] < goldUse {
					ctx.String(http.StatusBadRequest, "Not enough gold to upgrade champion")
					return
				}
				gold := data.GetCurrencies()[configs.Gold] - goldUse
				_, _, err := apiClient.InventoryServiceApi.UpdateUserVirtualCurrency(context.Background(), userId, configs.Gold).Body(snapser.UpdateUserVirtualCurrencyRequest{
					Amount: &gold,
				}).Token(sessionToken).Execute()
				if err != nil {
					ctx.String(http.StatusBadRequest, err.Error())
					return
				}
			}
			// Update champion info
			champions[i].Level += 1
			b["data"] = champions
			value, _ := json.Marshal(b)
			_, h2, err := apiClient.StorageServiceApi.ReplaceBlob(context.Background(), userId, "private", configs.ChampionDataBlob).Token(sessionToken).Body(snapser.ReplaceBlobRequest{
				Cas:   execute.Cas,
				Value: string(value),
			}).Execute()
			if err != nil {
				ctx.String(http.StatusBadRequest, err.Error())
				return
			}
			if h2.StatusCode != http.StatusOK {
				ctn, _ := io.ReadAll(h.Body)
				ctx.String(h.StatusCode, string(ctn))
				return
			}
			ctx.String(http.StatusOK, "Update champion success")
			return
		}
	}
	ctx.String(http.StatusNotFound, "Can not found champion with userId: "+userChampionUID)
}
