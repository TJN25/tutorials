package utils

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	return s
}

// Model represents our Bubbletea model
type Model struct {
	Questions []Question
	Width     int
	Height    int
	Index     int
	Styles    *Styles
}

type Question struct {
	Question string
	Answer   string
	Input    Input
}

func NewQuestion(question string) Question {
	return Question{Question: question}
}

func NewShortQuestion(question string) Question {
	q := NewQuestion(question)
	model := NewShortAnswerField()
	q.Input = model
	return q
}

func NewLongQuestion(question string) Question {
	q := NewQuestion(question)
	field := NewLongAnswerField()
	q.Input = field
	return q
}

func New(questions []Question) *Model {
	styles := DefaultStyles()
	AnswerField := textinput.New()
	AnswerField.Placeholder = "Your answer here"
	AnswerField.Focus()
	return &Model{
		Questions: questions,
		Styles:    styles,
	}
}

// Init is the Bubbletea initial command
func (m Model) Init() tea.Cmd {
	return nil
}

// Update is the Bubbletea update function
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.Questions[m.Index]
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			current.Answer = current.Input.Value()
			log.Printf("question: %s, answer: %s", current.Question, current.Answer)
			m.Next()
			return m, current.Input.Blur
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}
	current.Input, cmd = current.Input.Update(msg)
	return m, cmd
}

func (m *Model) Next() {
	if m.Index < len(m.Questions)-1 {
		m.Index++
	} else {
		m.Index = 0
	}

}

// View is the Bubbletea view function
func (m Model) View() string {
	if m.Width == 0 {
		return "loading"
	}

	return lipgloss.Place(

		m.Width,
		m.Height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.Questions[m.Index].Question,
			m.Styles.InputField.Render(m.AnswerField.View())))
}
