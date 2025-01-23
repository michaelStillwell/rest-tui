package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
)

func newTextarea(ln bool) textarea.Model {
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = "Type something"
	t.ShowLineNumbers = ln
	t.CharLimit = 0
	t.Cursor.Style = cursorStyle
	t.FocusedStyle.Placeholder = focusedPlaceholderStyle
	t.BlurredStyle.Placeholder = placeholderStyle
	t.FocusedStyle.CursorLine = cursorLineStyle
	t.FocusedStyle.Base = focusedBorderStyle
	t.BlurredStyle.Base = blurredBorderStyle
	t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	t.KeyMap.DeleteWordBackward.SetEnabled(false)
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Blur()
	return t
}

func newUrlInput() textarea.Model {
	u := textarea.New()
	u.Prompt = normalizeMethod("get")
	u.Placeholder = "http://localhost:3000"
	u.ShowLineNumbers = false
	u.Cursor.Style = cursorStyle
	u.FocusedStyle.Placeholder = focusedPlaceholderStyle
	u.BlurredStyle.Placeholder = placeholderStyle
	u.FocusedStyle.CursorLine = cursorLineStyle
	u.FocusedStyle.Base = focusedBorderStyle
	u.BlurredStyle.Base = blurredBorderStyle
	u.FocusedStyle.EndOfBuffer = endOfBufferStyle
	u.BlurredStyle.EndOfBuffer = endOfBufferStyle
	u.KeyMap.InsertNewline.SetEnabled(false)
	u.Focus()
	u.SetHeight(urlHeight)
	return u
}

func newResVp() viewport.Model {
	v := viewport.New(1, 1)

	v.Style.Border(focusedBorderStyle.GetBorder())
	return v
}
