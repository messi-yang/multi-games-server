package gziputil

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Ungzip(data []byte) ([]byte, error) {
	gunzip, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer gunzip.Close()
	compressedData, err := io.ReadAll(gunzip)
	if err != nil {
		return nil, err
	}
	return compressedData, nil
}

func Gzip(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	err = gz.Flush()
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
