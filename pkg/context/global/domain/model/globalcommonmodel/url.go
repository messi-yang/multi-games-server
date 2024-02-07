package globalcommonmodel

import (
	"fmt"
	"net/url"
)

type Url struct {
	value string
}

func NewUrl(rawUrl string) (u Url, err error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return u, err
	}

	if !parsedUrl.IsAbs() {
		return u, fmt.Errorf("invalid URL: %s", rawUrl)
	}

	if parsedUrl.Scheme != "https" {
		return u, fmt.Errorf("URL has to be https: %s", rawUrl)
	}

	return Url{value: rawUrl}, nil
}

func (u Url) String() string {
	return u.value
}
