package client

var jpegMimeType = [2]string{"image/jpeg", "image/jpg"}

func isJpegFileType(fileType string) (result bool) {
	for _, val := range jpegMimeType {
		if val == fileType {
			result = true
		}
	}

	return result
}
