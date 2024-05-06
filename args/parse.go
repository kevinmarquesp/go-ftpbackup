package args

import (
	"errors"
	"flag"
	"os"
)

func validateParsedValues(argv Args) error {
	if argv.Machine == "" {
		return errors.New("The -machine field is mandatory.")
	}

	if argv.Target == "" {
		return errors.New("The -target field is mandatory.")
	}

	if len(argv.Sources) == 0 {
		return errors.New("No sources to backup specified (list length 0 not allowed).")
	}

	return nil
}

// "flag" wrapper that constructs the args.Args{} struct for you with some
// default values already setted - the default port value (-1) is not a bug,
// it means that the user allows the program to connect without this option.
//
// The argument is necessary to make this code easy to test. In normal
// conditions, use the flag.CommandLine struct.
func Parse(fs *flag.FlagSet) (Args, error) {
	var sources sourcesFlag

	argv := Args{}

	fs.StringVar(&argv.Machine, "machine", "", "FTP server machine address.")
	fs.IntVar(&argv.Port, "port", -1, "FTP server port.")
	fs.StringVar(&argv.Username, "username", "", "FTP connection username.")
	fs.StringVar(&argv.Password, "password", "", "FTP connection password.")

	fs.Var(&sources, "src", "List of file paths that you want to copy.")
	fs.StringVar(&argv.Target, "target", "", "Path to the directory where everything will be saved.")
	fs.BoolVar(&argv.IsDryRun, "dry", false, "Just show the log information, without actually copying anything.")

	fs.Parse(os.Args[1:])

	// Copy the sources array before return to remove the String() and Set() methods.

	argv.Sources = make([]string, len(sources))

	copy(argv.Sources, sources)

	err := validateParsedValues(argv) // can be nil

	return argv, err
}
