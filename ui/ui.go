package ui

import (
	"github.com/AnuragProg/ssh-portfolio/ui/filler"
	notfound "github.com/AnuragProg/ssh-portfolio/ui/not_found"
	"github.com/AnuragProg/ssh-portfolio/ui/overview"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	ContentPageHeight = 25
	ContentPageWidth  = 70
)

type ContentPage string

const (
	Overview   = "Overview"
	Experience = "Experience"
	Projects   = "Projects"
	Contact    = "Contact"
)

var ContentPages = [4]ContentPage{Overview, Experience, Projects, Contact}


type UI struct {
	/** BELOW SECTION IS REFERENCES TO STUFF WE NEED **/
	tea.Model

	renderer *lipgloss.Renderer
	ContentPageMap  map[ContentPage]tea.Model // for routing between pages 

	/** BELOW SECTION IS ABOUT THE UI**/
	height int
	width  int

	// header section
	navbar tea.Model

	// content section

	// footer section
	help tea.Model
}

func NewUI(renderer *lipgloss.Renderer, height, width int) UI {
	return UI{

		renderer: renderer,
		ContentPageMap: map[ContentPage]tea.Model{
			Overview: overview.NewOverview(renderer, ContentPageHeight, ContentPageWidth),
		},

		height: height,
		width: width,

		navbar: NewNavBar(renderer),
		help:   NewHelpMenu(renderer),
	}
}

func (ui UI) Init() tea.Cmd {
	return nil
}

func (ui UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		ui.height = msg.Height
		ui.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), "q":
			return ui, tea.Quit
		}
	}

	ui.navbar, cmd = ui.navbar.Update(msg)
	cmds = append(cmds, cmd)
	ui.help, cmd = ui.help.Update(msg)
	cmds = append(cmds, cmd)

	return ui, tea.Batch(cmds...)
}

func (ui UI) View() string {

	header := ui.navbar.View()

	contentPage, ok := ui.ContentPageMap[ContentPages[ui.navbar.(NavBar).SelectedNavItemIdx]]
	if !ok {
		contentPage = notfound.NewNotFound(ui.renderer, ContentPageHeight, ContentPageWidth)
	}


	content := ui.renderer.NewStyle().
		Height(ContentPageHeight).
		Width(ContentPageWidth).
		Border(lipgloss.RoundedBorder()).
		Render(lipgloss.Place(ContentPageWidth, ContentPageHeight, lipgloss.Center, lipgloss.Center, contentPage.View()))

	footer := ui.help.View()

	screen := lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			header,
			filler.HorizontalFiller(ui.renderer, 3),
			content,
		), 
		filler.VerticalFiller(ui.renderer, 1),
		footer,
	)

	return lipgloss.Place(ui.width, ui.height, lipgloss.Center, lipgloss.Center, screen)
}
