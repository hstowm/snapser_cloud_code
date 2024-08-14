package storage

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"math"
	"net/http"
	"snapser/cloudecode/configs"
	"snapser/cloudecode/economy"
	"snapser/cloudecode/internal_connector"
	"snapser/cloudecode/response"
	proto "snapser/cloudecode/snapserpb/profiles"
	proto2 "snapser/cloudecode/snapserpb/storage"
	"strconv"
	"time"
)

type PVPGame struct {
	MatchID     string
	Host        string
	Client      string
	Winner      string
	TimeCreate  time.Time
	UpdateTIme  time.Time
	GameType    PvPGameType
	HostPower   float64
	ClientPower float64
}
type PvPGameType int

const (
	PlayerGame PvPGameType = 0
	BotGame    PvPGameType = 1
)

type Router struct {
	Route            *internal_connector.SnapserServiceConnector
	ListPvPGame      map[string]PVPGame // list game in server
	ListCampaignGame map[string]CampaignGame
}
type Session struct {
	SessionID  string `json:"sessionID"`
	TotalWin   int    `json:"totalWin"`
	TotalLose  int    `json:"totalLose"`
	Rank       int    `json:"rank"`
	Elo        int    `json:"elo"`
	TimeUpdate string `json:"lastPlayedTime"`
}
type CampaignGame struct {
	MatchId    string `json:"matchID"`
	TimeCreate string `json:"timeCreate"`
	Star       int    `json:"star"`
	Reward     string `json:"reward"`
}

func (session *Session) UpdateElo(elo int) {

	session.Elo = int(math.Max(0, float64(session.Elo+elo)))
}
func CalculateELo(winnerElo int, loserElo int) int {
	var result int
	result = int(configs.ELoK * 1 / (1 + math.Pow(10, math.Abs(float64(winnerElo-loserElo))/400)))
	return result
}
func CalculateCampaignFee(power int, star int) map[string]int32 {
	fees := make(map[string]int32)
	fees[configs.Gold] = int32(power * (star + 1))
	return fees
}
func CalculatePvPFee(enemyElo int) map[string]int32 {
	fees := make(map[string]int32)
	fees[configs.Gold] = int32(100 + enemyElo*2)
	return fees
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
// @Router       /v1/byosnap-skybull-cloudcode/create-pvp-game [post]
func (pgs *Router) CreatePvpGame(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	opponentId := ctx.PostForm("opponentID")
	sessionToken := ctx.PostForm(configs.SessionToken)
	matchID := ctx.PostForm("matchID")
	if userId == "" || sessionToken == "" || opponentId == "" || matchID == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}

	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	newGame := PVPGame{
		MatchID:    matchID,
		Client:     opponentId,
		Host:       userId,
		TimeCreate: time.Now(),
		GameType:   PlayerGame,
	}
	fmt.Println("Create match between player: " + newGame.Client + " and " + newGame.Host)
	hostProfile, err := pgs.Route.ProfileClient.GetProfile(c, &proto.GetProfileRequest{
		UserId: userId,
	})
	if err != nil {
		newGame.HostPower = 0
	} else {
		hostPower, err := strconv.ParseFloat(hostProfile.GetProfile().GetFields()[configs.PowerInProfiles].String(), 64)
		if err != nil {
			newGame.HostPower = 100
		} else {
			newGame.HostPower = hostPower
		}
	}

	if opponentId == configs.BotInPvPID {
		newGame.GameType = BotGame
		newGame.ClientPower = newGame.HostPower
	} else {
		clientProfile, err := pgs.Route.ProfileClient.GetProfile(c, &proto.GetProfileRequest{
			UserId: opponentId,
		})
		if err != nil {
			newGame.ClientPower = 0
		} else {
			clientPower, err := strconv.ParseFloat(clientProfile.GetProfile().GetFields()[configs.PowerInProfiles].String(), 64)
			if err != nil {
				newGame.ClientPower = 100
			} else {
				newGame.ClientPower = clientPower
			}
		}
	}
	// Sub energy in two player
	_, err = economy.UpdateEnergy(newGame.Host, -configs.EnergyUseWhenPlay, false, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResMessage{
			Code:       http.StatusInternalServerError,
			Message:    err.Error(),
			Currencies: nil,
		})
		return
	}
	if newGame.GameType == PlayerGame {
		_, err = economy.UpdateEnergy(newGame.Client, -configs.EnergyUseWhenPlay, false, c, pgs.Route)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ResMessage{
				Code:       http.StatusInternalServerError,
				Message:    err.Error(),
				Currencies: nil,
			})
		}
	}
	fmt.Println("Create match between player: " + newGame.Client + " and " + newGame.Client)

	if pgs.ListPvPGame == nil {
		pgs.ListPvPGame = make(map[string]PVPGame)
	}
	pgs.ListPvPGame[matchID] = newGame
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "Create PvP Game success",
		Currencies: nil,
	})
}

