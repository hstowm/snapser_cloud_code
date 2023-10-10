package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"snapser/cloudecode/docs"
	economy "snapser/cloudecode/economy"

	"github.com/gin-gonic/gin"
)

// @title           Skybull Snapser CloudeCode
// @version         0.0.1
// @description     Skybull snapser cloud code

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v2/byosnap-skybull")
	{
		v1.POST("/level-up-champion", economy.LevelUpChampionFee)
		v1.POST("/level-up-equipment", economy.LevelUpEquipmentFee)
		v1.POST("/sell-equipment", economy.SellEquipment)
		v1.POST("/re-roll-equipment-modifier", economy.ReRollEquipmentModifier)
		v1.POST("/sell-champion", economy.SellEquipment)
		v1.POST("/win-pvp", economy.WinPvPGame)
		v1.POST("/win-campaign", economy.WinCampaign)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
