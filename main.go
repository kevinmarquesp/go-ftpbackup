package main

import (
	"go-ftpbackup/args"
	"log"
)

func main() {
	argv := args.Parse()

	log.Println("- Arguments:", argv)
}
