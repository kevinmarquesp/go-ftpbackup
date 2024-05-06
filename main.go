package main

import (
	"flag"
	"fmt"
	"go-ftpbackup/args"
	"go-ftpbackup/connect"
	"log"
	"os"
)

func main() {
	argv, err := args.Parse(flag.CommandLine)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	addr := argv.Machine + connect.ConvertIntPortToString(argv.Port)

	conn, err := connect.GetConnection(addr, argv.Username, argv.Password)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	entries, err := conn.List("/")
	if err != nil {
		fmt.Println(err)

		os.Exit(3)
	}

	for _, entry := range entries {
		log.Println(entry.Name)
	}
}
