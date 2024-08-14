package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snapser/cloudecode/configs"
	"snapser/cloudecode/internal_connector"
	"snapser/cloudecode/response"
	proto2 "snapser/cloudecode/snapserpb/profiles"
	proto "snapser/cloudecode/snapserpb/storage"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	structpb2 "google.golang.org/protobuf/types/known/structpb"
)

type Router struct {
	Route *internal_connector.SnapserServiceConnector
}
type PlayerProfile struct {
	DisplayName string
	CreateTime  string
	Region      string
}
type CheckStatus struct {
	Day   int `json:"day"`
	State int `json:"state"`
}
type NewbieCheckStatus struct {
	Day   string `json:"day"`
	State int    `json:"state"`
}

const (
	Claimed   int = 0
	UnClaimed     = 1
	Claim         = 2
	NotComing     = 3
)

func GetStorageData(userId string, c context.Context, pgs *internal_connector.SnapserServiceConnector, blobName string, isAppendBlob bool, accessType string) (string, error) {
	if isAppendBlob {
		blob, err := pgs.StorageClient.GetBlob(c, &proto.GetBlobRequest{OwnerId: userId, AccessType: accessType, Key: blobName})
		if err != nil {
			return "", err
		}
		return blob.GetValue(), nil

	} else {
		blob, err := pgs.StorageClient.GetBlob(c, &proto.GetBlobRequest{OwnerId: userId, AccessType: accessType, Key: blobName})
		if err != nil {
			return "", err
		}
		return blob.GetValue(), nil

	}
}
func GetDailyLoginData(userId string, c context.Context, pgs *internal_connector.SnapserServiceConnector) ([]NewbieCheckStatus, error) {
	blob, err := pgs.StorageClient.GetAppendBlob(c, &proto.GetAppendBlobRequest{OwnerId: userId, AccessType: "private", Key: configs.LoginDataBlobName})
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	var loginDay []string
	var checkDay []time.Time
	today := time.Now()
	if blob != nil {
		loginDay = strings.Split(blob.GetValue(), ";")
		loginDay = loginDay[:len(loginDay)-1]
		// get day login in this month
		for _, value := range loginDay {
			parse, err := time.Parse(configs.TimeFormatDayOnly, value)
			if err == nil && parse.Year() == today.Year() && parse.Month() == today.Month() {
				checkDay = append(checkDay, parse)
			}
		}
	}
	// check from 1st day of this month to today
	var checkInProgress []NewbieCheckStatus
	for i := today.Day() - 1; i >= 0; i-- {
		var checkin_day = today.AddDate(0, 0, -i)
		checkInProgress = append(checkInProgress, NewbieCheckStatus{Day: checkin_day.Format(configs.TimeFormatDayOnly), State: UnClaimed})
	}
	fmt.Println("Login at day: " + strconv.Itoa(len(checkDay)))
	for i := 0; i < len(checkDay); i++ {
		checkInProgress[checkDay[i].Day()-1].State = Claimed
	}
	return checkInProgress, nil
}
func GetNewbieLoginData(userId string, c context.Context, pgs *internal_connector.SnapserServiceConnector) ([]NewbieCheckStatus, error) {
	blob, err := pgs.StorageClient.GetAppendBlob(c, &proto.GetAppendBlobRequest{OwnerId: userId, AccessType: "private", Key: configs.NewBieDataBlobName})
	if err != nil {
		fmt.Println("error " + err.Error())
	}
	profile, err := GetUserProfile(userId, c, pgs)
	if err != nil {
		return nil, err
	}
	createDay := profile.GetFields()[configs.TimeCreateKey]
	createDayStr := createDay.GetStringValue()
	fmt.Println("Create user from:" + createDayStr)
	createDayTime, _ := time.Parse(configs.TimeFormatTimeDay, createDayStr)
	today := time.Now()
	firstDay := time.Date(createDayTime.Year(), createDayTime.Month(), createDayTime.Day(), 0, 0, 0, 0, time.UTC)
	secondDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	var loginDay []string
	var checkDay []time.Time

	if blob != nil {
		loginDay = strings.Split(blob.GetValue(), ";")
		loginDay = loginDay[:len(loginDay)-1]
		for _, value := range loginDay {
			parse, err := time.Parse(configs.TimeFormatDayOnly, value)
			if err == nil {
				checkDay = append(checkDay, parse)
			}
		}
	}
	var day = int(secondDay.Sub(firstDay).Hours()) / 24
	fmt.Println("Time from create: " + strconv.Itoa(int(secondDay.Sub(firstDay).Hours())))
	fmt.Println("Create player for " + strconv.Itoa(day) + "day")
	var checkInProgress []NewbieCheckStatus
	if day > 9 {
		return checkInProgress, nil
	}
	for i := 0; i <= day; i++ {
		var check_Day = createDayTime.AddDate(0, 0, i)
		checkInProgress = append(checkInProgress, NewbieCheckStatus{Day: check_Day.Format(configs.TimeFormatDayOnly), State: UnClaimed})
	}
	for _, vDay := range checkDay {
		if len(checkDay) > 0 {
			var index = int(vDay.Sub(firstDay).Hours() / 24)
			if index < len(checkInProgress) {
				checkInProgress[index].State = Claimed
			}
		}
	}
	return checkInProgress, nil
}

