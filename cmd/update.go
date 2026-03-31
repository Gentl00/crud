/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"


	"github.com/spf13/cobra"
)



func NewUpdateCmd(s Server) *cobra.Command{
	// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update contact's detail (name, firstname, email, phone)",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			var errs []error
			ctx := cmd.Context()
			name , err1 := cmd.Flags().GetBool("name")
			errs = append(errs, err1)
			firstname, err4 := cmd.Flags().GetBool("firstname")
			errs = append(errs, err4)
			email, err2 := cmd.Flags().GetBool("email")
			errs = append(errs, err2)
			contact, err3 := cmd.Flags().GetBool("phone")
			errs = append(errs, err3)
			if errs != nil {
				fmt.Println(errs)
				return 
			}
			switch{
			case name:
				if len(args) == 0 {
					fmt.Println("veuillez fournir l'id du contact à modifier")
					return 
				}
				
				id, err := strconv.Atoi(args[0])
				if err != nil {
					fmt.Println(err)
				}
				
				if len(args) == 1{
					fmt.Println("cette action supprime le nom contact")
					err := s.UpdateName(ctx, id, "")
					if err != nil {
						fmt.Println(err)
						return 
					}
				}
				if len(args) == 2 {
					fmt.Println("cette action modifie le nom contact")
					err := s.UpdateName(ctx, id, args[1])
					if err != nil {
						fmt.Println(err)
						return 
					}
				}

			case firstname:
				if len(args) == 0 {
					fmt.Println("veuillez fournir l'id du contact à modifier")
					return 
				}
				
				id, err := strconv.Atoi(args[0])
				if err != nil {
					fmt.Println(err)
				}
				
				if len(args) == 1{
					fmt.Println("cette action supprime le prénom contact")
					err := s.UpdateFirstName(ctx, id, "")
					if err != nil {
						fmt.Println(err)
						return 
					}
				}
				if len(args) == 2 {
					fmt.Println("cette action modifie le prénom contact")
					err := s.UpdateFirstName(ctx, id, args[1])
					if err != nil {
						fmt.Println(err)
						return 
					}
				}

			case email:
				if len(args) == 0 {
					fmt.Println("veuillez fournir l'id du contact à modifier")
					return 
				}

				id, err := strconv.Atoi(args[0])
				if err != nil {
					fmt.Println(err)
				}
				
				if len(args) == 1{
					fmt.Println("cette action supprime le mail du contact")
					err := s.UpdateEmail(ctx, id, "")
					if err != nil {
						fmt.Println(err)
						return 
					}
				}
				if len(args) == 2 {
					fmt.Println("cette action modifie le mail du contact")
					err := s.UpdateEmail(ctx, id, args[1])
					if err != nil {
						fmt.Println(err)
						return 
					}
				}
			case contact:
				if len(args) == 0 {
					fmt.Println("veuillez fournir l'id du contact à modifier")
					return 
				}

				id, err := strconv.Atoi(args[0])
				if err != nil {
					fmt.Println(err)
				}
				
				if len(args) == 1{
					fmt.Println("cette action supprime le numéro du contact")
					err := s.UpdateContact(ctx, id, "")
					if err != nil {
						fmt.Println(err)
						return 
					}
				}
				if len(args) == 2 {
					fmt.Println("cette action modifie le numéro du contact")
					err := s.UpdateContact(ctx, id, args[1])
					if err != nil {
						fmt.Println(err)
						return 
					}
				}
			}
		},
	}

	updateCmd.Flags().Bool("name", false, "Update contact's name")
	updateCmd.Flags().Bool("firstname", false, "Update contact's firstname")
	updateCmd.Flags().Bool("email", false, "Update contact's email")
	updateCmd.Flags().Bool("phone", false, "Update contact's phonenumber")
	return updateCmd
}
