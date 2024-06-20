package project

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type projectItem struct {
	title, desc string
}
func (ei projectItem) Title() string { return ei.title }
func (ei projectItem) Description() string { return ei.desc }
func (ei projectItem) FilterValue() string { return ei.title }

type Project struct {
	renderer *lipgloss.Renderer
	height, width int

	// data to be shown
	projects list.Model
}

func NewProject(renderer *lipgloss.Renderer, height, width int) Project {
	experiences := list.New(
		[]list.Item{
			projectItem{"Project 1", "Description"},
			projectItem{"Project 2", "Description"},
			projectItem{"Project 3", "Description"},
		},
		list.NewDefaultDelegate(),
		width, height,
	)

	experiences.SetShowHelp(false)
	experiences.SetShowTitle(false)
	experiences.SetShowStatusBar(false)

	return Project{
		renderer,
		height, width,
		experiences,
	}
}

func (e Project) Init() tea.Cmd {
	return nil
}
func (e Project) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	e.projects, cmd = e.projects.Update(msg)
	return e, cmd
}

func (e Project) View() string {
	return lipgloss.PlaceHorizontal(
		e.width,
		lipgloss.Left,
		e.projects.View(),
	)
}
