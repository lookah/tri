/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/lookah/tri/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a TODO item on the list",
	Long:  `Fak dis šit`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addRun(cmd *cobra.Command, args []string) {
	fmt.Println("Adding item function")

	items := []todo.Item{}

	for _, x := range args {
		fmt.Println(x)
		items = append(items, todo.Item{Text: x})
	}

	fmt.Println(items)
}
