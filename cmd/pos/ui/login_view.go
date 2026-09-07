package ui

import (
	"fmt"
	"posgo/internal/usecase"

	"charm.land/lipgloss/v2"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
	usernameInput textinput.Model
	passwordInput textinput.Model
	focusIndex    int
	errMessage    string
	sessionUC     *usecase.SessionUseCase

	width  int
	height int
	//callback o senal para avisarle al MainModel que el login fue exitoso
	LoggedIn      bool
	LoggedInUser  string
	LoggedInUUIDs string
	SessionID     int
}

func NewLoginModel(sessionUC *usecase.SessionUseCase) LoginModel {
	userInput := textinput.New()
	userInput.Placeholder = "Usuario"
	userInput.Focus()
	userInput.CharLimit = 20
	userInput.Width = 30

	passwordInput := textinput.New()
	passwordInput.Placeholder = "Contraseña..."
	passwordInput.EchoMode = textinput.EchoPassword // Oculta los caracteres como contraseña
	passwordInput.EchoCharacter = '•'
	passwordInput.CharLimit = 50
	passwordInput.Width = 30

	return LoginModel{
		usernameInput: userInput,
		passwordInput: passwordInput,
		focusIndex:    0,
		sessionUC:     sessionUC,
	}
}

func (m LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m LoginModel) Update(msg tea.Msg) (LoginModel, tea.Cmd) {
	var cmds [2]tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyDown, tea.KeyUp:
			if m.focusIndex == 0 {
				m.focusIndex = 1
				m.usernameInput.Blur()
				m.passwordInput.Focus()
			} else {
				m.focusIndex = 0
				m.passwordInput.Blur()
				m.usernameInput.Focus()
			}

			return m, nil
		case tea.KeyEnter:
			username := m.usernameInput.Value()
			if username == "" {
				m.errMessage = "El usuario no puede estar vacío"
				return m, nil
			}

			password := m.passwordInput.Value()
			if password == "" {
				m.errMessage = "La contraseña no puede estar vacía."
				return m, nil
			}

			// Llamamos al caso de uso y capturamos la sesión real devuelta
			session, err := m.sessionUC.Login(username, password)
			if err != nil {
				m.errMessage = fmt.Sprintf("Error: %v", err)
				return m, nil
			}

			// ¡Listo! Aquí usamos los datos reales que vinieron de la base de datos
			m.LoggedIn = true
			m.LoggedInUser = username
			m.LoggedInUUIDs = session.UUIDs
			m.SessionID = session.ID // Este ID ya viene correcto de SQLite
			return m, nil
		}
	}

	if m.focusIndex == 0 {
		m.usernameInput, cmds[0] = m.usernameInput.Update(msg)
	} else {
		m.passwordInput, cmds[1] = m.passwordInput.Update(msg)
	}

	return m, tea.Batch(cmds[0], cmds[1])
}

func (m LoginModel) View() string {
	title := TitleStyle.Render(" INICIAR SESIÓN - POSGO ")

	// Estilo para los campos
	uiBox := BoxStyle.Copy().Width(40).Render(
		fmt.Sprintf(
			"Usuario :\n%s\n\nContraseña :\n%s\n\n%s",
			m.usernameInput.View(),
			m.passwordInput.View(),
			ErrorStyle.Render(m.errMessage),
		),
	)

	help := HelpStyle.Render("\n[Tab] Cambiar campo  |  [Enter] Ingresar  |  [Ctrl+C] Salir")

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"\n",
		uiBox,
		"\n",
		help,
	)

	if m.width == 0 || m.height == 0 {
		return content
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)

}
