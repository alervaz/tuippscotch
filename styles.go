package main

import "github.com/charmbracelet/lipgloss"

type styles struct {
	get                lipgloss.Style
	post               lipgloss.Style
	put                lipgloss.Style
	patch              lipgloss.Style
	delete             lipgloss.Style
	focusedAreaInput   lipgloss.Style
	unFocusedAreaInput lipgloss.Style
	focusedInput       lipgloss.Style
	unFocusedInput     lipgloss.Style
	title              lipgloss.Style
}

func getStyles() *styles {
	get := lipgloss.NewStyle().Foreground(lipgloss.Color("#10b981"))
	post := lipgloss.NewStyle().Foreground(lipgloss.Color("#f59e0b"))
	put := lipgloss.NewStyle().Foreground(lipgloss.Color("#3779e4"))
	patch := lipgloss.NewStyle().Foreground(lipgloss.Color("#006D5B"))
	deleteMethod := lipgloss.NewStyle().Foreground(lipgloss.Color("#ef4444"))
	focusedAreaInput := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#006D5B")).
		BorderStyle(lipgloss.RoundedBorder()).Width(80).Height(20)
	unFocusedAreaInput := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#5d636f")).
		BorderStyle(lipgloss.RoundedBorder()).Width(80).Height(6)
	focusedInput := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#006D5B")).
		BorderStyle(lipgloss.RoundedBorder()).Width(90).Height(1)
	unFocusedInput := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#5d636f")).
		BorderStyle(lipgloss.RoundedBorder()).Width(90).Height(1)
	title := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF")).MarginBottom(4)

	return &styles{
		get:                get,
		post:               post,
		put:                put,
		patch:              patch,
		delete:             deleteMethod,
		focusedAreaInput:   focusedAreaInput,
		unFocusedAreaInput: unFocusedAreaInput,
		focusedInput:       focusedInput,
		unFocusedInput:     unFocusedInput,
		title:              title,
	}
}
