package ui

import (
	"posgo/internal/usecase"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define los estados o pantallas de tu POS
type ScreenState int

const (
	StateLogin ScreenState = iota
	StateSales
	StateCredits
	StateClients
	StateProducts
	StateInventory
	StateOrders
	StateSettings
	StateInvoice
	StateReports
)

type tabBoundary struct {
	start int
	end   int
	tabID int
	state ScreenState
}

// MainModel es el modelo raíz que maneja toda la aplicación
type MainModel struct {
	state      ScreenState
	width      int
	height     int
	loginModel LoginModel // Instancia del modelo de login
	saleModel  *SaleModel
	// Referencias a los casos de uso para pasárselos a las vistas
	sessionUC *usecase.SessionUseCase
	productUC *usecase.ProductUseCase
	activeTab int
	// Datos de sesión activa tras un login exitoso
	username  string
	sessionID int
}

func NewMainModel(sessionUC *usecase.SessionUseCase, productUC *usecase.ProductUseCase) MainModel {
	return MainModel{
		state:      StateLogin,
		loginModel: NewLoginModel(sessionUC), // Inicializamos el login
		saleModel:  NewSaleModel(productUC),
		sessionUC:  sessionUC,
		productUC:  productUC,
	}
}

func (m MainModel) Init() tea.Cmd {
	// Inicializamos el parpadeo del cursor del login al arrancar
	return m.loginModel.Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		var cmd tea.Cmd
		switch m.state {
		case StateLogin:
			m.loginModel, cmd = m.loginModel.Update(msg)
		case StateSales:
			m.saleModel, cmd = m.saleModel.Update(msg)
		}
		return m, cmd
	case tea.MouseMsg:
		if m.state != StateLogin && msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft && msg.Y == 0 {
			// Obtenemos los límites dinámicos de las pestañas
			boundaries := m.getTabBoundaries()

			for _, b := range boundaries {
				// Verificamos si el clic (msg.X) cayó dentro del rango de esta pestaña
				if msg.X >= b.start && msg.X <= b.end {
					m.activeTab = b.tabID
					m.state = b.state

					// Si cambiamos al módulo de ventas, actualizamos su tamaño
					if b.state == StateSales && m.width > 0 && m.height > 0 {
						m.saleModel.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
					}
					break
				}
			}
			return m, nil
		}

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		if m.state != StateLogin {
			switch msg.String() {
			case "f1":
				m.activeTab = 0
				m.state = StateSales

				// ¡CLAVE! Si ya tenemos el tamaño de la pantalla, se lo mandamos al saleModel al presionar F1
				if m.width > 0 && m.height > 0 {
					m.saleModel.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
				}
				return m, nil

			case "f2":
				m.activeTab = 1
				m.state = StateInventory
				return m, nil

			case "f3":
				m.activeTab = 2
				m.state = StateReports
				return m, nil

			case "f4":
				m.activeTab = 3
				m.state = StateProducts
				return m, nil
			}
		}
	}

	// Delegamos los eventos según la pantalla activa
	switch m.state {
	case StateLogin:
		var cmd tea.Cmd
		m.loginModel, cmd = m.loginModel.Update(msg)

		// Si el login fue exitoso, capturamos los datos y cambiamos al módulo de ventas
		if m.loginModel.LoggedIn {
			m.username = m.loginModel.LoggedInUser
			m.sessionID = m.loginModel.SessionID
			m.state = StateSales
			m.activeTab = 0

			// ¡CLAVE! Al pasar del login a ventas por primera vez, le inyectamos
			// el ancho y alto actual para que la tabla abarque toda la pantalla de inmediato.
			if m.width > 0 && m.height > 0 {
				m.saleModel.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
			}
		}
		return m, cmd

	case StateSales:
		var cmd tea.Cmd
		m.saleModel, cmd = m.saleModel.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *MainModel) stateStateForTab(tab int) {
	switch tab {
	case 0:
		m.state = StateSales
	case 1:
		m.state = StateCredits
	case 2:
		m.state = StateProducts
	case 3:
		m.state = StateInventory
	case 4:
		m.state = StateClients
	case 5:
		m.state = StateOrders
	case 6:
		m.state = StateInvoice // O tu estado de configuración
	case 7:
		m.state = StateInvoice
	case 8:
		m.state = StateReports
	case 9:
		m.state = StateReports
	}
}

func (m MainModel) getTabBoundaries() []tabBoundary {
	// Usamos el mismo estilo base (padding 0, 2) para medir el ancho exacto de cada pestaña
	style := lipgloss.NewStyle().Padding(0, 2)

	tabs := []struct {
		text  string
		tabID int
		state ScreenState
	}{
		{"[F1] Ventas", 0, StateSales},
		{"Clientes", 4, StateClients},
		{"[F2] Creditos", 1, StateCredits},
		{"[F3] Producto", 2, StateProducts},
		{"[F4] Inventario", 3, StateInventory},
		{"Compras", 5, StateOrders},
		{"Configuracion", 6, StateSettings},
		{"Facturas", 7, StateInvoice},
		{"Corte", 8, StateReports},
		{"Reportes", 9, StateReports},
	}

	var boundaries []tabBoundary
	currentX := 0

	for _, t := range tabs {
		// Medimos el ancho que ocupa la pestaña renderizada
		w := lipgloss.Width(style.Render(t.text))

		boundaries = append(boundaries, tabBoundary{
			start: currentX,
			end:   currentX + w,
			tabID: t.tabID,
			state: t.state,
		})

		// Avanzamos X sumando el ancho de la pestaña + 1 espacio separador (" ")
		currentX += w + 1
	}

	return boundaries
}

func (m MainModel) View() string {
	// Dependiendo del estado, renderizamos la pantalla correspondiente
	if m.state == StateLogin {
		return m.loginModel.View()
	}

	topNav := m.renderTopNav()

	var moduleContent string
	switch m.state {
	case StateSales:
		moduleContent = m.saleModel.View()
	case StateInventory:
		moduleContent = "Modulo de Inventario en desarrollo..."
	case StateReports:
		moduleContent = "Modulo de reportes en desarrollo..."
	case StateProducts:
		moduleContent = "Moddulo de productos en desarrollo..."
	case StateClients:
		moduleContent = "Modulo de clientes en desarrollo..."
	case StateOrders:
		moduleContent = "Modulo de ordenes de compra en desarrollo..."
	case StateSettings:
		moduleContent = "Modulo de configuracion en desarrollo..."
	case StateInvoice:
		moduleContent = "Modulo de facturas en desarrollo..."

	default:
		moduleContent = "Cargando..."
	}

	return lipgloss.JoinVertical(lipgloss.Left, topNav, moduleContent)
}
