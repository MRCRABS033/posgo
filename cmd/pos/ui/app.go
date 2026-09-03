package ui

import (
	"posgo/internal/usecase"

	tea "github.com/charmbracelet/bubbletea"
)

// Define los estados o pantallas de tu POS
type ScreenState int

const (
	StateLogin ScreenState = iota
	StateSales
)

// MainModel es el modelo raíz que maneja toda la aplicación
type MainModel struct {
	state      ScreenState
	width      int
	height     int
	loginModel LoginModel // Instancia del modelo de login

	// Referencias a los casos de uso para pasárselos a las vistas
	sessionUC *usecase.SessionUseCase
	productUC *usecase.ProductUseCase

	// Datos de sesión activa tras un login exitoso
	username  string
	sessionID int
}

func NewMainModel(sessionUC *usecase.SessionUseCase, productUC *usecase.ProductUseCase) MainModel {
	return MainModel{
		state:      StateLogin,
		loginModel: NewLoginModel(sessionUC), // Inicializamos el login
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
		// Guardamos el ancho y alto actual de la terminal
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
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
		}
		return m, cmd

	case StateSales:
		// Aquí agregarás la lógica de actualización para el módulo de ventas más adelante
	}

	return m, nil
}

func (m MainModel) View() string {
	// Dependiendo del estado, renderizamos la pantalla correspondiente
	switch m.state {
	case StateLogin:
		return m.loginModel.View()

	case StateSales:
		// Contenido temporal para el módulo de ventas envuelto en el layout
		//salesContent := "¡Bienvenido al POS!\n[1] Registrar venta\n[2] Consultar producto"
		return RenderSalesLayout(m.username, m.sessionID, m.width, m.height)
	}

	return "Cargando..."
}
