package main

import (
	"flag"
	"fmt"
	"os"
)

func readFlags() (configFile *string) {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Error, you need to indicate the configuration file")
		fmt.Println("Usage:")
		fmt.Println("       ./xkcd -c /path/to/config/file")
		os.Exit(1)
	}

	configFile = flag.String("c", "", "configuration file")
	flag.Parse()

	return configFile
}
