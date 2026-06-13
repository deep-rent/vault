package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/deep-rent/nexus/app"
)

// The application version injected via -ldflags during build time.
var version = "v0.0.0"

func main() {
	if err := boot(os.Args, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func boot(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(stdout)

	showVersion := flags.Bool("v", false, "Display version and exit")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

  if *showVersion {
		_, _ = fmt.Fprintf(stdout, "Vouch %s\n", version)
		return nil
	}

  return app.RunAll(nil)
}
