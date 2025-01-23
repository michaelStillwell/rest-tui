package main

import (
	"net/http"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	focusUrl = iota
	focusBody
	focusRes
)

type Model struct {
	width      int
	height     int
	keymap     keymap
	help       help.Model
	bodyInput  textarea.Model
	resVp      viewport.Model
	resChannel chan string
	res        string
	urlInput   textarea.Model
	method     string
	focus      int
}

func NewModel() Model {
	m := Model{
		bodyInput:  newTextarea(true),
		resVp:      newResVp(),
		resChannel: make(chan string, 1),
		res:        "",
		urlInput:   newUrlInput(),
		method:     http.MethodGet,
		help:       help.New(),
		keymap: keymap{
			next: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "next"),
			),
			prev: key.NewBinding(
				key.WithKeys("shift+tab"),
				key.WithHelp("shift+tab", "prev"),
			),
			nextMethod: key.NewBinding(
				key.WithKeys("ctrl+n"),
				key.WithHelp("ctrl+n", "next"),
			),
			prevMethod: key.NewBinding(
				key.WithKeys("ctrl+p"),
				key.WithHelp("ctrl+p", "prev"),
			),
			quit: key.NewBinding(
				key.WithKeys("esc", "ctrl+c"),
				key.WithHelp("esc", "quit"),
			),
			format: key.NewBinding(
				key.WithKeys("ctrl+f"),
				key.WithHelp("ctrl+f", "format"),
			),
			submit: key.NewBinding(
				key.WithKeys("ctrl+enter", "ctrl+j"),
				key.WithHelp("ctrl+enter", "submit request"),
			),
			clear: key.NewBinding(
				key.WithKeys("ctrl+k"),
				key.WithHelp("ctrl+k", "clear response"),
			),
		},
	}

	m.urlInput.Focus()
	m.sizeInputs()
	m.setContent()
	return m
}

func (m *Model) resetBlur() {
	m.urlInput.Blur()
	// m.resVp.Blur()
	m.bodyInput.Blur()
}

func (m *Model) setFocus() tea.Cmd {
	var cmd tea.Cmd
	switch m.focus {
	case focusBody:
		cmd = m.bodyInput.Focus()
	case focusRes:
		// cmd = m.resVp.Focus()
	case focusUrl:
		cmd = m.urlInput.Focus()
	}
	return cmd
}
