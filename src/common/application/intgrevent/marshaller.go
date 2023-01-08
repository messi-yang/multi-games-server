package intgrevent

import "encoding/json"

func Unmarshal[T any](bytes []byte) (T, error) {
	var intgrEvent T
	err := json.Unmarshal(bytes, &intgrEvent)
	if err != nil {
		return intgrEvent, err
	}

	return intgrEvent, nil
}

func Marshal(event any) []byte {
	message, _ := json.Marshal(event)
	return message
}
