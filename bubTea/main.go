package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

//the program (model) is split 3 way: init - update - view

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

// this should be a func to return init state of model
// but can be init as var somewhere else
func initialModel() model {
	return model{
		choices: []string{"cpp", "c", "java", "python"},

		selected: make(map[int]struct{}),
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
		//enter or space to choose (toggle the state of choice)
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	//return the updated model m
	return m, nil
}

func (m model) View() string {
	s := "what are your favourite languages out of these?\n\n"

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

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
