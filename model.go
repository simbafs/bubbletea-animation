package main

import (
	"log"
	"time"

	"anima/animation"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	str string
}

func (m *Model) Init() tea.Cmd {
	log.Printf("Init")

	ani := animation.Msg{
		Delay: 100 * time.Millisecond,
	}

	return DeleteChar(Typewriter(ani, "Hello world"), 12).Cmd()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	log.Printf("%#v", msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case animation.Msg:
		log.Printf("Msg: %s", msg)
		switch job := msg.Shift().(type) {
		case Push:
			log.Printf("Push: %s", job)
			m.str += string(job)
			cmd = msg.Cmd()
		case Pop:
			if m.str != "" {
				m.str = m.str[:len(m.str)-int(job)]
			}
			cmd = msg.Cmd()
		case Set:
			m.str = string(job)
			cmd = msg.Cmd()
		default:
			msg.Push(job)
			cmd = msg.Cmd()
		}
	}

	return m, cmd
}

func (m *Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		"Hello bubbletea!",
		m.str,
		"",
	)
}

func NewModel() *Model {
	return &Model{}
}

type (
	Push string
	Pop  int
	Set  string
)

func Cmd[T any](t T) tea.Cmd {
	return func() tea.Msg {
		return t
	}
}

func Typewriter(ani animation.Msg, str string) animation.Msg {
	for _, r := range str {
		ani.Push(Push(string(r)))
	}

	return ani
}

func DeleteChar(ani animation.Msg, n int) animation.Msg {
	for i := 0; i < n; i++ {
		ani.Push(Pop(1))
	}

	return ani
}
