package connect

import (
	"fmt"
	"time"

	"github.com/jlaffaye/ftp"
)

// GetConnection attempts to establish an FTP connection to the specified
// address using the provided username and password.
//
// It continuously tries to dial the address until a connection is successfully
// established. Once a connection is established, it attempts to log in with
// the provided credentials.
func GetConnection(addr, username, password string) (*ftp.ServerConn, error) {
	var (
		conn *ftp.ServerConn
		err  error
	)

	// Try until it works because it's meat to be ran in background by some other script.

	for {
		conn, err = ftp.Dial(addr)
		if err == nil {
			break
		}

		fmt.Println(err)
		time.Sleep(time.Second * 3)
	}

	err = conn.Login(username, password)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
