package delivery

import "errors"

var (
	ErrBadParams        = errors.New("bad params width/height")
	ErrBadParsedAddress = errors.New("incorrect url for source file")
	ErrReadImage        = errors.New("reading image error")
	ErrImageProcessing  = errors.New("image processing is not possible")
)
