package experience

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type experienceItem struct {
	title, desc string
}
func (ei experienceItem) Title() string { return ei.title }
func (ei experienceItem) Description() string { return ei.desc }
func (ei experienceItem) FilterValue() string { return ei.title }

type Experience struct {
	renderer *lipgloss.Renderer
	height, width int

	// data to be shown
	experiences list.Model
}

func NewExperience(renderer *lipgloss.Renderer, height, width int) Experience {
	experiences := list.New(
		[]list.Item{
			experienceItem{"QuickGhy", "Android Developer Intern"},
			experienceItem{"ISOStats", "Full Stack Developer Intern"},
			experienceItem{"Contineu AI", "Full Stack Developer Intern"},
		},
		list.NewDefaultDelegate(),
		width, height,
	)

	experiences.SetShowHelp(false)
	experiences.SetShowTitle(false)
	experiences.SetShowStatusBar(false)

	return Experience{
		renderer,
		height, width,
		experiences,
	}
}

func (e Experience) Init() tea.Cmd {
	return nil
}
func (e Experience) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	e.experiences, cmd = e.experiences.Update(msg)
	return e, cmd
}

func (e Experience) View() string {
	return lipgloss.PlaceHorizontal(
		e.width,
		lipgloss.Left,
		e.experiences.View(),
	)
}
