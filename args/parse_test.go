package args

import (
	"flag"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("Should throw -machine error.", func(t *testing.T) {
		os.Args = []string{"cmd"}

		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		_, err := Parse(fs)

		if err == nil || err.Error() != "The -machine field is mandatory." {
			t.Errorf("Expected error for missing -machine, got %v", err)
		}
	})

	t.Run("Should throw -target error.", func(t *testing.T) {
		os.Args = []string{"cmd", "-machine", "localhost"}

		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		_, err := Parse(fs)

		if err == nil || err.Error() != "The -target field is mandatory." {
			t.Errorf("Expected error for missing -target, got %v", err)
		}
	})

	t.Run("Should throw empty -src list error.", func(t *testing.T) {
		os.Args = []string{"cmd", "-machine", "localhost", "-target", "/path/to/target"}

		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		_, err := Parse(fs)

		if err == nil || err.Error() != "No sources to backup specified (list length 0 not allowed)." {
			t.Errorf("Expected error for empty -src list, got %v", err)
		}
	})

	t.Run("Should assign default values.", func(t *testing.T) {
		os.Args = []string{"cmd", "-machine", "localhost", "-target", "/path/to/target", "-src", "/path/to/source"}

		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		args, err := Parse(fs)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if args.Port != -1 || args.Username != "" || args.Password != "" || args.IsDryRun {
			t.Errorf("Default values not correctly assigned in Args")
		}
	})

	t.Run("Should assign the user specified values.", func(t *testing.T) {
		os.Args = []string{"cmd", "-machine", "localhost", "-port", "21", "-username", "user", "-password", "pass", "-src", "/path/to/source", "-target", "/path/to/target", "-dry"}

		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		args, err := Parse(fs)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if args.Machine != "localhost" || args.Port != 21 || args.Username != "user" || args.Password != "pass" || args.Target != "/path/to/target" || len(args.Sources) != 1 || args.Sources[0] != "/path/to/source" || !args.IsDryRun {
			t.Errorf("User specified values not correctly assigned in Args")
		}
	})
}
