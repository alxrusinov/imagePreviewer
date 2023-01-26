package service

import "github.com/alxrusinov/imagePreviewer/internal/repository"

type CropperService struct {
	repo repository.Repo
}

func NewCropperService(repo repository.Repo) *CropperService {
	return &CropperService{repo: repo}
}

func (crp *CropperService) Fill(img []byte, params *CropperParams) ([]byte, error) {
	return nil, nil
}

type CropperParams struct {
	Address string
	Width   int
	Height  int
}

func NewCropperParams(address string, width int, height int) *CropperParams {
	return &CropperParams{Address: address, Width: width, Height: height}
}
