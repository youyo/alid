package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/youyo/alid/lib/alid"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		region  string
		all     bool
		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.StringVar(&region, "region", "OS Environment 'AWS_REGION'", "Use AWS Region.")
	flags.StringVar(&region, "r", "OS Environment 'AWS_REGION'", "Use AWS Region.(Short)")
	flags.BoolVar(&all, "all", false, "Print AMI all-information.")
	flags.BoolVar(&all, "i", false, "Print AMI all-information.(Short)")
	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		alid.VersionCheck(Version, cli.errStream)
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	r, _ := alid.SelectRegion(region)
	s, err := alid.NewService(alid.NewConfig(r))
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err)
		return ExitCodeError
	}
	amiInfo, err := s.FetchLatestAmiInfo()
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err)
		return ExitCodeError
	}

	if all {
		fmt.Fprintf(cli.outStream, "%v\n", amiInfo)
	} else {
		fmt.Fprintf(cli.outStream, "%s\n", *amiInfo.ImageId)
	}

	return ExitCodeOK
}
