package args

import "flag"

// "flag" wrapper that constructs the args.Args{} struct for you with some,
// default values already setted - you need to use this default values to
// validate the user input.
//
// Should be one of the first functions to be run by your app before executing
// the main logic.
func Parse() Args {
	var sources sourcesFlag

	argv := Args{}

	flag.StringVar(&argv.Machine, "machine", "", "FTP server machine address.")
	flag.IntVar(&argv.Port, "port", -1, "FTP server port.")
	flag.StringVar(&argv.Username, "username", "", "FTP connection username.")
	flag.StringVar(&argv.Password, "password", "", "FTP connection password.")

	flag.Var(&sources, "src", "List of file paths that you want to copy.")
	flag.StringVar(&argv.Target, "target", "", "Path to the directory where everything will be saved.")
	flag.BoolVar(&argv.IsDryRun, "dry", false, "Just show the log information, without actually copying anything.")

	flag.Parse()

	// Copy the sources array before return to remove the String() and Set() methods.

	argv.Sources = make([]string, len(sources))

	copy(argv.Sources, sources)

	return argv
}
