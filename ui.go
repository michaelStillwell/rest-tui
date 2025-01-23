package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	initialInputs = 2
	maxInputs     = 6
	minInputs     = 1
	urlHeight     = 1
	helpHeight    = 5
)

type keymap = struct {
	submit, clear, next, prev, nextMethod, prevMethod, quit, format key.Binding
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m *Model) sizeInputs() {
	m.bodyInput.SetWidth(m.width / 2)
	m.bodyInput.SetHeight(m.height - urlHeight - helpHeight)
	m.urlInput.SetWidth(m.width)
	m.resVp.Width = m.width / 2
	m.resVp.Height = m.height - urlHeight - helpHeight
}

func (m *Model) setContent() {
	m.resVp.SetContent(
		lipgloss.NewStyle().Width(m.width / 2).Render(m.res))
}

func (m Model) Help() string {
	return m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
		m.keymap.prev,
		m.keymap.format,
		m.keymap.submit,
		m.keymap.clear,
		m.keymap.quit,
	})
}

func (m Model) View() string {

	var views []string
	views = append(views, m.urlInput.View())

	var body []string

	body = append(body, m.bodyInput.View())
	if m.focus == focusRes {
		body = append(body, focusedBorderStyle.Render(m.resVp.View()))
	} else {
		body = append(body, blurredBorderStyle.Render(m.resVp.View()))
	}

	views = append(views, lipgloss.JoinHorizontal(lipgloss.Center, body...))
	views = append(views, m.Help())

	return lipgloss.JoinVertical(lipgloss.Top, views...)
}
