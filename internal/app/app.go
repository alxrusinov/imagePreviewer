package app

import (
	"github.com/alxrusinov/imagePreviewer/internal/client"
	"github.com/alxrusinov/imagePreviewer/internal/delivery"
	"github.com/alxrusinov/imagePreviewer/internal/repository/lru"
	"github.com/alxrusinov/imagePreviewer/internal/routes"
	"github.com/alxrusinov/imagePreviewer/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

const DefaultCap = 10

func Run() {
	httpClient := client.NewClient()

	// TODO: Add config later
	repo := lru.NewCache(DefaultCap)
	cropperService := service.NewCropperService(repo)
	services := service.NewServices(cropperService)
	handler := delivery.NewHttpHandler(services, httpClient)

	router := gin.Default()

	api := router.Group("fill")

	api.POST(routes.FILL, handler.FillHandler)

	err := router.Run(":3000")

	if err != nil {
		log.Fatalln(err)
	}
}
