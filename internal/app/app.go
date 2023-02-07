package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alxrusinov/imagePreviewer/internal/client"
	"github.com/alxrusinov/imagePreviewer/internal/delivery"
	"github.com/alxrusinov/imagePreviewer/internal/repository/lru"
	"github.com/alxrusinov/imagePreviewer/internal/routes"
	"github.com/alxrusinov/imagePreviewer/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	DefaultCacheSize = 10
	DefaultPort      = "80"
)

func Run() {
	port := DefaultPort
	cacheSize := DefaultCacheSize

	if portEnv, exist := os.LookupEnv("PORT"); exist {
		port = portEnv
	}

	if cacheEnv, exist := os.LookupEnv("CACHE"); exist {
		if size, err := strconv.Atoi(cacheEnv); err == nil {
			cacheSize = size
		}
	}

	httpClient := client.NewClient()

	repo := lru.NewCache(cacheSize)
	cropperService := service.NewCropperService(repo)
	services := service.NewServices(cropperService)
	handler := delivery.NewHTTPHandler(services, httpClient)

	router := gin.Default()

	api := router.Group("fill")

	api.GET(routes.FILL, handler.FillHandler)

	addr := fmt.Sprintf("0.0.0.0:%s", port)

	if err := router.Run(addr); err != nil {
		log.Fatalln(err)
	}
}
