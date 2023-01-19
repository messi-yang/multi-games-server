package jsonmarshaller

import "encoding/json"

func Unmarshal[T any](bytes []byte) (T, error) {
	var intEvent T
	err := json.Unmarshal(bytes, &intEvent)
	if err != nil {
		return intEvent, err
	}

	return intEvent, nil
}

func Marshal(event any) []byte {
	message, _ := json.Marshal(event)
	return message
}
