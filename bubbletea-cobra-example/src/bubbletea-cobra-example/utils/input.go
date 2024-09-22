package utils

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Blink() tea.Msg
	Blur() tea.Msg
	Focus() tea.Cmd
	SetValue(string)
	Value() string
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}

type ShortAnswerField struct {
	textinput textinput.Model
}

type LongAnswerField struct {
	textarea textarea.Model
}

func (sa *ShortAnswerField) Value() string {
	return sa.textinput.Value()
}

func (la *LongAnswerField) Value() string {
	return la.textarea.Value()
}

func (sa *ShortAnswerField) Blur() {
	sa.textinput.Blur()
}

func (la *LongAnswerField) Blur() {
	la.textarea.Blur()
}

// textinput
func NewShortAnswerField() *ShortAnswerField {
	ti := textinput.New()
	return &ShortAnswerField{ti}
}

// textarea
func NewLongAnswerField() *LongAnswerField {
	ta := textarea.New()
	return &LongAnswerField{ta}
}
