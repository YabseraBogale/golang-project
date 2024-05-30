package main

import (
	"fmt"
	"os"
	// "github.com/mymmrac/telego"
)

func main() {
	// Get Bot token from environment variables
	botToken := os.Getenv("TOKEN")
	fmt.Println(string(botToken))

}
