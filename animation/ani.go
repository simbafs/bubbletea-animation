package animation

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Msg struct {
	Delay time.Duration
	jobs  []any
}

func (m *Msg) Push(job ...any) {
	m.jobs = append(m.jobs, job...)
}

func (m *Msg) Pop() any {
	if len(m.jobs) == 0 {
		return nil
	}

	job := m.jobs[len(m.jobs)-1]
	m.jobs = m.jobs[:len(m.jobs)-1]
	return job
}

func (m *Msg) Shift() any {
	if len(m.jobs) == 0 {
		return nil
	}

	job := m.jobs[0]
	m.jobs = m.jobs[1:]
	return job
}

func (m Msg) Cmd() tea.Cmd {
	if len(m.jobs) == 0 {
		return nil
	}
	return tea.Tick(m.Delay, func(time.Time) tea.Msg {
		return m
	})
}
