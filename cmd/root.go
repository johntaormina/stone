package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "stone",
		Short: "Minimal, distractionless note taking.",
		// Long:  `Long description`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	printHello()
		// },
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	if !configExists() {
		cobra.CheckErr("Please initialize config with `stone init`.")
	}
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AddConfigPath(stoneHome)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	viper.SetDefault("editor", "vim")
	viper.SetDefault("date", time.Now().Format("2006-01-02"))
	viper.Debug()
}
