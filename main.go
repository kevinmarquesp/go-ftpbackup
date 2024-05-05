package main

import (
	"fmt"
	"go-ftpbackup/args"
	"go-ftpbackup/connect"
	"log"
	"os"
)

func main() {
	argv, err := args.Parse()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	addr := argv.Machine + connect.ConvertIntPortToString(argv.Port)

	log.Println("- Address:", addr)
}
