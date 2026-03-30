/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)


func NewAddCmd(s Server) *cobra.Command{
	var addcontactCmd = &cobra.Command{
	Use:   "addcontact",
	Short: "A brief description of your command",	
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		err := s.AddName(ctx, args[0], args[1])
		var errs []error
		if err != nil {
			fmt.Println(err)
			}
		email, err2 := cmd.Flags().GetBool("email")
		errs = append(errs, err2)
		contact, err3 := cmd.Flags().GetBool("contact")
		errs = append(errs, err3)
		if errs != nil {
			fmt.Println(errs)
			return 
			}
		switch{
		case email:
			if len(args) == 0 {
				fmt.Println("veuillez foournir l'id du contact à modifier")
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
				fmt.Println("veuillez foournir l'id du contact à modifier")
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
	return addcontactCmd
}





