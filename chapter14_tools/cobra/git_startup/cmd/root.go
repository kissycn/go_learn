package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git",
	Short: "Git is a distributed version control system.",
	Long:  `Git is a free and open source distributed version control system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git run")
	},
}

func Execute() {
	rootCmd.Execute()
}
