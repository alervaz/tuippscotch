package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type verb string

var verbs = []verb{GET, POST, PUT, PATCH, DELETE}

const (
	GET    verb = http.MethodGet
	POST   verb = http.MethodPost
	PUT    verb = http.MethodPut
	PATCH  verb = http.MethodPatch
	DELETE verb = http.MethodDelete
)

type inputStyle int

const (
	URL inputStyle = iota
	BODY
	NONE
	RESPONSE
)

type model struct {
	width             int
	height            int
	currentMode       int
	currentVerb       int
	currentInput      Input
	currentInputStyle inputStyle
	urlInput          *UrlInput
	bodyInput         *BodyInput
	styles            *styles
	response          *ResponseInput
}

func initModel() *model {
	return &model{
		urlInput:          newUrlInput(),
		bodyInput:         newBodyInput(),
		styles:            getStyles(),
		currentInputStyle: NONE,
		response:          newResponseInput(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.currentInput = m.urlInput
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+z", "ctrl+d", "q":
			return m, tea.Quit
		case "tab":
			if m.currentMode == 0 {
				m.currentMode++
			} else {
				m.currentMode--
			}
			return m, nil
		case "up":
			if m.currentInputStyle == URL {
				if m.currentVerb != 0 {
					m.currentVerb--
				} else {
					m.currentVerb = len(verbs) - 1
				}
			}
		case "down":
			if m.currentInputStyle == URL {
				if m.currentVerb != len(verbs)-1 {
					m.currentVerb++
				} else {
					m.currentVerb = 0
				}
			}
		case "ctrl+u":
			m.currentInputStyle = URL
			m.currentInput = m.urlInput
			cmd := m.urlInput.Focus()
			return m, cmd
		case "ctrl+n":
			m.currentInput.Blur()
			m.currentInputStyle = NONE
			return m, nil
		case "ctrl+b":
			m.currentInputStyle = BODY
			m.currentInput = m.bodyInput
			cmd := m.bodyInput.Focus()
			return m, cmd
		case "ctrl+r":
			m.currentInputStyle = RESPONSE
			m.currentInput = m.response
			cmd := m.response.Focus()
			return m, cmd
		case "enter":
			if m.currentInputStyle != BODY {
				url := m.urlInput.Value()
				body := m.bodyInput.Value()
				reader := strings.NewReader(body)
				log.Println(url, body, reader)

				req, err := http.NewRequest(string(verbs[m.currentVerb]), url, reader)
				if err != nil {
					m.response.SetValue(fmt.Sprintf("%s", err.Error()))
					return m, nil
				}

				client := &http.Client{Timeout: time.Second * 60}
				res, err := client.Do(req)
				if err != nil {
					m.response.SetValue(fmt.Sprintf("%d: %s", 404, err.Error()))
					return m, nil
				}
				defer res.Body.Close()

				buf, err := io.ReadAll(res.Body)
				if err != nil {
					m.response.SetValue(fmt.Sprintf("%s", err.Error()))
					return m, nil
				}
				m.response.SetValue(fmt.Sprintf("Code %d: %s", res.StatusCode, string(buf)))
				return m, nil
			}
		}
	}

	m.currentInput, cmd = m.currentInput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	title := m.styles.title.Render("Tuippscotch")
	verbView := ""
	verbStyles := map[verb]lipgloss.Style{
		GET:    m.styles.get,
		POST:   m.styles.post,
		PUT:    m.styles.put,
		PATCH:  m.styles.patch,
		DELETE: m.styles.delete,
	}
	for i, verb := range verbs {
		if i == m.currentVerb {
			verbView += verbStyles[verb].Render(string(verb))
		}
	}

	urlInput := ""
	if m.currentInputStyle == URL {
		urlInput += m.styles.focusedInput.Height(1).Render(m.urlInput.View())
	} else {
		urlInput += m.styles.unFocusedInput.Render(m.urlInput.View())
	}

	bodyInput := ""
	if m.currentInputStyle == BODY {
		bodyInput += m.styles.focusedInput.Height(20).Render(m.bodyInput.View())
	} else {
		bodyInput += m.styles.unFocusedInput.Render(m.bodyInput.View())
	}

	response := ""
	if m.currentInputStyle == RESPONSE {
		response += m.styles.focusedInput.Height(20).Render(m.response.View())
	} else {
		response += m.styles.unFocusedInput.Render(m.response.View())
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				verbView,
				urlInput,
			),
			bodyInput,
			response,
		),
	)
}
