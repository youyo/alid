package alid

import (
	"fmt"
	"io"

	latest "github.com/tcnksm/go-latest"
)

const (
	owner      string = "youyo"
	repository string = "alid"
)

// VersionCheck is checking what is it correct version.
func VersionCheck(version string, errStream io.Writer) {
	githubTag := &latest.GithubTag{
		Owner:      owner,
		Repository: repository,
	}
	if res, err := latest.Check(githubTag, version); err == nil {
		if res.Outdated {
			fmt.Fprintf(errStream,
				"%s is not latest, you should upgrade to %s\n",
				version, res.Current)
		}
	} else {
		fmt.Fprintf(errStream,
			"Network is not unreachable. Can not check version.\n")
	}
}
