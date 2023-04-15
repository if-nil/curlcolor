package main

import (
	"fmt"
	"github.com/if-nil/curlcolor"
	"os"
)

func main() {
	if err := curlcolor.Run(os.Args[1:]); err != nil {
		fmt.Printf("\033[31m%s\033[0m\n", "curlcolor: "+err.Error())
		os.Exit(1)
	}
}
