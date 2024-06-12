package filler

import "github.com/charmbracelet/lipgloss"

func HorizontalFiller(renderer *lipgloss.Renderer, width int) string {
	return renderer.NewStyle().Width(width).String()
}
func VerticalFiller(renderer *lipgloss.Renderer, height int) string {
	return renderer.NewStyle().Height(height).String()
}
