package contact

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Contact struct {
	renderer      *lipgloss.Renderer
	height, width int
}

func NewContact(renderer *lipgloss.Renderer, height, width int) Contact {
	return Contact{
		renderer,
		height, width,
	}
}

func (c Contact) Init() tea.Cmd {
	return nil
}

func (c Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c Contact) View() string {
	title := `
███████╗ █████╗ ██╗   ██╗    ██╗  ██╗ ██████╗ ██╗      █████╗ 
██╔════╝██╔══██╗╚██╗ ██╔╝    ██║  ██║██╔═══██╗██║     ██╔══██╗
███████╗███████║ ╚████╔╝     ███████║██║   ██║██║     ███████║
╚════██║██╔══██║  ╚██╔╝      ██╔══██║██║   ██║██║     ██╔══██║
███████║██║  ██║   ██║       ██║  ██║╚██████╔╝███████╗██║  ██║
╚══════╝╚═╝  ╚═╝   ╚═╝       ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝
	`
	return title
}
