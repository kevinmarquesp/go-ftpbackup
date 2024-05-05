package args

import "fmt"

// This struct will be used by the application to access *all* input data that
// the user passed. If more the user needs more data (not CLI flags only) to
// pass, it should be acced here.
type Args struct {
	Machine  string
	Port     int
	Username string
	Password string
	Target   string
	Sources  []string
	IsDryRun bool
}

// NOTE: The struct below is used as a custom []string type for the "flags" package.

type sourcesFlag []string

func (sf *sourcesFlag) String() string {
	res := ""

	for i := range *sf {
		if i == len(*sf)-1 {
			res += (*sf)[i]

			continue
		}

		res += fmt.Sprintf("%s, ", (*sf)[i])
	}

	return res
}

func (sf *sourcesFlag) Set(val string) error {
	*sf = append(*sf, val)

	return nil
}
