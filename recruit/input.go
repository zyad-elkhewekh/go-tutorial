package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Value() string
	Blur() tea.Msg
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}

func (sa *ShortAnswerField) Value() string {
	return sa.textinput.Value()
}
func (sa *ShortAnswerField) Blur() tea.Msg {
	return sa.textinput.Blur
}

func (la *LongAnswerField) Value() string {
	return la.textarea.Value()
}
func (la *LongAnswerField) Blur() tea.Msg {
	return la.textarea.Blur
}

type ShortAnswerField struct {
	textinput textinput.Model
}
type LongAnswerField struct {
	textarea textarea.Model
}

func NewLongAnswerField() *LongAnswerField {
	ta := textarea.New()
	ta.Placeholder = "ctrl+s to save +c to exit"
	ta.Focus()
	return &LongAnswerField{ta}
}
func NewShortAnswerField() *ShortAnswerField {
	ti := textinput.New()
	ti.Placeholder = "ctrl+s to save +c to exit"
	ti.Focus()
	return &ShortAnswerField{ti}
}

func (sa *ShortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	sa.textinput, cmd = sa.textinput.Update(msg)
	return sa, cmd
}
func (la *LongAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	la.textarea, cmd = la.textarea.Update(msg)
	return la, cmd
}

func (sa *ShortAnswerField) View() string {
	return sa.textinput.View()
}
func (la *LongAnswerField) View() string {
	return la.textarea.View()
}
