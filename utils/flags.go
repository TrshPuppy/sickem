package utils

import (
	"flag"
	"fmt"
)

type Flags struct {
	nmap *bool
}

func GetFlags() Flags {
	nmapFLag := flag.Bool("nmap", false, "Ask sickem to run nmap")

	fl := Flags{nmap: nmapFLag}

	fmt.Printf("fl: %s\n", fl)

	return fl
}
