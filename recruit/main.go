package main

import (
	//"fmt"
	"log"
	//"os"

	//"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen int

const (
	firstScreen screen = iota
	registerationScreen
	questionScreen
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func defaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("45")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(120)
	return s
}

type model struct {
	questions  []Question
	user_input []string
	movement   []string
	screen     screen
	cursor     int
	height     int
	width      int
	indx       int
	//answerField textinput.Model
	styles *Styles
}

type Question struct {
	question string
	answer   string
	input    Input
}

func newQuestion(question string) Question {
	return Question{question: question}
}

func newShortQuestion(question string) Question {
	q := newQuestion(question)
	field := NewShortAnswerField()
	q.input = field
	return q
}

func newLongQuestion(question string) Question {
	q := newQuestion(question)
	field := NewLongAnswerField()
	q.input = field
	return q
}

func New(questions []Question) *model {
	styles := defaultStyles()
	//answerField := textinput.New()
	//answerField.Placeholder = "answer here"
	//answerField.Focus()
	return &model{questions: questions /*answerField: answerField,*/, styles: styles}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.questions[m.indx]
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+s":
			current.answer = current.input.Value()
			//m.answerField.SetValue("")
			log.Printf("question: %s, answer: %s", current.question, current.answer)
			m.Next()
			return m, current.input.Blur
		}

	}
	current.input, cmd = current.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	current := m.questions[m.indx]
	if m.width == 0 {
		return "loading..."
	}
	return lipgloss.JoinVertical(lipgloss.Center, m.questions[m.indx].question, m.styles.InputField.Render(current.input.View()))
}

func (m *model) Next() {
	if m.indx < len(m.questions)-1 {
		m.indx++
	} else {
		m.indx = 0
	}
}

func main() {

	questions := []Question{
		newShortQuestion("name: "),
		newShortQuestion("level: "),
		newShortQuestion("id: "),
		newShortQuestion("email: "),
		newLongQuestion("dummy answer: "),
	} //pull from repo or localized db?
	//more questions variables for different scenes?
	m := New(questions)

	f, err := tea.LogToFile("debug.log", "DBG: ")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
