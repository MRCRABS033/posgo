package ui

import "github.com/charmbracelet/lipgloss"

func (m MainModel) renderTopNav() string {
	inactiveStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Background(lipgloss.Color("235")).
		Padding(0, 2)

	activeStile := lipgloss.NewStyle().
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("62")).
		Bold(true).
		Padding(0, 2)

	f1Tab := inactiveStyle.Render("[F1] Ventas")
	f2Tab := inactiveStyle.Render("[F2] Creditos")
	Tab1 := inactiveStyle.Render("Clientes")
	f3Tab := inactiveStyle.Render("[F3] Producto")
	f4Tab := inactiveStyle.Render("[F4] Inventario")
	//Tab2 := inactiveStyle.Render("Compras")
	//Tab3 := inactiveStyle.Render("Configuracion")
	//Tab4 := inactiveStyle.Render("Facturas")
	//Tab5 := inactiveStyle.Render("Corte")
	//Tab6 := inactiveStyle.Render("Reportes")
	switch m.activeTab {
	case 0:
		f1Tab = activeStile.Render("[F1] Ventas")
	case 1:
		f2Tab = activeStile.Render("[F2] Creditos")
	case 2:
		f3Tab = activeStile.Render("[F3] Productos")
	case 3:
		f4Tab = activeStile.Render("[F4] Inventario")
	case 4:
		Tab1 = activeStile.Render("Clientes")

	}

	navBar := lipgloss.JoinHorizontal(lipgloss.Top, f1Tab, " ", Tab1, " ", f2Tab, " ", f3Tab, " ", f4Tab)

	barStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("240")).
		Width(80)

	return barStyle.Render(navBar)
}
