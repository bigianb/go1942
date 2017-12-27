package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "A 1942 enulator written in go\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n\n")
	fmt.Fprintf(os.Stderr, "\tgo1942 [options]\n\n")
	fmt.Fprintf(os.Stderr, "Options are:\n\n")
	flag.PrintDefaults()
}

func main() {
	help := flag.Bool("help", false, "Show usage")
	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
		return
	}

	fmt.Fprintf(os.Stderr, "Done - exiting\n")
	return
}
