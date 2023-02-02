package app

import (
	"log"

	"github.com/alxrusinov/imagePreviewer/internal/client"
	"github.com/alxrusinov/imagePreviewer/internal/delivery"
	"github.com/alxrusinov/imagePreviewer/internal/repository/lru"
	"github.com/alxrusinov/imagePreviewer/internal/routes"
	"github.com/alxrusinov/imagePreviewer/internal/service"
	"github.com/gin-gonic/gin"
)

const DefaultCap = 10

func Run() {
	httpClient := client.NewClient()

	// TODO: Add config later
	repo := lru.NewCache(DefaultCap)
	cropperService := service.NewCropperService(repo)
	services := service.NewServices(cropperService)
	handler := delivery.NewHTTPHandler(services, httpClient)

	router := gin.Default()

	api := router.Group("fill")

	api.POST(routes.FILL, handler.FillHandler)

	err := router.Run("0.0.0.0:80")
	if err != nil {
		log.Fatalln(err)
	}
}
