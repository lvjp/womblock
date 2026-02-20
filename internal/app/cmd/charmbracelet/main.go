package charmbracelet

import (
	"fmt"

	"git.sr.ht/~lvjp/wtf-go/internal/pkg/cmd/util"
	"git.sr.ht/~lvjp/wtf-go/pkg/buildinfo"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Run(ctx *util.Context) error {
	p := tea.NewProgram(
		initialModel(),
		tea.WithInput(ctx.Input),
		tea.WithOutput(ctx.Output),
		tea.WithContext(ctx),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("failed to run the Charmbracelet TUI: %w", err)
	}

	return nil
}

type model struct {
	message string
}

func initialModel() model {
	return model{message: "Hello, wtf-go!"}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	msg := fmt.Sprintf("%s\n\n%s\n\nPress q to quit.", m.message, buildinfo.Get().String())

	// Send the UI for rendering
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Render(msg)
}
