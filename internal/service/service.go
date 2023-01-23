package service

import "github.com/alxrusinov/imagePreviewer/internal/repository"

type CropperService struct {
	repo repository.Repo
}

func NewCropperService(repo repository.Repo) *CropperService {
	return &CropperService{repo: repo}
}

type Services struct {
	CropperService CropperService
}

func NewServices(cropperService *CropperService) *Services {
	return &Services{CropperService: *cropperService}
}
