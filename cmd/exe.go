package cmd

import (
	"fmt"

	"github.com/trshpuppy/sickem/utils"
)

func Sickem() string {
	// Get command line args
	args := utils.GetFlags()

	return fmt.Sprintf("args: %b\n", args.Nmap)
}
