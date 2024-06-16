package rpc_test

import (
	"testing"

	"github.com/YabseraBogale/golang-project/golanglsp/rpc"
)

type EncodeTesting struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodingMessage(EncodeTesting{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	expected := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	msg, contentLength, err := rpc.DecodingMessage([]byte(expected))
	if err != nil {
		t.Fatalf("Expected: %d Actual: %d", 15, contentLength)
	}
	if contentLength != 15 {
		t.Fatalf("Expected: %d Actual: %d", 15, contentLength)
	}
	if msg != "hi" {
		t.Fatalf("Expected: hi Actual: %s", msg)
	}
}
