package main

import (
	"flag"
	"fmt"
	"go-ftpbackup/args"
	"go-ftpbackup/connect"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	argv, err := args.Parse(flag.CommandLine)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	addr := argv.Machine + connect.ConvertIntPortToString(argv.Port)

	var conn *ftp.ServerConn

	// Try until it works because it's meat to be ran in background by some other script.

	for {
		conn, err = ftp.Dial(addr)
		if err == nil {
			break
		}

		fmt.Println(err)
		time.Sleep(time.Second * 3)
	}

	defer conn.Quit()

	err = conn.Login(argv.Username, argv.Password)
	if err != nil {
		fmt.Println(err)

		os.Exit(2)
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
