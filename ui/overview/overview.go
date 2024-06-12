package overview

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type Overview struct {
	renderer *lipgloss.Renderer
	height, width int
}

func NewOverview(renderer *lipgloss.Renderer,height, width int) Overview {
	return Overview{renderer, height, width}
}

func (o Overview) Init() tea.Cmd { return nil }

func (o Overview) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return o, nil }

func (o Overview) View() string { 

	title := o.renderer.
		NewStyle().
		Foreground(color.White).
		Render("Anurag Singh - Software Developer")

	desc := o.renderer.
		NewStyle().
		Foreground(color.LightGray).
		Width(lipgloss.Width(title)).
		Render("I'm currently a student and a software developer with interest in building something new every now and then. I use NVIM BTW...")

	return lipgloss.JoinVertical(lipgloss.Center, title, desc)
}
