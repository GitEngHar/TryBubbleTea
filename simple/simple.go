package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

var (
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6b6d75"))
	keyStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffa500"))
	separator  = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4B5563")).
			PaddingLeft(2).
			PaddingRight(2).
			Render(" | ")
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
				fmt.Sprintf("%s %s", keyStyle.Render("[c]"), valueStyle.Render("ChangeCalcModec")),
			),
		),
	)
)

type model struct {
	input    textinput.Model
	editing  bool
	calcMode string
	count    int
}

func (m *model) Init() tea.Cmd {
	ti := textinput.New()
	ti.Placeholder = "PLUS or MINUS"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 30

	m.input = ti
	m.calcMode = "PLUS"
	return nil
}

// 状態の更新
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "c":
			m.editing = true
			m.input.SetValue(m.calcMode)
			return &m, textinput.Blink
		case "enter":
			if m.editing {
				m.calcMode = m.input.Value()
				m.editing = false
			}
		case "esc":
			m.editing = false
		case "q":
			return &m, tea.Quit
		case " ":
			if m.calcMode == "PLUS" {
				m.count++
			} else if m.calcMode == "MINUS" {
				m.count--
			}

		}
	}

	if m.editing {
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return &m, cmd
	}

	return &m, nil
}
func (m *model) View() string {
	currentRow := ""
	if m.editing {
		currentRow = lipgloss.JoinHorizontal(lipgloss.Left, titleStyle.Render("CalcMode:"), m.input.View())
	} else {
		currentRow = lipgloss.JoinHorizontal(lipgloss.Left, titleStyle.Render("CalcMode:"), valueStyle.Render(m.calcMode))
	}
	var infoBlock = lipgloss.JoinVertical(
		lipgloss.Left,
		currentRow,
	)
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		infoBlock,
		separator,
		appDetails,
	) + fmt.Sprintf("\nCount: %d", m.count)
}

func main() {
	// これなんだ??
	p := tea.NewProgram(&model{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Erorr", err)
		os.Exit(1)
	}
}
