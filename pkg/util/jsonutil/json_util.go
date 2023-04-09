package jsonutil

import "encoding/json"

func Unmarshal[T any](bytes []byte) (T, error) {
	var data T
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func Marshal(data any) []byte {
	message, _ := json.Marshal(data)
	return message
}
