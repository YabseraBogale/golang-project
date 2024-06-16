package main

import (
	"bufio"
	"os"

	"github.com/YabseraBogale/golang-project/golanglsp/rpc"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Text()
		handler(msg)
	}
}

func handler(_ any) {}
