package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var version string

func main() {
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.StringVar(&version, "version", "v1.0.0", "Print version information and quit.")

	fmt.Println(version)
}