func GetUserProfile(userId string, c context.Context, pgs *internal_connector.SnapserServiceConnector) (*structpb.Struct, error) {
	data, err := pgs.ProfileClient.GetUserData(c, &proto2.GetUserDataRequest{UserId: userId})
	if err != nil {
		fmt.Println("not found profile with error:  " + err.Error())
		return nil, err
	}
	return data.Profile, nil
}

func SetupNewUser(userId string, c context.Context, pgs *internal_connector.SnapserServiceConnector) error {
	// Update user profile
	createTime := time.Now().Format(configs.TimeFormatTimeDay)
	var profile structpb.Struct
	profile.Fields = make(map[string]*structpb2.Value)
	fmt.Println("create user profile")
	profile.Fields[configs.TimeCreateKey] = &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: createTime}}
	_, err := pgs.ProfileClient.UpsertProfile(c, &proto2.UpsertProfileRequest{UserId: userId, Profile: &profile})
	if err != nil {
		return err
	}
	return nil
}

// @Summary      Level up equipment
// @Description  Level up equipment
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	matchId	formData	string	true	"id of the pvp game"
// @Param	opponentID	formData	string	true	"id of opponent"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/account-info[get]
func (pgs *Router) GetAccountInfo(ctx *gin.Context) {
	userId, _ := ctx.GetQuery(configs.UserId)
	if userId == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	profile, err := GetUserProfile(userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	if profile == nil || profile.Fields == nil {
		err := SetupNewUser(userId, c, pgs.Route)
		if err != nil {
			ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusNoContent, Message: "Create new user but error: " + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusNoContent, Message: "Create new user "})
		return
	} else {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusOK, Message: "New user"})
		return
	}
}

