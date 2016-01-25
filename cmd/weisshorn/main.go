package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func main() {
	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Displays all registered cards",
		Long:  `list is to display all registered cards known to the server.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list: " + strings.Join(args, " "))
		},
	}

	var cmdAdd = &cobra.Command{
		Use:   "add [compass card no] [name]",
		Short: "Registers a compass card with the server. If cards exists, it's a noop.",
		Long:  `e.g. add 01640812374011110001 "mom's card"`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add: " + strings.Join(args, " "))
		},
	}

	var cmdImport = &cobra.Command{
		Use:   "import [compass card no]",
		Short: "Imports a CSV dump.",
		Long:  `e.g. import 01640812374011110001`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("import: " + strings.Join(args, " "))
			Import()
		},
	}

	var rootCmd = &cobra.Command{Use: "weisshorn"}
	rootCmd.AddCommand(cmdList, cmdAdd, cmdImport)
	rootCmd.Execute()
}
