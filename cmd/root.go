/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a TODO managing application",
	Long: `Tri will help you organize and handle all your TODO tasks.
	It's supposed to be simple and easy to use.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tri.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory, please set data file using --datafile")
	}

	RootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".testDataFile.json", "data file to store TODOs")

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
