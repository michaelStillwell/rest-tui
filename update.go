package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.resetBlur()
			return m, tea.Quit

		case key.Matches(msg, m.keymap.nextMethod):
			handleNextMethod(&m)

		case key.Matches(msg, m.keymap.prevMethod):
			handlePrevMethod(&m)

		case key.Matches(msg, m.keymap.next):
			cmd = handleNext(&m)
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keymap.prev):
			cmd := handlePrev(&m)
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.keymap.format):
			v := m.bodyInput.Value()
			m.bodyInput.SetValue(jsonFormat([]byte(v)))

		case key.Matches(msg, m.keymap.submit):
			go callUrl(m.method, m.urlInput.Value(), m.bodyInput.Value(), m.resChannel)

		case key.Matches(msg, m.keymap.clear):
			m.res = ""
			m.setContent()
		}

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	default:
		select {
		case res := <-m.resChannel:
			m.res = res
			m.setContent()

		default:
		}
	}

	m.sizeInputs()

	switch m.focus {
	case focusBody:
		m.bodyInput, cmd = m.bodyInput.Update(msg)
		cmds = append(cmds, cmd)
	case focusRes:
		m.resVp, cmd = m.resVp.Update(msg)
		cmds = append(cmds, cmd)
	case focusUrl:
		fallthrough
	default:
		m.urlInput, cmd = m.urlInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
