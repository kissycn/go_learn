package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var echoCmd = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long:  `echo is for echoing anything back.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))
	},
}
