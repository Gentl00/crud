/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command

func NewDeleteCmd(s Server) *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if len(args) == 0 {
				fmt.Println("Veuillez fournir l'ID")
				return
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := s.Delete(ctx, id); err != nil {
				fmt.Println(err)
				return
			}
		},
	}
	return deleteCmd
}
