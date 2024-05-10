package connect

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetConnection(t *testing.T) {
	t.Run("(Integration) Should connect to the FTP server without any errors.", func(t *testing.T) {
		done := make(chan struct{})

		go func() {
			_, err := GetConnection("localhost:2121", "kevin", "kevin123")
			assert.NoError(t, err)
		}()

		select {
		case <-time.After(5 * time.Second):
			t.Log("Connection timeout. Did you start the integration server?")

		case <-done:
			t.Log("Successfully connected to the FTP server.")
		}
	})

	t.Run("Should keep trying for a non-existing server for three seconds", func(t *testing.T) {
		done := make(chan struct{})

		go func() {
			GetConnection("nonexisting.server", "username", "password")
			close(done)
		}()

		select {
		case <-time.After(3 * time.Second):
			t.Log("Test ran for at least 3 seconds")

		case <-done:
			t.Error("GetConnection exited before 3 seconds")
		}
	})
}
