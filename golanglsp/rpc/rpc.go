package rpc

import "encoding/json"

func Encoding(msg string) ([]byte, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return content, nil
}
