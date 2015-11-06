package main 

import (
	"fmt"
	"data"
	"os"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Fprintln(os.Stderr, "Error: not enough arguments")
		return
	}
	x := data.Task{
		Id: 1,
		Description: os.Args[1],
	}
	fmt.Println(x)
}