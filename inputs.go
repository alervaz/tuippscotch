package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Blur()
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
	Focus() tea.Cmd
	Value() string
	SetValue(string)
}

type UrlInput struct {
	tt textinput.Model
}

func (u *UrlInput) Blur() {
	u.tt.Blur()
}

func (u *UrlInput) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	u.tt, cmd = u.tt.Update(msg)
	return u, cmd
}

func (u *UrlInput) View() string {
	return u.tt.View()
}

func (u *UrlInput) Focus() tea.Cmd {
	return u.tt.Focus()
}

func (u *UrlInput) Value() string {
	return u.tt.Value()
}

func (u *UrlInput) SetValue(value string) {
	u.tt.SetValue(value)
}

func newUrlInput() *UrlInput {
	urlInput := textinput.New()
	urlInput.Placeholder = "URL..."
	return &UrlInput{
		tt: urlInput,
	}
}

type BodyInput struct {
	ta textarea.Model
}

func (b *BodyInput) Blur() {
	b.ta.Blur()
}

func (b *BodyInput) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	b.ta, cmd = b.ta.Update(msg)
	return b, cmd
}

func (b *BodyInput) View() string {
	return b.ta.View()
}

func (b *BodyInput) Focus() tea.Cmd {
	return b.ta.Focus()
}

func (b *BodyInput) Value() string {
	return b.ta.Value()
}

func (b *BodyInput) SetValue(value string) {
	b.ta.SetValue(value)
}

func newBodyInput() *BodyInput {
	bodyInput := textarea.New()
	bodyInput.Placeholder = "Body..."

	return &BodyInput{
		ta: bodyInput,
	}
}

type ResponseInput struct {
	ra textarea.Model
}

func (r *ResponseInput) Blur() {
	r.ra.Blur()
}

func (r *ResponseInput) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	r.ra, cmd = r.ra.Update(msg)
	return r, cmd
}

func (r *ResponseInput) View() string {
	return r.ra.View()
}

func (r *ResponseInput) Focus() tea.Cmd {
	return r.ra.Focus()
}

func (r *ResponseInput) Value() string {
	return r.ra.Value()
}

func (r *ResponseInput) SetValue(value string) {
	r.ra.SetValue(value)
}

func newResponseInput() *ResponseInput {
	responseInput := textarea.New()
	responseInput.Placeholder = "Response..."

	return &ResponseInput{
		ra: responseInput,
	}
}

type HeadersInput struct {
	ha textarea.Model
}

func (h *HeadersInput) Blur() {
	h.ha.Blur()
}

func (h *HeadersInput) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	h.ha, cmd = h.ha.Update(msg)
	return h, cmd
}

func (h *HeadersInput) View() string {
	return h.ha.View()
}

func (h *HeadersInput) Focus() tea.Cmd {
	return h.ha.Focus()
}

func (h *HeadersInput) Value() string {
	return h.ha.Value()
}

func (h *HeadersInput) SetValue(value string) {
	h.ha.SetValue(value)
}

func newHeadersInput() *HeadersInput {
	headersInput := textarea.New()
	headersInput.Placeholder = "Headers..."

	return &HeadersInput{
		ha: headersInput,
	}
}
