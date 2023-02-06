package delivery

import (
	"fmt"
	"net/url"
	"strings"
)

func createImageAddress(link string) (*url.URL, error) {
	trimmedLink := strings.TrimLeft(link, "/")

	joinedPath, err := url.JoinPath("http://", trimmedLink)
	if err != nil {
		return nil, err
	}

	parsed, err := url.Parse(joinedPath)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func createFileName(link string, width, height int) string {
	splintedLink := strings.Split(link, "/")

	originalFileName := splintedLink[len(splintedLink)-1]

	trimmedFileName := strings.TrimLeft(originalFileName, "_")

	splintedFileName := strings.Split(trimmedFileName, ".")

	extension := splintedFileName[1]
	original := splintedFileName[0]

	name := strings.Split(original, "_")[0]

	result := fmt.Sprintf("%s_%dx%d.%s", name, width, height, extension)

	return result
}