// @Summary      Call when start a campaign game
// @Description  Create campaign game and subject energy
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/create-campaign-game [post]
func (pgs *Router) CreateCampaignGame(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	if userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	// Sub energy in two player
	currencies, err := economy.UpdateEnergy(userId, -configs.EnergyUseWhenPlay, false, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code:       http.StatusInternalServerError,
			Message:    err.Error(),
			Currencies: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "Create campaign Game success",
		Currencies: currencies,
	})
}

// @Summary      Call when player win a pvp game
// @Description  Call when player win a pvp game
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	winnerID	formData	string	true	"opponent snapser id "
// @Param	matchID	formData	string	true	"opponent snapser id "
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/win-pvp [post]
func (pgs *Router) WinPvPGame(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	winner := ctx.PostForm("winnerID")
	matchID := ctx.PostForm("matchID")
	sessionToken := ctx.PostForm(configs.SessionToken)
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")

	if userId == "" || sessionToken == "" || winner == "" || matchID == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	//c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	pvpGame, isExist := pgs.ListPvPGame[matchID]
	if !isExist {
		ctx.String(http.StatusBadRequest, "Not found match with id: "+matchID)
	}

	if pvpGame.Winner != "" {
		currencies, err := economy.GetUserCurrencies(userId, c, pgs.Route)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ResMessage{Code: 200, Message: "Result is updated", Currencies: nil})
		}
		ctx.JSON(http.StatusOK, response.ResMessage{Code: 200, Message: "Result is updated", Currencies: currencies})
		return
	}
	pvpGame.Winner = winner
	pvpGame.UpdateTIme = time.Now()
	pgs.ListPvPGame[matchID] = pvpGame
	var loserPower float64
	loser := ""
	if winner == pvpGame.Host {
		loser = pvpGame.Client
		loserPower = pvpGame.ClientPower
	} else {
		loser = pvpGame.Host
		loserPower = pvpGame.HostPower
	}
	print(loser)
	if pvpGame.GameType == BotGame && winner == configs.BotInPvPID {
		ctx.String(http.StatusOK, "Update game result done")
		return
	}
	currenciesGain := CalculatePvPFee(int(math.Floor(loserPower)))
	currencies, err := economy.UpdateCurrencies(currenciesGain, userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code: http.StatusInternalServerError, Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "update result pvp done, gain money" + strconv.Itoa(int(currencies[configs.Gold])),
		Currencies: currencies,
	})
	str, err := json.Marshal(pvpGame)
	if err != nil {
		return
	}
	_, err = pgs.Route.StorageClient.UpdateAppendBlob(c, &proto2.UpdateAppendBlobRequest{Key: configs.PvPDataBlobHistory, AccessType: "private", OwnerId: pvpGame.Host, Value: string(str)})
	if err != nil {
		return
	} else {
		print("update pvp match data to host's storage done")
	}
	_, err = pgs.Route.StorageClient.UpdateAppendBlob(c, &proto2.UpdateAppendBlobRequest{Key: configs.PvPDataBlobHistory, AccessType: "private", OwnerId: pvpGame.Client, Value: string(str)})
	if err != nil {
		return
	} else {
		print("update pvp match data to client's storage done")
	}
}

// @Summary      Call when win a campaign
// @Description   Call when win a campaign
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login with snapser"
// @Param	sessionToken	formData	string	true	"session token for login version"
// @Param	stagePower	formData	string	true	"power stage player win"
// @Param	star	formData	string	true	"the start they reach"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/win-campaign [post]
func (pgs *Router) WinCampaign(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	stagePower, _ := strconv.Atoi(ctx.PostForm("stagePower"))
	star, _ := strconv.Atoi(ctx.PostForm("star"))

	if userId == "" || sessionToken == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	currenciesGain := CalculateCampaignFee(stagePower, star)
	currencies, err := economy.UpdateCurrencies(currenciesGain, userId, c, pgs.Route)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResMessage{
			Code: http.StatusInternalServerError, Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:       http.StatusOK,
		Message:    "update result campaign done, gain money" + strconv.Itoa(int(currencies[configs.Gold])),
		Currencies: currencies,
	})

	newGame := CampaignGame{MatchId: uuid.New().String(), Star: star, TimeCreate: time.Now().Format(configs.TimeFormatTimeDay), Reward: strconv.Itoa(int(currenciesGain[configs.Gold]))}
	str, err := json.Marshal(newGame)
	if err != nil {
		return
	}
	_, err = pgs.Route.StorageClient.UpdateAppendBlob(c, &proto2.UpdateAppendBlobRequest{Key: configs.CampaignBlobHistory, AccessType: "private", OwnerId: userId, Value: string(str)})
	if err != nil {
		return
	} else {
		print("update campaign match data to storage done")
	}

}
