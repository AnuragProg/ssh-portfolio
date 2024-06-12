package ui

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	help.KeyMap

	Left  key.Binding
	Right key.Binding
	Quit  key.Binding
	Help  key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Help, k.Quit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Left, k.Right},
		{k.Quit, k.Help},
	}
}

type HelpMenu struct {

	renderer *lipgloss.Renderer

	keys KeyMap
	help help.Model
}

func NewHelpMenu(renderer *lipgloss.Renderer) HelpMenu {
	return HelpMenu{
		renderer: renderer,
		keys: KeyMap{
			Left: key.NewBinding(
				key.WithKeys("shift+tab"),
				key.WithHelp("shift+tab", "go left"),
			),
			Right: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "go right"),
			),
			Quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
			Help: key.NewBinding(
				key.WithKeys("?"),
				key.WithHelp("?", "toggle help"),
			),
		},
		help: help.New(),
	}
}

func (hf HelpMenu) Init() tea.Cmd {
	return nil
}
func (hf HelpMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		hf.help.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "?":
			hf.help.ShowAll = !hf.help.ShowAll
		}
	}

	return hf, nil
}
func (hf HelpMenu) View() string {
	hf.help.Styles.ShortDesc = hf.renderer.NewStyle().Foreground(color.Gray)
	hf.help.Styles.FullDesc = hf.renderer.NewStyle().Foreground(color.Gray)
	return hf.help.View(hf.keys)
}
