// Main bubbletea model and view switcher
package ui

import (
	"fmt"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"github.com/Ryan-Rivard/gledger/core"
	"github.com/charmbracelet/bubbles/textinput"
)

type sessionState int

const (
	listView sessionState = iota
	formView
)

// updated model is returned on each update cycle and replaces the original model
// but the ledger is a pointer, so there is only a copy of the pointer
// ledger updates happen in place
type model struct {
	state        sessionState
	ledger       *core.Ledger // mutate the ledger in place
	inputs       []textinput.Model // form inputs for Description, Account1, Account2, Amount
	focusedInput int
	err          error
}

func InitialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case formView:
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "1":
				m.state = listView
				m.ledger.Transactions = append(m.ledger.Transactions, core.Transaction{})
			}
		}

	case listView:
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "1":
				m.state = formView
				m.ledger.Transactions = append(m.ledger.Transactions, core.Transaction{})
			}
		}
	default:
		panic(fmt.Sprintf("unexpected main.sessionState: %#v", m.state))
	}

	return m, nil
}

func (m model) View() tea.View {
	s := "1 to switch views, q to quit\r\n\r\n"
	switch m.state {
	case formView:
		s += "formview "
		s += strconv.Itoa(len(m.ledger.Transactions))
	case listView:
		s += "listview "
		s += strconv.Itoa(len(m.ledger.Transactions))
	default:
		panic(fmt.Sprintf("unexpected main.sessionState: %#v", m.state))
	}
	return tea.NewView(s)
}
