package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := initModel()

	file, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
