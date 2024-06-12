package overview

import (
	"github.com/AnuragProg/ssh-portfolio/ui/color"
	"github.com/AnuragProg/ssh-portfolio/ui/filler"
	"github.com/AnuragProg/ssh-portfolio/ui/model"
	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Overview struct {
	renderer      *lipgloss.Renderer
	height, width int
	cursorModel   cursor.Model
}

func NewOverview(renderer *lipgloss.Renderer, height, width int) Overview {
	cursorModel := cursor.New()
	cursorModel.SetMode(cursor.CursorBlink)
	cursorModel.SetChar(" ")
	cursorModel.Style = renderer.NewStyle()     //.SetString("")
	cursorModel.TextStyle = renderer.NewStyle() //.SetString("")
	cursorModel.Focus()
	return Overview{
		renderer,
		height, width,
		cursorModel,
	}
}

func (o Overview) Init() tea.Cmd {
	return nil
}

func (o Overview) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// switch msg.(type) {
	// case cursor.BlinkMsg:
	// 	log.Println("getting blink message")
	// }
	o.cursorModel, cmd = o.cursorModel.Update(msg)
	return o, cmd
}

func (o Overview) View() string {

	drawing := o.renderer.NewStyle().
		Foreground(color.LightGray).
		Render(`
   @@@@@@@@@@@@@@++@   @@@@@@@@@@@@ 
  @*-----------@+++-+@@===========#@
  @+::      .:-%#+++++@::.       :*@
   @@+.     -#@%*++++*%%.      .+@@ 
    @+.     -#%*++++#*.      .=@@@  
    @+.     -#%*++##.      .=%@@    
    @+.     -#%*##.       -#%*%@    
   @%+.     -#@%:       :*@#++==%   
 @+-#+.     -*:       .+@#++++++-+@ 
@++*#+.     -       .=@%+++++++++++@
 @@#%+.           +-#%+++++++++*%@@ 
   @@+.          =..%++++++++*#@@   
    @+.        -#@%%*%%%#*##%@@@@@  
    @+.      :%@@  @*@  .-. .-:  #@ 
    @+.    :#@##. %+*-:%%%:-%@*.-%  
    @+.  .*@@#*+ :#+% =@@# =@@: #@  
    @%++*@@@@@@+=@%%++@@%++@@#=+@@  
      @@@     @@#**%@@              
                @@@@                
	`)

	title := o.renderer.
		NewStyle().
		Foreground(color.White).
		Render("anurag singh - a software developer")

	desc := o.renderer.
		NewStyle().
		Foreground(color.LightGray).
		Width(lipgloss.Width(title)).
		Blink(true).
		Render("i'm currently a student and a software developer with interest in building something new every now and then. i use nvim btw..." + o.cursorModel.View())

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		drawing,
		filler.HorizontalFiller(o.renderer, 4),
		lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			filler.VerticalFiller(o.renderer, 2),
			desc,
		),
	)
}


func (o Overview) Resume() model.ResumableModel{
	o.cursorModel.Focus()
	return o
}
func (o Overview) Pause() model.ResumableModel {
	return o
}
