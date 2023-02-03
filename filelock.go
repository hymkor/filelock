package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	if f, err := os.Open(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	} else {
		defer f.Close()
	}
	ioutil.ReadAll(os.Stdin)
}
