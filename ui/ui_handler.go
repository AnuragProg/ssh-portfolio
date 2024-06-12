package ui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
)

func UIHandler(sshSession ssh.Session) (tea.Model, []tea.ProgramOption) {

	pty, _, _ := sshSession.Pty()
	log.Println(pty.Term)

	renderer := bubbletea.MakeRenderer(sshSession)
	return NewUI(renderer, pty.Window.Height, pty.Window.Width), []tea.ProgramOption{tea.WithAltScreen()}
}
