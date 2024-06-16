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
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	contentLength, err := rpc.DecodingMessage([]byte(expected))
	if err != nil {
		t.Fatalf("Expected: %d Actual: %d", expected, contentLength)
	}
	if contentLength != 16 {
		t.Fatalf("Expected: %d Actual: %d", 16, contentLength)
	}
}
