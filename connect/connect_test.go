package connect

import (
	"testing"
	"time"
)

func TestGetConnection(t *testing.T) {
	t.Run("Should keep trying for a non-existing server for three seconds", func(t *testing.T) {
		done := make(chan struct{})

		go func() {
			GetConnection("nonexisting.server", "username", "password")
			close(done)
		}()

		select {
		case <-time.After(3 * time.Second):
			// After at least 3 seconds, stop trying to connect
			t.Log("Test ran for at least 3 seconds")
		case <-done:
			// If connection is successful, stop the timer
			t.Error("GetConnection exited before 3 seconds")
		}
	})

	// TODO: Add more tests with a dummy FTP local server.
}
