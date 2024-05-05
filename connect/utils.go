package connect

import "fmt"

// The "github.com/jlaffaye/ftp" package reads the port along side with the
// address string, this function helps convert the user specified argument from
// integer to a string with a ":" at the begining.
//
// 0 or -1 values will just return an empty string, this will make the
// connection function uses the default port.
func ConvertIntPortToString(port int) string {
	if port <= 0 {
		return ""
	}

	return fmt.Sprintf(":%d", port)
}
