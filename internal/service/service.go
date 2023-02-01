package service

var (
	DecodeImageError = NewError("decoding image error")
	EncodeImageError = NewError("encoding processed image error")
)

type Services struct {
	CropperService CropperService
}

func NewServices(cropperService *CropperService) *Services {
	return &Services{CropperService: *cropperService}
}

type Error struct {
	msg string
}

func NewError(msg string) *Error {
	return &Error{msg: msg}
}

func (err *Error) Error() string {
	return err.msg
}
