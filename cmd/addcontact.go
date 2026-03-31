/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewAddCmd(s Server) *cobra.Command {
	var firstname string
	var email string
	var phone string
	var addcontactCmd = &cobra.Command{
		Use:   "add",
		Short: "Add contact to the diary",
		Long:  `Add contact to the diary.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if len(args) == 0 {
				fmt.Println("Veuillez fournir un nom")
				return
			}
			err := s.AddName(ctx, args[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			if firstname != "" {
				err := s.AddFirstName(ctx, firstname)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Prénom ajouté")
			}

			if email != "" {
				err := s.AddMail(ctx, email)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("email ajouté")
			}

			if phone != "" {
				err := s.AddPhone(ctx, phone)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		},
	}
	addcontactCmd.Flags().StringVar(&firstname, "firstname", "", "Add firstname")
	addcontactCmd.Flags().StringVar(&email, "email", "", "Add email")
	addcontactCmd.Flags().StringVar(&phone, "phone", "", "Add phone number")
	return addcontactCmd
}
