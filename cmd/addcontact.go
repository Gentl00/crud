/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


func NewAddCmd(s Server) *cobra.Command{
	var addcontactCmd = &cobra.Command{
	Use:   "add",
	Short: "Add contact to the diary",	
	Long: `Add contact to the diary.`,
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
		var errs []error
		firstname , err1 := cmd.Flags().GetBool("firstname")
		errs = append(errs, err1)
		email, err2 := cmd.Flags().GetBool("email")
		errs = append(errs, err2)
		contact, err3 := cmd.Flags().GetBool("phone")
		errs = append(errs, err3)
		if errs != nil {
			fmt.Println(errs)
			return 
			}
		switch{
		case firstname:
			if len(args) == 0 {
				fmt.Println("veuillez fournir le prénom")
				return 
			}
			err := s.AddFirstName(ctx, args[0])
			if err != nil {
				fmt.Println(err)
				return 
			}

		case email:
			if len(args) == 0 {
				fmt.Println("veuillez foournir l'email")
				return 
			}
			err := s.AddMail(ctx, args[0])
			if err != nil {
				fmt.Println(err)
				return 
			}
			
		case contact:
			if len(args) == 0 {
				fmt.Println("veuillez foournir le numéro de contact")
				return 
			}

			err := s.AddPhone(ctx, args[0])
			if err != nil {
				fmt.Println(err)
				return 
			}
		}
		},
	}
	addcontactCmd.Flags().Bool("firstname", false, "Add firstname")
	addcontactCmd.Flags().Bool("email", false, "Add email")
	addcontactCmd.Flags().Bool("phone", false, "Add phone number")
	return addcontactCmd
}





