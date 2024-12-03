package main

import tea "github.com/charmbracelet/bubbletea"

func main() {
	tea.LogToFile("log.txt", "")

	m := NewModel()

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
