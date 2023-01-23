package delivery

import "errors"

var (
	BadParsedAddressError = errors.New("incorrect url for source file")
	ImageDownloadingError = errors.New("can not download image")
)
