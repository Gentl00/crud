/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func NewRootCmd(s Server) *cobra.Command{
	// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crud",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	}
	rootCmd.AddCommand(NewAddCmd(s))
	rootCmd.AddCommand(NewListCmd(s))
	rootCmd.AddCommand(NewSearchCmd(s))
	return rootCmd
}



// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.



type Server interface{
	AddName(ctx context.Context, name string, lastname string) error
	UpdateName(ctx context.Context, id int, name string, lastname string) error
	UpdateEmail(ctx context.Context, id int, email string) error
	UpdateContact(ctx context.Context, id int, contact string) error
	List(ctx context.Context) error
	ListSearch(ctx context.Context, name string) error
}
