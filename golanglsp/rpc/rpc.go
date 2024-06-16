package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func EncodingMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}
func DecodingMessage(msg []byte) (string, int, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", 0, errors.New("not found")
	}
	contentLengthByte := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthByte))
	if err != nil {
		return "", 0, err
	}
	_ = content
	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", 0, err
	}
	return baseMessage.Method, contentLength, nil

}
