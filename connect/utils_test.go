package connect

import "testing"

func TestConvertIntPortToString(t *testing.T) {
	t.Run("Should return empty as default values.", func(t *testing.T) {
		got := ConvertIntPortToString(0)
		want := ""

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		got = ConvertIntPortToString(-1)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Should return a simple port string.", func(t *testing.T) {
		got := ConvertIntPortToString(21)
		want := ":21"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
