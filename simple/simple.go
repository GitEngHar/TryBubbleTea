package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	count int
}

// 初期化
func (m model) Init() tea.Cmd {
	return nil
}

// 状態の更新
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case " ":
			m.count++
		}
	}
	return m, nil
}
func (m model) View() string {
	return fmt.Sprintf("Count: %d\n[Space] increment\n[q] quit\n", m.count)
}

func main() {
	// これなんだ??
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Erorr", err)
		os.Exit(1)
	}
}
