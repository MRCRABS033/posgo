package ui

import "charm.land/lipgloss/v2"

var (
	ColorPrimary    = lipgloss.Color("#4F46E5") // Índigo principal
	ColorSecondary  = lipgloss.Color("#818CF8") // Índigo claro para acentos
	ColorBackground = lipgloss.Color("#1F2937") // Gris oscuro / Fondo
	ColorText       = lipgloss.Color("#F9FAFB") // Blanco suave para texto
	ColorMuted      = lipgloss.Color("#9CA3AF") // Gris apagado para textos secundarios
	ColorSuccess    = lipgloss.Color("#10B981") // Verde para éxitos o totales
	ColorError      = lipgloss.Color("#EF4444") // Rojo para errores o alertas
)

var (
	// Estilo para títulos principales
	TitleStyle = lipgloss.NewStyle().
			Foreground(ColorText).
			Background(ColorPrimary).
			Bold(true).
			Padding(0, 2)

	// Estilo para cajas o contenedores con bordes redondeados
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorSecondary).
			Padding(1, 2)

	// Estilo para textos de ayuda o subtítulos
	HelpStyle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Italic(true)

	// Estilo para mensajes de error
	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	// Estilo para mensajes de éxito
	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true)
)
