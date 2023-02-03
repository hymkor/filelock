package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	if f, err := os.Open(os.Args[1]); err != nil {
		println(os.Stderr, err.Error())
		return
	} else {
		defer f.Close()
	}
	io.ReadAll(os.Stdin)
}
