package notfound

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NotFound struct {
	renderer *lipgloss.Renderer
}

func NewNotFound(renderer *lipgloss.Renderer) NotFound {
	return NotFound{renderer}
}

func (nf NotFound) Init() tea.Cmd { return nil }


func (nf NotFound) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return nf, nil }


func (nf NotFound) View() string { 
	return nf.renderer.NewStyle().
		Foreground(color.Red).
		SetString("Not Found").
		String()
}


