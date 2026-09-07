package ui

import (
	"charm.land/lipgloss/v2"
)

// RenderSalesLayout arma la pantalla principal del POS combinando header y contenido
func RenderSalesLayout(username string, sessionID int, width, height int) string {
	// 1. Renderizamos el Header
	header := RenderHeader(username, sessionID)

	// 2. Cuerpo temporal de la vista de ventas
	bodyStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorSecondary).
		Width(width - 4).
		Height(height - lipgloss.Height(header) - 4) // Ajusta el alto disponible

	body := bodyStyle.Render("Aquí irá la tabla de productos del ticket y el total a pagar...")

	// 3. Unimos todo de forma vertical rellenando la pantalla
	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
	)
}
