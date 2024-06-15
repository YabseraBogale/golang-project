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
