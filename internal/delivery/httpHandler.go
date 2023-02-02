package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alxrusinov/imagePreviewer/internal/client"
	"github.com/alxrusinov/imagePreviewer/internal/repository"
	"github.com/alxrusinov/imagePreviewer/internal/service"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Services *service.Services
	Client   *client.Client
}

func NewHTTPHandler(services *service.Services, client *client.Client) *HTTPHandler {
	return &HTTPHandler{Services: services, Client: client}
}

func (handler *HTTPHandler) FillHandler(ctx *gin.Context) {
	width := ctx.Param("width")
	height := ctx.Param("height")
	link := ctx.Param("link")

	widthParam, err := strconv.Atoi(width)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrBadParams.Error()})
		return
	}

	heightParam, err := strconv.Atoi(height)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrBadParams.Error()})
		return
	}

	imageAddress, err := createImageAddress(link)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrBadParsedAddress.Error()})
		return
	}

	header := ctx.Request.Header.Clone()

	img, err := handler.Client.GetWithHeaders(imageAddress, header)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("%s: %w", ErrReadImage.Error(), err).Error()})
		return
	}

	rawURL := ctx.Request.URL.String()

	result, ok := handler.Services.CropperService.GetByURL(repository.Key(rawURL))

	if ok {
		fileName := createFileName(link, widthParam, heightParam)

		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
		ctx.Data(http.StatusOK, "image/jpg", result)
		return
	}

	cropperParams := service.NewCropperParams(link, widthParam, heightParam)

	cropped, err := handler.Services.CropperService.Fill(img, cropperParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": ErrImageProcessing.Error()})
		return
	}

	_ = handler.Services.CropperService.SaveToCache(repository.Key(rawURL), cropped)

	fileName := createFileName(link, widthParam, heightParam)

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	ctx.Data(http.StatusOK, "image/jpg", cropped)
}
