package worldcommonmodel

import (
	"fmt"
	"regexp"
)

type EmbedCode struct {
	value string
}

func NewEmbedCode(embedCodeString string) (embedCode EmbedCode, err error) {
	matched, err := regexp.Match("^<iframe.*?src=[\"'](.*?)[\"'].*?></iframe>$", []byte(embedCodeString))
	if err != nil {
		return embedCode, err
	}
	if !matched {
		return embedCode, fmt.Errorf("the embed code %s is not valid", embedCodeString)
	}

	return EmbedCode{value: embedCodeString}, nil
}

func (u EmbedCode) String() string {
	return u.value
}
