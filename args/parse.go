package args

import "flag"

// TODO: add a documentation for this functions
func Parse() Args {
	var sources sourcesFlag

	argv := Args{}

	flag.StringVar(&argv.Machine, "machine", "", "FTP server machine address.")
	flag.IntVar(&argv.Port, "port", -1, "FTP server port.")
	flag.StringVar(&argv.Username, "username", "", "FTP connection username.")
	flag.StringVar(&argv.Password, "password", "", "FTP connection password.")

	flag.Var(&sources, "src", "List of file paths that you want to copy.")

	copy(argv.Sources, sources) // removes the String() and Set() methods of the custom flags list type

	flag.StringVar(&argv.Target, "target", "", "Path to the directory where everything will be saved.")
	flag.BoolVar(&argv.IsDryRun, "dry", false, "Just show the log information, without actually copying anything.")

	flag.Parse()

	return argv
}
