/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func NewRootCmd(s Server) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "crud",
		Short: "Contact Diary",
		Long: `CLI tool for storing and managing contacts with basic 
	CRUD operations.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	rootCmd.AddCommand(NewAddCmd(s))
	rootCmd.AddCommand(NewListCmd(s))
	rootCmd.AddCommand(NewSearchCmd(s))
	rootCmd.AddCommand(NewUpdateCmd(s))
	rootCmd.AddCommand(NewDeleteCmd(s))
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

type Server interface {
	AddName(ctx context.Context, name string) error
	AddFirstName(ctx context.Context, firstname string) error
	AddMail(ctx context.Context, email string) error
	AddPhone(ctx context.Context, phone string) error
	Delete(ctx context.Context, id int) error
	UpdateName(ctx context.Context, id int, name string) error
	UpdateFirstName(ctx context.Context, id int, firstname string) error
	UpdateEmail(ctx context.Context, id int, email string) error
	UpdateContact(ctx context.Context, id int, contact string) error
	List(ctx context.Context) error
	ListSearch(ctx context.Context, name string) error
}