// @Summary      Level up equipment
// @Description  Level up equipment
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	matchId	formData	string	true	"id of the pvp game"
// @Param	opponentID	formData	string	true	"id of opponent"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/account-info[get]
func (pgs *Router) GetAccountInfo2(ctx *gin.Context) {
	userId, _ := ctx.GetQuery(configs.UserId)
	if userId == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	profile, err := GetUserProfile(userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	isNewUser := false
	if profile == nil || profile.Fields == nil {
		err := SetupNewUser(userId, c, pgs.Route)
		if err != nil {
			ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusNoContent, Message: "Create new user but error: " + err.Error()})
			return
		}
		isNewUser = true
	}

	dataMap := make(map[string]string)
	profileData, err := GetStorageData(userId, c, pgs.Route, configs.ProfileDataBlob, false, "protected")
	if err == nil {
		dataMap[configs.ProfileDataBlob] = profileData
	}
	championData, err := GetStorageData(userId, c, pgs.Route, configs.ChampionDataBlob, false, "private")
	if err == nil {
		dataMap[configs.ChampionDataBlob] = championData
	}
	equipmentData, err := GetStorageData(userId, c, pgs.Route, configs.EquipmentDataBlob, false, "private")
	if err == nil {
		dataMap[configs.EquipmentDataBlob] = equipmentData
	}
	cardSkillData, err := GetStorageData(userId, c, pgs.Route, configs.CardSkillDataBlob, false, "private")
	if err == nil {
		dataMap[configs.CardSkillDataBlob] = cardSkillData
	}
	rankData, err := GetStorageData(userId, c, pgs.Route, configs.RankDataBlob, false, "private")
	if err == nil {
		dataMap[configs.RankDataBlob] = rankData
	}
	teamData, err := GetStorageData(userId, c, pgs.Route, configs.UserTeamsBlob, false, "private")
	if err == nil {
		dataMap[configs.UserTeamsBlob] = teamData
	}
	campaignData, err := GetStorageData(userId, c, pgs.Route, configs.UserCampaignBlob, false, "private")
	if err == nil {
		dataMap[configs.UserCampaignBlob] = campaignData
	}
	questData, err := GetStorageData(userId, c, pgs.Route, configs.UserQuestBlob, false, "private")
	if err == nil {
		dataMap[configs.UserQuestBlob] = questData
	}
	pvpData, err := GetStorageData(userId, c, pgs.Route, configs.UserPvp, false, "private")
	if err == nil {
		dataMap[configs.UserPvp] = pvpData
	}
	lootBoxData, err := GetStorageData(userId, c, pgs.Route, configs.UserLootBox, false, "private")
	if err == nil {
		dataMap[configs.UserLootBox] = lootBoxData
	}
	itemPurchaseData, err := GetStorageData(userId, c, pgs.Route, configs.ItemPurchaseData, false, "protected")
	if err == nil {
		dataMap[configs.ItemPurchaseData] = itemPurchaseData
	}
	generalUtilsData, err := GetStorageData(userId, c, pgs.Route, configs.GeneralUtilsData, false, "private")
	if err == nil {
		dataMap[configs.GeneralUtilsData] = generalUtilsData
	}
	dealDailyShopData, err := GetStorageData(userId, c, pgs.Route, configs.DealDailyShopData, false, "private")
	if err == nil {
		dataMap[configs.DealDailyShopData] = dealDailyShopData
	}
	offerShop, err := GetStorageData(userId, c, pgs.Route, configs.OfferShopData, false, "private")
	if err == nil {
		dataMap[configs.OfferShopData] = offerShop
	}
	userQuestData, err := GetStorageData(userId, c, pgs.Route, configs.QuestData, false, "private")
	if err == nil {
		dataMap[configs.QuestData] = userQuestData
	}
	tutorialRecord, err := GetStorageData(userId, c, pgs.Route, configs.TutorialRecordBlob, false, "private")
	if err == nil {
		dataMap[configs.TutorialRecordBlob] = tutorialRecord
	}
	gameItems, err := GetStorageData(userId, c, pgs.Route, configs.GameItems, false, "private")
	if err == nil {
		dataMap[configs.GameItems] = gameItems
	}
	loginData, err := GetStorageData(userId, c, pgs.Route, configs.LoginData, false, "public")
	if err == nil {
		dataMap[configs.LoginData] = loginData
	}
	NewBieData, err := GetNewbieLoginData(userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: "Create new user but error: " + err.Error()})
		return
	}
	DailyLoginData, err := GetDailyLoginData(userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: "Create new user but error: " + err.Error()})
		return
	}
	newbieDataString, err := json.Marshal(NewBieData)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	dailyDataString, err := json.Marshal(DailyLoginData)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	dataMap[configs.LoginDataBlobName] = string(dailyDataString)
	dataMap[configs.NewBieDataBlobName] = string(newbieDataString)
	data, _ := json.Marshal(dataMap)
	if isNewUser {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusNoContent, Message: "Create new user ", Data: string(data)})
		return
	} else {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusOK, Message: "Create new user ", Data: string(data)})
		return
	}
}

