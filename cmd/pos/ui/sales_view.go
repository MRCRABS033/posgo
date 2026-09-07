package ui

import (
	"fmt"
	"posgo/internal/domain"
	"posgo/internal/usecase"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SaleModel struct {
	barcodeInput textinput.Model
	focusIndex   int
	errMessage   string
	productUC    *usecase.ProductUseCase
	itemList     []*domain.Product
	table        table.Model
	width        int
	height       int
}

// Estilo de borde que abarcara todo el ancho disponible
var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func NewSaleModel(productUC *usecase.ProductUseCase) *SaleModel {
	barcodeInput := textinput.New()
	barcodeInput.Placeholder = "Escriba el codigo del producto..."
	barcodeInput.Focus()
	barcodeInput.Width = 50

	// Definimos columnas iniciales (se adaptarán al ancho de la terminal)
	columns := []table.Column{
		{Title: "ID", Width: 6},
		{Title: "Producto", Width: 30},
		{Title: "Stock", Width: 8},
		{Title: "Cantidad", Width: 10},
		{Title: "Precio", Width: 12},
		{Title: "Subtotal", Width: 12},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}),
		table.WithFocused(false),
	)

	s := table.DefaultStyles()
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(s)

	return &SaleModel{
		barcodeInput: barcodeInput,
		table:        t,
		focusIndex:   0,
		productUC:    productUC,
	}
}

func (m SaleModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *SaleModel) Update(msg tea.Msg) (*SaleModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// 1. ANCHO TOTAL: La tabla abarca el 100% de la pantalla horizontal
		tableWidth := m.width
		if tableWidth < 40 {
			tableWidth = 40
		}

		// Distribuimos el ancho dinámicamente entre las columnas fijas y hacemos que la del "Producto" crezca
		fixedWidth := 6 + 8 + 10 + 12 + 12          // ID + Stock + Cantidad + Precio + Subtotal
		productWidth := tableWidth - fixedWidth - 4 // -4 por los bordes de la tabla
		if productWidth < 15 {
			productWidth = 15
		}

		columns := []table.Column{
			{Title: "ID", Width: 6},
			{Title: "Producto", Width: productWidth},
			{Title: "Stock", Width: 8},
			{Title: "Cantidad", Width: 10},
			{Title: "Precio", Width: 12},
			{Title: "Subtotal", Width: 12},
		}
		m.table.SetColumns(columns)
		m.table.SetWidth(tableWidth)

		// 2. ALTO TOTAL: Calculamos el espacio vertical restante para que la tabla crezca hacia abajo.
		// El MainModel tiene el topNav (~2 líneas), más los títulos, input y ayudas (~9 líneas en total).
		reservedHeight := 11
		availableHeight := m.height - reservedHeight
		if availableHeight < 5 {
			availableHeight = 5
		}

		// Descontamos las líneas del header de la tabla (2 líneas)
		tableHeaderHeight := 2
		m.table.SetHeight(availableHeight - tableHeaderHeight)

		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab:
			if m.focusIndex == 0 {
				m.focusIndex = 1
				m.barcodeInput.Blur()
				m.table.Focus()
			} else {
				m.focusIndex = 0
				m.barcodeInput.Focus()
				m.table.Blur()
			}
			return m, nil

		case tea.KeyEnter:
			product := m.barcodeInput.Value()
			if product == "" {
				m.errMessage = "Escriba el codigo del producto por favor."
				return m, nil
			}
			m.errMessage = ""

			result, err := m.productUC.GetProductByCode(product)
			if err != nil {
				m.errMessage = fmt.Sprintf("Error: %v", err)
				return m, nil
			}

			m.itemList = append(m.itemList, result)
			m.barcodeInput.SetValue("")

			var rows []table.Row
			for _, p := range m.itemList {
				rows = append(rows, table.Row{
					fmt.Sprintf("%d", p.ID),
					p.ProductName,
					"10", // Stock de ejemplo
					"1",  // Cantidad por defecto
					fmt.Sprintf("$%.2f", p.UnitSellPrice),
					fmt.Sprintf("$%.2f", p.UnitSellPrice),
				})
			}
			m.table.SetRows(rows)
			return m, nil
		}
	}

	if m.focusIndex == 0 {
		m.barcodeInput, cmd = m.barcodeInput.Update(msg)
	} else {
		m.table, cmd = m.table.Update(msg)
	}

	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *SaleModel) View() string {
	s := "=== MÓDULO DE VENTAS (CAJA) ===\n\n"

	s += "Código de barras:\n"
	s += m.barcodeInput.View() + "\n\n"

	if m.errMessage != "" {
		errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
		s += errStyle.Render(m.errMessage) + "\n\n"
	}

	s += "Ticket de Venta:\n"

	// Forzamos al borde de la tabla a medir exactamente el ancho completo de la terminal
	containerWidth := m.width
	if containerWidth < 40 {
		containerWidth = 40
	}
	fullWidthStyle := baseStyle.Copy().Width(containerWidth)

	s += fullWidthStyle.Render(m.table.View()) + "\n\n"

	if m.focusIndex == 1 {
		s += lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("[Usa las flechas ↑/↓ para navegar en la tabla | Tab para volver al input]")
	} else {
		s += lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("[Presiona Tab para moverte a la tabla]")
	}

	return s
}
