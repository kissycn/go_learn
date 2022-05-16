package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	Execute()
}

var (
	name string
	age  int
)

var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "Test Demo",
	Long:  `A cobra test demo`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("My name is %s, my age is %d\n", name, age)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("init")
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "personal name")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "personal age")
}

func initConfig() {
	fmt.Println("init config")
}
