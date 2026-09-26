package main

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4")).Underline(true)
	errorStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF5555"))
	helpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262"))
	tableHead   = lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#282A36"))
)

func (m Model) renderDashboard() string {
	builder := strings.Builder{}
	builder.WriteString(headerStyle.Render("=== GLEDGER DASHBOARD ==="))
	builder.WriteString("\n\n")

	tHeader := fmt.Sprintf("%-12s  %-25s  %-22s  %-22s  %10s", "Date", "Description", "Debit Account (+)", "Credit Account (-)", "Amount")
	builder.WriteString(tableHead.Render(tHeader))
	builder.WriteString("\n")

	for _, tx := range m.Ledger.Transactions {
		amtFormatted := fmt.Sprintf("$%.2f", float64(tx.AmountCents)/100.0)
		row := fmt.Sprintf("%-12s  %-25s  %-22s  %-22s  %10s",
			tx.Date,
			tx.Description,
			tx.DebitAcc,
			tx.CreditAcc,
			amtFormatted,
		)
		builder.WriteString(row)
		builder.WriteString("\n")
	}

	builder.WriteString("\n")
	builder.WriteString(helpStyle.Render("[n] Add New Transaction  |  [q] Quit System"))
	builder.WriteString("\n")
	return builder.String()
}

func (m Model) renderEntryForm() string {
	builder := strings.Builder{}
	builder.WriteString(headerStyle.Render("=== POST NEW JOURNAL ENTRY TRANSACTION ==="))
	builder.WriteString("\n\n")

	labels := []string{"Date:       ", "Description:", "Debit Acc:  ", "Credit Acc: ", "Amt Cents:  "}
	for i, input := range m.Inputs {
		prefix := "  "
		if i == m.FocusedInput {
			prefix = "> "
		}
		fmt.Fprintf(&builder, "%s%s %s\n", prefix, labels[i], input.View())
	}

	if m.FormErr != "" {
		builder.WriteString("\n")
		builder.WriteString(errorStyle.Render(m.FormErr))
		builder.WriteString("\n")
	}

	builder.WriteString("\n")
	builder.WriteString(helpStyle.Render("[Tab/Arrows] Navigate Fields  |  [Enter] Submit at bottom  |  [Esc] Cancel"))
	return builder.String()
}
