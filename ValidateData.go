package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"snapser"
	"snapser/cloudecode/configs"
	"strconv"
	"time"
)

type energyData struct {
	energy int
	time   string
}

func ValidateEnergyData(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	newEnergy, err := strconv.Atoi(ctx.PostForm("energy"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "text/plain", "energy wrong format")
		return
	}
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())
	res, _response, err := apiClient.ProfilesServiceApi.GetProfile(context.Background(), userId).Token(sessionToken).Execute()
	if _response == nil {
		ctx.String(http.StatusBadRequest, "cannot send request to get profile"+err.Error()+" with userID: ."+userId+" and token: "+sessionToken+".")
		return
	}
	if _response.StatusCode != http.StatusNotFound && _response.StatusCode != http.StatusOK {
		ctn, _ := io.ReadAll(_response.Body)
		ctx.String(http.StatusBadRequest, string(ctn))
		return
	}
	if _response.StatusCode == http.StatusOK {

		if res.Profile != nil {
			energy := energyData{
				energy: int(res.Profile[configs.EnergyKey].(float64)),
				time:   res.Profile[configs.EnergyUpdateKey].(string),
			}
			oldTime, _err := time.Parse(configs.TimeFormatTimeDay, energy.time)
			if _err != nil {
				goto validateDone
			}
			dis := newEnergy - energy.energy
			current := time.Now()

			timeDis := current.Sub(oldTime).Minutes()
			a := timeDis / 5
			if 0 > int(a)-dis {
				ctx.String(http.StatusBadRequest, "energy data is invalid ")
				return
			}
		}
	}
validateDone:
	var profile = make(map[string]interface{})
	if res != nil {
		profile = res.Profile
	}
	profile[configs.EnergyKey] = newEnergy
	profile[configs.EnergyUpdateKey] = time.Now().Format(configs.TimeFormatTimeDay)
	print("current time %s", time.Now().Format(configs.TimeFormatTimeDay))
	// validate new energy
	// Set new energy
	_, status, _ := apiClient.ProfilesServiceApi.UpsertProfile(context.Background(), userId).Token(sessionToken).Body(snapser.UpsertProfileRequest{
		Profile: profile,
	}).Execute()
	if status.StatusCode == http.StatusOK {
		ctx.String(http.StatusOK, "update energy success")
		return
	}
}

func ValidateCurrency(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	currencyType := ctx.PostForm("type")
	gold, err := strconv.Atoi(ctx.PostForm("value"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "text/plain", "currency value invalid!")
		return
	}
	source := ctx.PostForm("goldSource")
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())
	data, _, _ := apiClient.InventoryServiceApi.GetUserCurrencies(context.Background(), userId).Token(sessionToken).Execute()

	// Validate gold
	if data == nil {
		//ctx.String(http.StatusBadRequest, "text/plain", "not found currencies data")
		//return
		print("currency do not setup ")
	}
	print("update gold %d from %s", gold, source)
	updateGold := int32(gold)
	//data.Currencies[currencyType]
	_, res, _ := apiClient.InventoryServiceApi.UpdateUserVirtualCurrency(context.Background(), userId, currencyType).Token(sessionToken).Body(snapser.UpdateUserVirtualCurrencyRequest{
		Amount: &updateGold,
	}).Execute()
	if res == nil {
		ctx.String(http.StatusBadRequest, "cannot update currency!")
		return
	}
	if res.StatusCode == http.StatusOK {
		ctx.String(http.StatusOK, "update currency done!")
	} else {
		ctx.String(http.StatusOK, "update currency fail!")
	}
}
