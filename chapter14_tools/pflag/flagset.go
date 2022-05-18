package main

import (
	"flag"
	"fmt"
)

func main() {
	args := []string{"-intflag", "10", "-stringflag", "test"}

	var intflag int
	var boolflag bool
	var stringflag string

	fs := flag.NewFlagSet("MyFlagSet", flag.ContinueOnError)
	fs.IntVar(&intflag, "intflag", 0, "int flag value")
	fs.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	fs.StringVar(&stringflag, "stringflag", "default", "string flag value")

	fs.Parse(args)

	fmt.Println("int flag:", intflag)
	fmt.Println("bool flag:", boolflag)
	fmt.Println("string flag:", stringflag)
	//flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	//flagSet.StringVar(&version, "version", "v1.0.0", "Print version information and quit.")
	//
	//flagSet.Parse([]string{"--version", "v1.0.0"})
	//fmt.Println(version)
}
