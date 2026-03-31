/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}

func NewSearchCmd(s Server) *cobra.Command {
	var searchCmd = &cobra.Command{
		Use:   "search",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if len(args) == 0 {
				fmt.Println("Veuillez fournir le nom")
				return
			}
			err := s.ListSearch(ctx, args[0])
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	return searchCmd
}
