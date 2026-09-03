package ui

import (
	"fmt"
	"time"

	"charm.land/lipgloss/v2"
)

func RenderHeader(username string, sessionID int) string {
	title := TitleStyle.Render(" POSGO - SISTEMA DE VENTAS ")

	currentTime := time.Now().Format("02/01/2006 15:04")
	infoText := fmt.Sprintf(" Cajero: %s  |  Sesión: #%d  |  %s ", username, sessionID, currentTime)

	infoStyle := lipgloss.NewStyle().
		Foreground(ColorText).
		Background(lipgloss.Color("#3730A3")).
		Padding(0, 2)

	info := infoStyle.Render(infoText)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		info,
	)
}
