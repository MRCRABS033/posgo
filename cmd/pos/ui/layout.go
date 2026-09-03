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

// RenderLoginLayout arma la pantalla de inicio de sesión centrada
func RenderLoginLayout(width, height int) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorPrimary).
		Padding(2, 4)

	loginBox := boxStyle.Render("=== INICIAR SESIÓN ===")

	// Si tenemos las dimensiones de la pantalla, podemos centrar la caja
	if width > 0 && height > 0 {
		return lipgloss.Place(
			width, height,
			lipgloss.Center, lipgloss.Center,
			loginBox,
		)
	}

	return loginBox
}
