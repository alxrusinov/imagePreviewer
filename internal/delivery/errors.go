package delivery

import "errors"

var (
	BadParamsError        = errors.New("bad params width/height")
	BadParsedAddressError = errors.New("incorrect url for source file")
	ReadImageError        = errors.New("reading image error")
	ImageProcessingError  = errors.New("image processing is not possible")
)
