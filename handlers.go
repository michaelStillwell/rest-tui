package main

import (
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func handleNextMethod(m *Model) {
	switch strings.ToUpper(m.method) {
	case http.MethodGet:
		m.method = http.MethodPost
	case http.MethodPost:
		m.method = http.MethodPut
	case http.MethodPut:
		m.method = http.MethodDelete
	case http.MethodDelete:
		m.method = http.MethodPatch
	case http.MethodPatch:
		m.method = http.MethodGet
	}
	m.urlInput.Prompt = normalizeMethod(m.method)
}

func handlePrevMethod(m *Model) {
	switch strings.ToUpper(m.method) {
	case http.MethodGet:
		m.method = http.MethodPatch
	case http.MethodPost:
		m.method = http.MethodGet
	case http.MethodPut:
		m.method = http.MethodPost
	case http.MethodDelete:
		m.method = http.MethodPut
	case http.MethodPatch:
		m.method = http.MethodDelete
	}
	m.urlInput.Prompt = normalizeMethod(m.method)
}

func handleNext(m *Model) tea.Cmd {
	m.resetBlur()
	m.focus++
	if m.focus > focusRes {
		m.focus = focusUrl
	}
	return m.setFocus()
}

func handlePrev(m *Model) tea.Cmd {
	m.resetBlur()
	m.focus--
	if m.focus < focusUrl {
		m.focus = focusRes
	}
	return m.setFocus()
}
