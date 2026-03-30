package ui

import (
	"strconv"
	"os"
	"crud/internal/storage"

	_"github.com/charmbracelet/lipgloss"
	"github.com/olekukonko/tablewriter"
)

func TaskStyle(tasks []storage.ContactDTO) error{
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Nom", "Prenom", "Email", "Numéro"})
	for i := range tasks{
		task := tasks[i]
		styleID := strconv.Itoa(task.ID)
		table.Append([]string{styleID, task.Nom, task.Prenom, task.Email, task.Numero})

	}
	err := table.Render()
	if err != nil {
		return err
	}

	return nil
}