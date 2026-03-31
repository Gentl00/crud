/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func NewUpdateCmd(s Server) *cobra.Command {
	// updateCmd represents the update command
	var name string
	var firstname string
	var email string
	var phone string
	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update contact's detail (name, firstname, email, phone)",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if len(args) == 0 {
				fmt.Println("Veuillez fournir un nom")
				return
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			if name != "" {
				err := s.UpdateName(ctx, id, name)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if firstname != "" {
				err := s.UpdateFirstName(ctx, id, firstname)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if email != "" {
				err := s.UpdateEmail(ctx, id, email)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if phone != "" {
				err := s.UpdateContact(ctx, id, phone)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

		},
	}
	updateCmd.Flags().StringVar(&firstname, "firstname", "", "Update contact's name")
	updateCmd.Flags().StringVar(&name, "name", "", "Update contact's firstname")
	updateCmd.Flags().StringVar(&email, "email", "", "Update contact's email")
	updateCmd.Flags().StringVar(&phone, "phone", "", "Update contact's phonenumber")
	return updateCmd
}
