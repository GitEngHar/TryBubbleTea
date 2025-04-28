package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

var (
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6b6d75"))
	keyStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffa500"))
	infoBlock  = lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Left, titleStyle.Render("Current:"), valueStyle.Render("/User/deletor/Download")),
		lipgloss.JoinHorizontal(lipgloss.Left, titleStyle.Render("Format:"), valueStyle.Render("txt")),
		lipgloss.JoinHorizontal(lipgloss.Left, titleStyle.Render("Size:"), valueStyle.Render("10MB")),
		lipgloss.JoinHorizontal(lipgloss.Left, titleStyle.Render("Exclude FileName:"), valueStyle.Render("backup")),
		lipgloss.JoinHorizontal(lipgloss.Left, titleStyle.Render("Show Hidden Files:"), valueStyle.Render("false")),
	)
	separator = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4B5563")).
			PaddingLeft(2).
			PaddingRight(2).
			Render("/")
	appDetails = lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Left,
			fmt.Sprintf("%s %s", keyStyle.Render("[Space]"), valueStyle.Render("Incremental")),
			fmt.Sprintf("%s %s", keyStyle.Render("[q]"), valueStyle.Render("quit")),
		),
		lipgloss.NewStyle().PaddingLeft(4).Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				fmt.Sprintf("%s %s", keyStyle.Render("[c]"), valueStyle.Render("ChangeMode")),
			),
		),
	)
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

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		infoBlock,
		separator,
		appDetails,
	) + fmt.Sprintf("\nCount: %d", m.count)
}

func main() {
	// これなんだ??
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Erorr", err)
		os.Exit(1)
	}
}
