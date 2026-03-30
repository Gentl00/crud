/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"crud/cmd"
	"crud/internal/config"
	"crud/internal/service"
	"crud/internal/storage"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	var conf config.ConfigDB
	repo := storage.NewStorage(&conf)
	if err := repo.NewPool(ctx); err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 3 {
		fmt.Println("Veuillez remplir le dernier champ")
		return
	}
	serv := service.NewContactService(repo)
	rootCmd := cmd.NewRootCmd(serv)
	rootCmd.Execute()
}
