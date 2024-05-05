package main

import (
	"fmt"
	"go-ftpbackup/args"
	"log"
	"os"
)

func main() {
	argv, err := args.Parse()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	log.Println("- Arguments:", argv)
}
