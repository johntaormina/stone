package cmd

import "os"

var (
	home, err = os.UserHomeDir()
	stoneHome = home + "/.stone"
)
