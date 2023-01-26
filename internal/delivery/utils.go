package delivery

import (
	"net/url"
	"strings"
)

func createImageAddress(link string) (*url.URL, error) {

	trimmedLink := strings.TrimLeft(link, "/")

	joinedPath, err := url.JoinPath("https://", trimmedLink)

	if err != nil {
		return nil, err
	}

	parsed, err := url.Parse(joinedPath)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}
