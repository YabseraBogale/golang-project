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
		msg := scanner.Text()
		handler(msg)
	}
}

func handler(_ any) {}
func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey you didn't giveme a good file")
	}
	return log.New(logfile, "[golanglsp]", log.Ldate|log.LUTC|log.Lshortfile)
}
