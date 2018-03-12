package main

import (
	"fmt"
	"os"
)

// ANSII Colors
const (
	CLR_0 = "\x1b[30;1m"

	CLR_R = "\x1b[31;1m"

	CLR_G = "\x1b[32;1m"

	CLR_Y = "\x1b[33;1m"

	CLR_B = "\x1b[34;1m"

	CLR_M = "\x1b[35;1m"

	CLR_C = "\x1b[36;1m"

	CLR_W = "\x1b[37;1m"

	CLR_N = "\x1b[0m"
)

func main() {
	var dirname string
	if len(os.Args) == 1 {
		dirname = "."
	} else {
		dirname = os.Args[1]
	}

	d, err := os.Open(dirname)
	if err != nil {
		panic(err)
	}
	defer d.Close()
	file, err := d.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for _, file := range file {
		if file.Mode().IsRegular() {
			fmt.Println(file.Mode(), CLR_M, file.Name(), CLR_N, file.Size(), "byte")
		} else {
			fmt.Println(file.Mode(), CLR_C, file.Name(), CLR_N)
		}
	}

}