func (pgs *Router) GetDailyRewardCheckingStatus(ctx *gin.Context) {
	userId, _ := ctx.GetQuery(configs.UserId)
	if userId == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	blob, err := pgs.Route.StorageClient.GetAppendBlob(c, &proto.GetAppendBlobRequest{OwnerId: userId, AccessType: "private", Key: configs.LoginDataBlobName})
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	var loginDay []string
	var checkDay []time.Time
	today := time.Now()
	if blob != nil {
		loginDay = strings.Split(blob.GetValue(), ";")
		loginDay = loginDay[:len(loginDay)-1]
		// get day login in this month
		for _, value := range loginDay {
			parse, err := time.Parse(configs.TimeFormatDayOnly, value)
			if err == nil && parse.Year() == today.Year() && parse.Month() == today.Month() {
				checkDay = append(checkDay, parse)
			}
		}
	}
	// check from 1st day of this month to today
	var checkInProgress []CheckStatus
	for i := 1; i <= today.Day(); i++ {
		if len(checkDay) > 0 {
			// if login this day, add status
			if i == checkDay[0].Day() {
				checkInProgress = append(checkInProgress, CheckStatus{Day: i, State: Claimed})
				checkDay = checkDay[1:]
			} else {
				checkInProgress = append(checkInProgress, CheckStatus{Day: i, State: UnClaimed})
			}
		} else {
			checkInProgress = append(checkInProgress, CheckStatus{Day: i, State: UnClaimed})
		}
	}
	value, err := json.Marshal(checkInProgress)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusOK, Message: string(value)})
	return
}

