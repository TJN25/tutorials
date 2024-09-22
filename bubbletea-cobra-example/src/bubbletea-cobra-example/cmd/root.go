/*
Copyright © 2024 dani
*/
package cmd

import (
	"log"
	"os"

	"github.com/TJN25/bubbletea-cobra-example/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bubbletea-cobra-example",
	Short: "Learn how to use Bubbletea with Cobra",
	Long:  `Learn how to use Bubbletea with Cobra. This example is a simple project.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: rootRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bubbletea-cobra-example.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func rootRun(cmd *cobra.Command, args []string) {
	questions := []utils.Question{
		utils.NewShortQuestion("what is your name?"),
		utils.NewShortQuestion("what is your favourite editor?"),
		utils.NewLongQuestion("what is your favourite quote?")}
	m := utils.New(questions)

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}
