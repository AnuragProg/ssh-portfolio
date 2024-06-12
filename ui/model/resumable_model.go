package model

import tea "github.com/charmbracelet/bubbletea"


/** Mixture of tea.Model and resumable functions like pause and resume to allow control of state inside of those model on navigation **/
type ResumableModel interface {
	tea.Model
	Resume() ResumableModel
	Pause() ResumableModel
}
