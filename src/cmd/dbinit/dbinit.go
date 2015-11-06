package main

import (
	"data"
	"fmt"
	"os"
)

func main() {
	err := data.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error initializing database ", err)
	}
}