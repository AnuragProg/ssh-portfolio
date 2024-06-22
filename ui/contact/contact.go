package contact

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SocialLink struct {
	tag, link string
}

// INFO: The tags must be equal in length to make the ui not look weird
var socialLinks = []SocialLink{
	{"github  ", "https://github.com/AnuragProg"},	
	{"linkedin", "https://www.linkedin.com/in/anurag-singh-bisht"},	
}

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
	

	tagStyle := c.renderer.NewStyle()
	linkStyle := c.renderer.NewStyle().
		Foreground(color.LightBlue).
		Underline(true).
		PaddingTop(1)

	links := []string{}
	for _, socialLink := range socialLinks {
		tag := socialLink.tag
		link := socialLink.link
		links = append(
			links,
			lipgloss.JoinHorizontal(lipgloss.Center, tagStyle.Render(tag), " - ", linkStyle.Render(link)),
		)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		links...,
	)
	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.Place(
			c.width,
			lipgloss.Height(title),
			lipgloss.Center,
			lipgloss.Center,
			title,
		),
		content,
	)
}
