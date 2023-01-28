package service

import (
	"bytes"
	"github.com/alxrusinov/imagePreviewer/internal/repository"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
)

type CropperService struct {
	repo repository.Repo
}

func NewCropperService(repo repository.Repo) *CropperService {
	return &CropperService{repo: repo}
}

func (crp *CropperService) Fill(img []byte, params *CropperParams) ([]byte, error) {

	imgReader := bytes.NewReader(img)

	originalImg, _, err := image.Decode(imgReader)

	if err != nil {
		return nil, err
	}

	croppedImg := imaging.Fill(originalImg, params.Width, params.Height, imaging.Center, imaging.Lanczos)

	buf := new(bytes.Buffer)

	err = jpeg.Encode(buf, croppedImg, nil)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type CropperParams struct {
	Address string
	Width   int
	Height  int
}

func NewCropperParams(address string, width int, height int) *CropperParams {
	return &CropperParams{Address: address, Width: width, Height: height}
}
