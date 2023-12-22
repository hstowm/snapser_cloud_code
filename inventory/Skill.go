package inventory

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"net/http"
	"snapser/cloudecode/configs"
	"snapser/cloudecode/response"
	proto "snapser/cloudecode/snapserpb/storage"
	"strings"
)

type SkillData struct {
	UserCardUID string `json:"userCardUID"`
	CardUid     string `json:"cardUID"`
}

func (pgs *Router) GetPlayerAllCardSkill(userID string, c context.Context) ([]SkillData, error) {
	blob, err := pgs.Route.StorageClient.GetBlob(c, &proto.GetBlobRequest{AccessType: "private", Key: configs.CardSkillDataBlob, OwnerId: userID})
	if err != nil {
		return nil, errors.New("error when call api to snapser server " + err.Error())
	}
	var b map[string][]SkillData
	jsonData := blob.GetValue()
	err = json.Unmarshal([]byte(jsonData), &b)
	if err != nil {
		return nil, errors.New("can't parse data from json to list equipment")
	}
	var cardData = b["data"]
	return cardData, err
}

// @Summary      Recycle champion
// @Description  Recycle champion
// @Accept       mpfd
// @Produce      mpfd
// @Param token header string true "Authorization Token" Format(uuid)
// @Param	userId	formData	string	true	"userId when they login"
// @Param	sessionToken	formData	string	true	"session token"
// @Param	cardSkill	formData	string	true	"list card want to recycle"
// @Success      200  {array}  	string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/byosnap-skybull-cloudcode/recycle-card-skill [post]
func (pgs *Router) RecycleCardSkill(ctx *gin.Context) {
	userId := ctx.PostForm(configs.UserId)
	sessionToken := ctx.PostForm(configs.SessionToken)
	cardSkills := ctx.PostForm("cardSkill")
	if userId == "" || sessionToken == "" || cardSkills == "" {
		ctx.String(http.StatusBadRequest, "data is not enough")
		return
	}
	userCardSkillsRecycle := strings.Split(cardSkills, ",")
	if len(userCardSkillsRecycle) < configs.NumCardSkillToRecycle {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: "Not enough number card"})
		return
	}
	c := metadata.AppendToOutgoingContext(ctx.Request.Context(), "gateway", "internal")
	usercardSkills, err := pgs.GetPlayerAllCardSkill(userId, c)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	if len(usercardSkills) < configs.NumCardSkillToRecycle {
		ctx.JSON(http.StatusNotFound, response.ResMessage{Code: http.StatusNotFound, Message: "Player don't have enough cardSkill"})
	}
	for _, cardSkillID := range userCardSkillsRecycle {
		index := -1
		for i, cardSkill := range usercardSkills {
			if cardSkill.UserCardUID == cardSkillID {
				index = i
			}
		}
		if index >= 0 {
			copy(usercardSkills[index:], usercardSkills[index+1:])
			usercardSkills[len(usercardSkills)-1] = usercardSkills[0] // or the zero value of T
			usercardSkills = usercardSkills[:len(usercardSkills)-1]
		} else {
			ctx.JSON(http.StatusOK, response.ResMessage{
				Code:    http.StatusBadRequest,
				Message: "Player don't own cardSkill with id: " + cardSkillID,
			})
		}
	}
	ctx.JSON(http.StatusOK, response.ResMessage{
		Code:    http.StatusOK,
		Message: "can merger",
	})
	return
}
