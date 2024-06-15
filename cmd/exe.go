package cmd

import (
	"fmt"

	"github.com/trshpuppy/sickem/cmd/config"
	"github.com/trshpuppy/sickem/utils"
)

func Sickem() string {
	// Get command line args
	args := utils.GetFlags()

	// Check for config
	conf := config.CheckForConfig()

	return fmt.Sprintf("args: %b, conf: %v\n", args.Nmap, conf)
}
