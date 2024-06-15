package utils

import (
	"flag"
)

type Flags struct {
	Nmap bool
}

func GetFlags() Flags {
	nmapFLag := flag.Bool("nmap", false, "Ask sickem to run nmap")

	flag.Parse()

	fl := Flags{Nmap: *nmapFLag}

	return fl
}
