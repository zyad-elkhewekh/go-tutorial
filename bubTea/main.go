package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	firstScreen screen = iota
	secondScreen
	thirdScreen
)

//the program (model) is split 3 way: init - update - view

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	screen   screen
}

// this should be a func to return init state of model
// but can be init as var somewhere else
func initialModel() model {
	return model{
		choices: []string{"cpp", "c", "java", "python"},

		selected: make(map[int]struct{}),
		screen:   firstScreen,
	}
}

// init will be used for i/o
func (m model) Init() tea.Cmd {
	return nil //havent figured tht out yet
}

// update func is to tell trhat something happend
// and return action (in this case an updated model)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//a switch case on the var msg that'll carry
	//some values of pressed keys
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		//space to choose (toggle the state of choice)
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "enter":
			switch m.screen {
			case firstScreen:
				m.screen = secondScreen
			case secondScreen:
				m.screen = thirdScreen
			}
		}
	}
	//return the updated model m
	return m, nil
}

func (m model) View() string {
	switch m.screen {
	case firstScreen:
		return m.firstView()
	case secondScreen:
		return m.secView()
	case thirdScreen:
		return m.thrdView()
	default:
		return ""
	}
}

func (m model) firstView() string {
	s := "what are your favourite languages out of these? (press space to pick and enter to confirm)\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">" //this is the actual cursor visually
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress q to quit.\n"
	return s
}
func (m model) secView() string {
	s := "what are your most frequently used languages out of these? (ignore likability)\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">" //this is the actual cursor visually
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress q to quit.\n"
	return s
}

func (m model) thrdView() string {
	s := "you\n\n"
	var lifeIndx int = 0
	var mastery int = 0
	for i := range m.selected {
		switch m.choices[i] {
		case "python":
			mastery--
			lifeIndx++
		case "cpp":
			mastery++
			lifeIndx--
		case "c":
			mastery += 2
			lifeIndx -= 2
		case "java":
			mastery++
			lifeIndx++
		}
	}

	if mastery >= 3 || lifeIndx < 2 {
		s += "\n need to touch grass\n"
	} else if mastery < 0 {
		s += "need to learn how to code!\n"
	} else {
		s += "\nwitch!\n"
	}
	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
