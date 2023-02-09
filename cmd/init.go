package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Init config and required folders for storage.",
		Run: func(cmd *cobra.Command, args []string) {
			command()
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func command() {
	if !configExists() {
		createConfig()
	} else {
		fmt.Println("./stone folder already exists.")
	}
}

func configExists() bool {
	_, err = os.Stat(stoneHome)
	// err nil if folder exists
	return err == nil
}

func createConfig() {
	// write ./stone file to home dir
	err := os.Mkdir(stoneHome, 0750)
	cobra.CheckErr(err)
	// write config.yaml to ./stone/
	err = os.WriteFile(stoneHome+"/config.txt", []byte("name: bob barker"), 0666)
	cobra.CheckErr(err)
	// creates notes folder to ./stone/
	err = os.Mkdir(stoneHome+"/notes", 0750)
	cobra.CheckErr(err)
}
