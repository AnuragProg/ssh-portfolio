package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NavBar struct {
	SelectedNavItemIdx int

	renderer *lipgloss.Renderer

	navItemStyle lipgloss.Style
}

func NewNavBar(renderer *lipgloss.Renderer) NavBar {
	return NavBar{
		SelectedNavItemIdx: 0,
		renderer:           renderer,
		navItemStyle: renderer.NewStyle().
			MaxHeight(5).
			Width(15).
			Align(lipgloss.Center),
	}
}

func (nb NavBar) Init() tea.Cmd {
	return nil
}

func (nb NavBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyShiftTab:
			nb.SelectedNavItemIdx = nb.SelectedNavItemIdx - 1
			if nb.SelectedNavItemIdx < 0 {
				nb.SelectedNavItemIdx = len(ContentPages) - 1
			}
		case tea.KeyTab:
			nb.SelectedNavItemIdx = (nb.SelectedNavItemIdx + 1) % len(ContentPages)
		}
	}

	return nb, nil
}

func (nb NavBar) View() string {
	navbar := []string{}
	for idx, contentPage := range ContentPages {
		if idx == nb.SelectedNavItemIdx {
			navbar = append(navbar, nb.navItemStyle.Copy().AlignHorizontal(lipgloss.Right).SetString("-> "+nb.renderer.NewStyle().Underline(true).Render(string(contentPage))).Render())
		} else {
			navbar = append(navbar, nb.navItemStyle.Copy().AlignHorizontal(lipgloss.Right).SetString(string(contentPage)).Render())
		}

		if idx != len(ContentPages)-1 {
			navbar = append(navbar, nb.renderer.NewStyle().Width(1).Render())
		}
	}
	return lipgloss.JoinVertical(lipgloss.Center, navbar...)
}