func (pgs *Router) GetNewbieRewardCheckingStatus(ctx *gin.Context) {
	userId, _ := ctx.GetQuery(configs.UserId)
	if userId == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	blob, err := pgs.Route.StorageClient.GetAppendBlob(c, &proto.GetAppendBlobRequest{OwnerId: userId, AccessType: "private", Key: configs.NewBieDataBlobName})
	if err != nil {
		fmt.Println("error " + err.Error())
	}
	profile, err := GetUserProfile(userId, c, pgs.Route)
	if err != nil {
		return
	}
	createDay := profile.GetFields()[configs.TimeCreateKey]
	createDayStr := createDay.GetStringValue()
	fmt.Println("Create user from:" + createDayStr)
	createDayTime, _ := time.Parse(configs.TimeFormatTimeDay, createDayStr)
	today := time.Now()
	firstDay := time.Date(createDayTime.Year(), createDayTime.Month(), createDayTime.Day(), 0, 0, 0, 0, time.UTC)
	secondDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	var loginDay []string
	var checkDay []time.Time

	if blob != nil {
		loginDay = strings.Split(blob.GetValue(), ";")
		loginDay = loginDay[:len(loginDay)-1]
		for _, value := range loginDay {
			parse, err := time.Parse(configs.TimeFormatDayOnly, value)
			if err == nil && parse.Year() == today.Year() && parse.Month() == today.Month() {
				checkDay = append(checkDay, parse)
			}
		}
	}
	var day = int(secondDay.Sub(firstDay).Hours()) / 24
	fmt.Println("Time from create: " + strconv.Itoa(int(secondDay.Sub(firstDay).Hours())))
	fmt.Println("Create player for " + strconv.Itoa(day) + "day")
	var checkInProgress []CheckStatus
	for i := day; i >= 0; i-- {
		vDay := today.AddDate(0, 0, -i)
		if len(checkDay) > 0 {
			if vDay.Day() == checkDay[0].Day() && vDay.Month() == checkDay[0].Month() {
				checkInProgress = append(checkInProgress, CheckStatus{Day: day - i + 1, State: Claimed})
				checkDay = checkDay[1:]
			} else {
				checkInProgress = append(checkInProgress, CheckStatus{Day: day - i + 1, State: UnClaimed})
			}
		} else {
			checkInProgress = append(checkInProgress, CheckStatus{Day: day - i + 1, State: UnClaimed})
		}
	}
	value, err := json.Marshal(checkInProgress)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusOK, Message: string(value)})
	return
}
func (pgs *Router) CheckinDaily(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	blob, _ := pgs.Route.StorageClient.GetAppendBlob(c, &proto.GetAppendBlobRequest{OwnerId: userId, AccessType: "private", Key: configs.LoginDataBlobName})
	today := time.Now()
	isLogin := false
	if blob != nil {
		fmt.Println("data login: " + blob.GetValue())
		var loginDay []string
		loginDay = strings.Split(blob.GetValue(), ";")
		fmt.Println("select day:" + loginDay[len(loginDay)-2])
		if len(loginDay) > 1 {
			lastLoginDay, _ := time.Parse(configs.TimeFormatDayOnly, loginDay[len(loginDay)-2])
			if lastLoginDay.Year() == today.Year() && lastLoginDay.Month() == today.Month() && lastLoginDay.Day() == today.Day() {
				isLogin = true
			}
		}

	}
	if isLogin {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusBadRequest, Message: "you are already checkin"})
		return
	} else {
		_, err := pgs.Route.StorageClient.UpdateAppendBlob(c, &proto.UpdateAppendBlobRequest{Key: configs.LoginDataBlobName, OwnerId: userId, AccessType: "private", Value: today.Format(configs.TimeFormatDayOnly) + ";"})
		if err != nil {
			ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: "Checkin fail with error:" + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusOK, Message: "Checkin success"})
		return
	}
}
func (pgs *Router) CheckinNewbie(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(context.Background(), "gateway", "internal")
	blob, _ := pgs.Route.StorageClient.GetAppendBlob(c, &proto.GetAppendBlobRequest{OwnerId: userId, AccessType: "private", Key: configs.NewBieDataBlobName})

	profile, err := GetUserProfile(userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	createDay := profile.Fields[configs.TimeCreateKey]
	today := time.Now()
	createDayStr := createDay.GetStringValue()
	if createDayStr != "" {
		createDayTime, _ := time.Parse(configs.TimeFormatTimeDay, createDayStr)
		firstDay := time.Date(createDayTime.Year(), createDayTime.Month(), createDayTime.Day(), 0, 0, 0, 0, time.UTC)
		secondDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
		var day = secondDay.Sub(firstDay).Hours() / 24
		if day > 7 {
			ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusBadRequest, Message: "Your newbie packet is expired"})
			return
		}
	}
	isLogin := false

	var loginDay []string
	if blob != nil {
		fmt.Println("data login: " + blob.GetValue())
		loginDay = strings.Split(blob.GetValue(), ";")
		if len(loginDay) > 1 {
			fmt.Println("select day:" + loginDay[len(loginDay)-2])
			lastLoginDay, _ := time.Parse(configs.TimeFormatDayOnly, loginDay[len(loginDay)-2])
			if lastLoginDay.Year() == today.Year() && lastLoginDay.Month() == today.Month() && lastLoginDay.Day() == today.Day() {
				isLogin = true
			}
		}
	}
	if isLogin {
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusBadRequest, Message: "you are already checkin"})
		return
	} else {
		_, err := pgs.Route.StorageClient.UpdateAppendBlob(c, &proto.UpdateAppendBlobRequest{Key: configs.NewBieDataBlobName, OwnerId: userId, AccessType: "private", Value: today.Format(configs.TimeFormatDayOnly) + ";"})
		if err != nil {
			ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusInternalServerError, Message: "Checkin fail with error:" + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, response.ResMessage{Code: http.StatusOK, Message: "Checkin success"})
		return
	}
}
func (pgs *Router) ReclaimNewbie(ctx *gin.Context) {

}
func (pgs *Router) ReclaimDailyReward(ctx *gin.Context) {

}
