package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"snapser"
	"strconv"
	"time"
)

type energyData struct {
	energy int
	time   string
}

const energy_key = "energy"
const energy_update_key = "energy_update_time"
const time_fomat = "2006-01-02T15:04:05Z07:00"

// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       mpfd
// @Produce      mpfd
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	energy	formData	string	true	"energy value need to update"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /energy [post]
func ValidateEnergyData(ctx *gin.Context) {
	userId := ctx.PostForm("userId")
	sessionToken := ctx.PostForm("sessionToken")
	newEnergy, err := strconv.Atoi(ctx.PostForm("energy"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "text/plain", "energy wrong format")
		return
	}
	apiClient := snapser.NewAPIClient(snapser.NewConfiguration())
	res, _response, _ := apiClient.ProfilesServiceApi.GetProfile(context.Background(), userId).Token(sessionToken).Execute()
	if _response.StatusCode != http.StatusNotFound && _response.StatusCode != http.StatusOK {
		ctn, _ := io.ReadAll(_response.Body)
		ctx.String(http.StatusBadRequest, string(ctn))
		return
	}
	if _response.StatusCode == http.StatusOK {

		if res.Profile[energy_key] != nil {
			energy := energyData{
				energy: int(res.Profile[energy_key].(float64)),
				time:   res.Profile[energy_update_key].(string),
			}
			oldTime, _err := time.Parse(time_fomat, energy.time)
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
	profile[energy_key] = newEnergy
	profile[energy_update_key] = time.Now().Format(time_fomat)
	print("current time %s", time.Now().Format(time_fomat))
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

// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       mpfd
// @Produce      mpfd
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	type	formData	string	true	"type of currency: SORT_CURRENCY or PREMIUM_CURRENCY"
// @Param	value	formData	string	true	"value need to update"
// @Param	goldSource	formData	string	true	"source of gold"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       //update-currency [post]
func ValidateCurrency(ctx *gin.Context) {
	userId := ctx.PostForm("userId")
	sessionToken := ctx.PostForm("sessionToken")
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
