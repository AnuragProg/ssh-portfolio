package notfound

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NotFound struct {
	renderer *lipgloss.Renderer
	height, width int
}

func NewNotFound(renderer *lipgloss.Renderer,height, width int) NotFound {
	return NotFound{renderer, height, width}
}

func (o NotFound) Init() tea.Cmd { return nil }


func (o NotFound) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return o, nil }


func (o NotFound) View() string { 
	return o.renderer.NewStyle().
		Foreground(color.Red).
		SetString("Not Found").
		String()
}
