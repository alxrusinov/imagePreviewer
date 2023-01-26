package service

type Services struct {
	CropperService CropperService
}

func NewServices(cropperService *CropperService) *Services {
	return &Services{CropperService: *cropperService}
}
