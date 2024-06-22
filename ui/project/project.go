package project

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tagItem struct {
	tag                string
	textColor, bgColor lipgloss.Color
}

type projectItem struct {
	title, desc, link string
	tags              []tagItem
}

type Project struct {
	renderer      *lipgloss.Renderer
	height, width int

	// data to be shown
	projects []projectItem
}

func NewProject(renderer *lipgloss.Renderer, height, width int) Project {
	return Project{
		renderer,
		height, width,
		[]projectItem{
			{
				"printit",
				"a microservices platfom to place orders for file printing.",
				"https://github.com/anuragprog/printit-microservices-monorepo",
				[]tagItem{
					{"golang", color.White, color.Blue},
					{"rust", color.White, color.DarkGray},
					{"python", color.White, color.Blue},
					{"js", color.Black, color.Yellow},
				},
			},
			{
				"job-ki-khoj",
				"an extension to PGRKAM govt website adding realtime functionalities.",
				"https://github.com/job-ki-khoj-smart-india-hackathon-2023/backend-monorepo",
				[]tagItem{
					{"js", color.Black, color.Yellow},
				},
			},
			{
				"gps-attendance-tracker",
				"an automated personal attendance tracker through gps location android app.",
				"https://github.com/anuragprog/gps-attendance-tracker",
				[]tagItem{
					{"kotlin", color.Black, color.Purple},
				},
			},
		},
	}
}

func (e Project) Init() tea.Cmd {
	return nil
}
func (e Project) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return e, nil
}

func (e Project) View() string {

	view := make([]string, 0, len(e.projects))

	titleStyle := e.renderer.NewStyle().
		Foreground(color.White)

	descStyle := e.renderer.NewStyle().
		Foreground(color.Gray)

	linkStyle := e.renderer.NewStyle().
		Underline(true).
		Foreground(color.LightBlue)


	for _, project := range e.projects {

		tagViews := make([]string, 0, len(project.tags))
		for _, tag := range project.tags {

			tagStyle := e.renderer.NewStyle().
				Background(tag.bgColor).
				Foreground(tag.textColor).
				Padding(0, 1)

			tagViews = append(
				tagViews,
				tagStyle.Render(tag.tag),
				" ",
			)
		}

		// title, desc, link string
		// tags []string
		// NOTE: to combine tags and title into one line
		titleAndTagViews := make([]string, 0, len(project.tags)+1)
		titleAndTagViews = append(titleAndTagViews, titleStyle.Render(project.title), " ")
		for _, tagView := range tagViews {
			titleAndTagViews = append(titleAndTagViews, tagView)
		}
		view = append(
			view,
			lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.JoinHorizontal(
					lipgloss.Center,
					titleAndTagViews...,
				),
				linkStyle.Render(project.link),
				descStyle.Render(project.desc),
			),
			"", // NOTE: for spacing b/w project cards
		)
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
