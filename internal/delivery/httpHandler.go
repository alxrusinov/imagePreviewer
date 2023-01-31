package delivery

import (
	"fmt"
	"github.com/alxrusinov/imagePreviewer/internal/client"
	"github.com/alxrusinov/imagePreviewer/internal/repository"
	"github.com/alxrusinov/imagePreviewer/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HttpHandler struct {
	Services *service.Services
	Client   *client.Client
}

func NewHttpHandler(services *service.Services, client *client.Client) *HttpHandler {
	return &HttpHandler{Services: services, Client: client}
}

func (handler *HttpHandler) FillHandler(ctx *gin.Context) {
	width := ctx.Param("width")
	height := ctx.Param("height")
	link := ctx.Param("link")

	widthParam, err := strconv.Atoi(width)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": BadParamsError.Error()})
		return
	}

	heightParam, err := strconv.Atoi(height)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": BadParamsError.Error()})
		return
	}

	imageAddress, err := createImageAddress(link)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": BadParsedAddressError.Error()})
		return
	}

	header := ctx.Request.Header.Clone()

	img, err := handler.Client.GetWithHeaders(imageAddress, header)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": ReadImageError.Error()})
		return
	}

	rawUrl := ctx.Request.URL.String()

	result, ok := handler.Services.CropperService.GetByUrl(repository.Key(rawUrl))

	if ok {
		fileName := createFileName(link, widthParam, heightParam)

		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
		ctx.Data(http.StatusOK, "image/jpg", result)
		return
	}

	cropperParams := service.NewCropperParams(link, widthParam, heightParam)

	cropped, err := handler.Services.CropperService.Fill(img, cropperParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": ImageProcessingError.Error()})
		return
	}

	_ = handler.Services.CropperService.SaveToCache(repository.Key(rawUrl), cropped)

	fileName := createFileName(link, widthParam, heightParam)

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	ctx.Data(http.StatusOK, "image/jpg", cropped)
}
