package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	// Initialize UI program state context cleanly from layout definition module
	p := tea.NewProgram(NewModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Fatal execution crash: %v\n", err)
		os.Exit(1)
	}
}
