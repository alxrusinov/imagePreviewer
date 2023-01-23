package delivery

import "net/url"

func createImageAddress(link string) (*url.URL, error) {

	parsed, err := url.Parse(link)

	if err != nil {
		return nil, err
	}

	parsed.Scheme = "https"

	return parsed, nil
}
