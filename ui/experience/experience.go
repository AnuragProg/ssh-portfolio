package experience

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type experienceItem struct {
	companyName, role, from, to string
}

type Experience struct {
	renderer      *lipgloss.Renderer
	height, width int
	experiences   []experienceItem
}

func NewExperience(renderer *lipgloss.Renderer, height, width int) Experience {
	return Experience{
		renderer,
		height, width,
		[]experienceItem{
			{"quickghy", "android developer intern", "sep/2022", "feb/2023"},
			{"isostats", "full stack developer intern", "apr/2023", "aug/2023"},
			{"contineu.ai", "backend intern", "may/2023", "present"},
		},
	}
}

func (e Experience) Init() tea.Cmd {
	return nil
}
func (e Experience) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return e, nil
}

func (e Experience) View() string {
	view := make([]string, 0, len(e.experiences)*2) // INFO: * 2 so that space can be added between the experiences

	roleStyle := e.renderer.NewStyle().
		Foreground(color.White)
	fromToStyle := e.renderer.NewStyle().
		Foreground(color.Gray)
	companyNameStyle := e.renderer.NewStyle().
		Foreground(color.Gray)

	for idx, experience := range e.experiences {
		view = append(
			view,
			lipgloss.JoinVertical(
				lipgloss.Left,
				roleStyle.Render(experience.role),
				companyNameStyle.Render(experience.companyName),
				fromToStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, experience.from, ":", experience.to)),
			),
		)
		if idx != len(e.experiences) -1 {
			view = append(view, "")
		}
	}

	return lipgloss.Place(
		e.width,
		e.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			view...,
		),
	)
}
