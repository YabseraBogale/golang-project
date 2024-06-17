package main

import (
	"bufio"
	"log"
	"os"

	"github.com/YabseraBogale/golang-project/golanglsp/rpc"
)

func main() {
	logger := getLogger("./logger.txt")
	logger.Println("Hey i started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodingMessage(msg)
		if err != nil {
			logger.Println("Got error: %s", err)
		}

		handler(logger, method, content)
	}
}

func handler(logger *log.Logger, method string, content []byte) {
	logger.Println(method)
}
func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey you didn't giveme a good file")
	}
	return log.New(logfile, "[golanglsp]", log.Ldate|log.LUTC|log.Lshortfile)
}
