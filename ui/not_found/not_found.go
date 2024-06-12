package notfound

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	"github.com/AnuragProg/ssh-portfolio/ui/model"
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

func (nf NotFound) Init() tea.Cmd { return nil }


func (nf NotFound) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return nf, nil }


func (nf NotFound) View() string { 
	return nf.renderer.NewStyle().
		Foreground(color.Red).
		SetString("Not Found").
		String()
}


func (nf NotFound)Resume() model.ResumableModel {
	return nf
}

func (nf NotFound)Pause() model.ResumableModel {
	return nf
}
