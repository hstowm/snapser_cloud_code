package main

import (
	"log"
	"snapser/cloudecode/router"
)
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"       // swagger embed files

// @title           Skybull Snapser CloudeCode
// @version         0.0.1
// @description     Skybull snapser cloud code

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	/*	docs.SwaggerInfo.Title = "Swagger Example API"
		docs.SwaggerInfo.Description = "This is a sample server Petstore server."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}*/
	pgs, err := router.New()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	r := pgs.NewRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
