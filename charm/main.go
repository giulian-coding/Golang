/*
package main

// A simple program that opens the alternate screen buffer then counts down
// from 5 and then exits.

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model int

type tickMsg time.Time

func main() {
	p := tea.NewProgram(model(10), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		m--
		if m <= 0 {
			return m, tea.Quit
		}
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("\n\n     Hi. This program will exit in %d seconds...", m)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
*/

package main

// A simple program that handled losing and acquiring focus.

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model{
		// assume we start focused...
		focused:   true,
		reporting: true,
	}, tea.WithReportFocus())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	focused   bool
	reporting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.FocusMsg:
		m.focused = true
	case tea.BlurMsg:
		m.focused = false
	case tea.KeyMsg:
		switch msg.String() {
		case "t":
			m.reporting = !m.reporting
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Hi. Focus report is currently "
	if m.reporting {
		s += "enabled"
	} else {
		s += "disabled"
	}
	s += ".\n\n"

	if m.reporting {
		if m.focused {
			s += "This program is currently focused!"
		} else {
			s += "This program is currently blurred!"
		}
	}
	return s + "\n\nTo quit sooner press ctrl-c, or t to toggle focus reporting...\n"
}
