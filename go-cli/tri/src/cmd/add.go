/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"

	"fmt"

	"github.com/TJN25/tri/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	Long: `Take input from the user and save it as a new task.`,
	Run: addRun,
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

	items := []todo.Item{};


	for _,item := range args {
		items = append(items, todo.Item{Text:item})
	}

	err := todo.SaveItems("/Users/nicth99p/bin/projects/tutorials/go-cli/tri/test.json", items)

	if err != nil {
		fmt.Errorf("%v", err);
	}




	}
