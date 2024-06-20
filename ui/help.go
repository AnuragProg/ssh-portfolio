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

	Up   key.Binding
	Down key.Binding
	Quit key.Binding
	Help key.Binding

	ContentPageUp   key.Binding
	ContentPageDown key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Help, k.Quit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.ContentPageUp, k.ContentPageDown},
		{k.Help, k.Quit},
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
			Up: key.NewBinding(
				key.WithKeys("shift+tab"),
				key.WithHelp("shift+tab", "go up"),
			),
			Down: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "go down"),
			),
			Quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
			Help: key.NewBinding(
				key.WithKeys("?"),
				key.WithHelp("?", "toggle help"),
			),
			ContentPageUp: key.NewBinding(
				key.WithKeys("↑"),
				key.WithHelp("↑/k", "scroll content up"),
			),
			ContentPageDown: key.NewBinding(
				key.WithKeys("↓"),
				key.WithHelp("↓/j", "scroll content down"),
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
