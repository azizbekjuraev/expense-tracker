package main

import (
	"fmt"
	"os"
	"expense-tracker/cmd"
)

func main () {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
