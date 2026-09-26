package main

import (
	"fmt"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type SessionState int

const (
	ViewDashboard SessionState = iota
	ViewEntryForm
)

const (
	InputDate = iota
	InputDesc
	InputDebit
	InputCredit
	InputAmount
)

type Model struct {
	State        SessionState
	Ledger       Ledger
	Inputs       []textinput.Model
	FocusedInput int
	FormErr      string
}

func NewModel() Model {
	m := Model{
		State:        ViewDashboard,
		FocusedInput: 0,
		Ledger: Ledger{
			Transactions: []Transaction{
				{Date: "2026-09-20", Description: "Opening Balance Checking", DebitAcc: "Assets:Checking", CreditAcc: "Equity:Opening", AmountCents: 500000},
				{Date: "2026-09-24", Description: "Weekly Groceries Buy", DebitAcc: "Expenses:Food", CreditAcc: "Assets:Checking", AmountCents: 12450},
			},
		},
	}

	m.Inputs = make([]textinput.Model, 5)

	m.Inputs[InputDate] = textinput.New()
	m.Inputs[InputDate].Placeholder = "YYYY-MM-DD (e.g. 2026-09-25)"
	m.Inputs[InputDate].Focus()

	m.Inputs[InputDesc] = textinput.New()
	m.Inputs[InputDesc].Placeholder = "Description (e.g. Coffee)"

	m.Inputs[InputDebit] = textinput.New()
	m.Inputs[InputDebit].Placeholder = "Debit Account (e.g. Expenses:Food)"

	m.Inputs[InputCredit] = textinput.New()
	m.Inputs[InputCredit].Placeholder = "Credit Account (e.g. Assets:Checking)"

	m.Inputs[InputAmount] = textinput.New()
	m.Inputs[InputAmount].Placeholder = "Amount in cents (e.g. 450)"

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.State == ViewDashboard {
				return m, tea.Quit
			}
		case "n":
			if m.State == ViewDashboard {
				m.State = ViewEntryForm
				m.FormErr = ""
				m.FocusedInput = 0
				m.Inputs[InputDate].Focus()
				for i := 1; i < len(m.Inputs); i++ {
					m.Inputs[i].Blur()
				}
				return m, nil
			}
		case "esc":
			if m.State == ViewEntryForm {
				m.State = ViewDashboard
				return m, nil
			}
		case "tab", "shift+tab", "up", "down", "enter":
			if m.State == ViewEntryForm {
				if msg.String() == "enter" && m.FocusedInput == len(m.Inputs)-1 {
					return m.submitForm()
				}

				if msg.String() == "tab" || msg.String() == "down" || msg.String() == "enter" {
					m.Inputs[m.FocusedInput].Blur()
					m.FocusedInput = (m.FocusedInput + 1) % len(m.Inputs)
					m.Inputs[m.FocusedInput].Focus()
				} else {
					m.Inputs[m.FocusedInput].Blur()
					m.FocusedInput--
					if m.FocusedInput < 0 {
						m.FocusedInput = len(m.Inputs) - 1
					}
					m.Inputs[m.FocusedInput].Focus()
				}
				return m, nil
			}
		}
	}

	if m.State == ViewEntryForm {
		var cmd tea.Cmd
		m.Inputs[m.FocusedInput], cmd = m.Inputs[m.FocusedInput].Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) submitForm() (tea.Model, tea.Cmd) {
	amount, err := strconv.ParseInt(m.Inputs[InputAmount].Value(), 10, 64)
	if err != nil {
		m.FormErr = "Error: Amount must be a valid integer raw representation of cents."
		return m, nil
	}

	newTx, err := m.Ledger.ValidateEntry(
		m.Inputs[InputDate].Value(),
		m.Inputs[InputDesc].Value(),
		m.Inputs[InputDebit].Value(),
		m.Inputs[InputCredit].Value(),
		amount,
	)

	if err != nil {
		m.FormErr = fmt.Sprintf("Error: %v", err)
		return m, nil
	}

	m.Ledger.Transactions = append(m.Ledger.Transactions, *newTx)

	for i := range m.Inputs {
		m.Inputs[i].SetValue("")
	}
	m.State = ViewDashboard
	return m, nil
}

// View switches between dashboard and data entry views cleanly
func (m Model) View() tea.View {
	var v tea.View
	v.AltScreen = true

	var docStyle = lipgloss.NewStyle().Padding(1, 2)
	var content string

	if m.State == ViewDashboard {
		content = m.renderDashboard()
	} else {
		content = m.renderEntryForm()
	}

	v.Content = docStyle.Render(content)
	return v
}
