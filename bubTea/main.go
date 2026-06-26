package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/zyad-elkhewekh/go-tutorial/bubTea/internals/handlers"

	//for api
	tea "github.com/charmbracelet/bubbletea"
	//"github.com/zyad-elkhewekh/go-tutorial/bubTea/internals/tools"
)

type screen int

const (
	namescreen screen = iota
	firstScreen
	secondScreen
	thirdScreen
)

//the program (model) is split 3 way: init - update - view

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	screen   screen
	username string
}

// this should be a func to return init state of model
// but can be init as var somewhere else
func initialModel() model {
	return model{
		choices: []string{"cpp", "c", "java", "python"},

		selected: make(map[int]struct{}),
		screen:   namescreen,
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

		if m.screen == namescreen {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "enter":
				if m.username != "" {
					m.screen = firstScreen
				}
			case "backspace":
				if len(m.username) > 0 {
					m.username = m.username[:len(m.username)-1]
				}
			default:
				if len(msg.String()) == 1 {
					m.username += msg.String()
				}
			}
			return m, nil
		}
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
			case namescreen:
				switch msg.String() {
				case "enter":
					if m.username != "" {
						m.screen = firstScreen
					}
				case "backspace":
					if len(m.username) > 0 {
						m.username = m.username[:len(m.username)-1]
					}
				default:
					// only append printable single chars
					if len(msg.String()) == 1 {
						m.username += msg.String()
					}
				}
			case firstScreen:
				m.screen = secondScreen
				m.saveChoices()
				//m.selected = make(map[int]struct{}) // reset for screen 2
				m.cursor = 0
			case secondScreen:
				m.screen = thirdScreen
				m.saveChoices()
			}
		}
	}
	//return the updated model m
	return m, nil
}

func (m model) View() string {
	switch m.screen {
	case namescreen:
		return m.nameView()
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

func (m model) nameView() string {
	return fmt.Sprintf("Enter your username:\n\n> %s\n\nPress enter to continue.", m.username)
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

func (m model) saveChoices() {
	var picked []string
	for i := range m.selected {
		picked = append(picked, m.choices[i])
	}

	body, _ := json.Marshal(map[string]string{
		"username": m.username,
		"choice":   strings.Join(picked, ","),
	})

	resp, err := http.Post(
		"http://localhost:8080/choice",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	p := tea.NewProgram(initialModel())

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	go func() {
		fmt.Println("starting my first api..")
		err := http.ListenAndServe("localhost:8080", r)
		if err != nil {
			log.Error(err)
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
