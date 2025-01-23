package main

import "github.com/charmbracelet/lipgloss"

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	cursorLineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("230"))

	placeholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("238"))

	endOfBufferStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("235"))

	focusedPlaceholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("238"))

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("57"))

	blurredBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.HiddenBorder())
)
