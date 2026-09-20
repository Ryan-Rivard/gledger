package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/textinput"
)

type AccountType string

const (
	Asset     AccountType = "Asset"
	Liability AccountType = "Liability"
	Equity    AccountType = "Equity"
)

type Account struct {
	Id   string
	Name string
	Type AccountType
}

type Posting struct {
	AccountId string
	Amount    float64 // Positive for Debit, Negative for Credit? or explict db/cr fields
}

type Transaction struct {
	Date        time.Time
	Description string
	Postings    []Posting // must sum to 0 to balance
}

type Ledger struct {
	Accounts     map[string]*Account
	Transactions []Transaction
}

type sessionState int

const (
	listView sessionState = iota
	formView
)

type model struct {
	state        sessionState
	ledger       Ledger
	inputs       []textinput.Model // form inputs for Description, Account1, Account2, Amount
	focusedInput int
	err          error
}

func initialModel() model {
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
		s += "formview"
	case listView:
		s += "listview"
	default:
		panic(fmt.Sprintf("unexpected main.sessionState: %#v", m.state))
	}
	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
