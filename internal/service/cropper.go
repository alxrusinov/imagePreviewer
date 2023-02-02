package service

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/alxrusinov/imagePreviewer/internal/repository"
	"github.com/disintegration/imaging"
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
		return nil, DecodeImageError
	}

	croppedImg := imaging.Fill(originalImg, params.Width, params.Height, imaging.Center, imaging.Lanczos)

	buf := new(bytes.Buffer)

	err = jpeg.Encode(buf, croppedImg, nil)

	if err != nil {
		return nil, EncodeImageError
	}

	return buf.Bytes(), nil
}

func (crp *CropperService) GetByURL(rawURL repository.Key) ([]byte, bool) {
	result, exist := crp.repo.Get(rawURL)

	if !exist {
		return nil, exist
	}

	if img, ok := result.([]byte); ok {
		return img, ok
	}

	return nil, exist
}

func (crp *CropperService) SaveToCache(key repository.Key, value []byte) bool {
	ok := crp.repo.Set(key, value)

	return ok
}

type CropperParams struct {
	Address string
	Width   int
	Height  int
}

func NewCropperParams(address string, width int, height int) *CropperParams {
	return &CropperParams{Address: address, Width: width, Height: height}
}
