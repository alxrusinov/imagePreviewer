package delivery

import (
	"fmt"
	"github.com/alxrusinov/imagePreviewer/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	services *service.Services
}

func NewHttpHandler(services *service.Services) *HttpHandler {
	return &HttpHandler{services: services}
}

func (handler *HttpHandler) FillHandler(ctx *gin.Context) {

	width := ctx.Param("width")
	height := ctx.Param("height")
	link := ctx.Param("link")

	imageAddress, err := createImageAddress(link)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BadParsedAddressError.Error()})
	}

	fmt.Printf("imagePath is %v\n", imageAddress)

	ctx.JSON(http.StatusOK, gin.H{
		"width":  width,
		"height": height,
		"link":   link,
	})
}
