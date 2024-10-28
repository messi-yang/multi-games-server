package globalcommonmodel

import (
	"fmt"
	"net/url"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type Url struct {
	value string
}

// Interface Implementation Check
var _ domain.ValueObject[Url] = (*Url)(nil)

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

func (url Url) IsEqual(otherUrl Url) bool {
	return url == otherUrl
}

func (url Url) String() string {
	return url.value
}
