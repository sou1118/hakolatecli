package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type runtimeModel struct {
	choices  []string
	cursor   int
	selected string
}

func (m runtimeModel) Init() tea.Cmd {
	return nil
}

func (m runtimeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			m.selected = m.choices[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m runtimeModel) View() string {
	s := "Choose a container runtime:\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s
}

func installRuntime(runtime string) {
	fmt.Printf("You selected %s, installing...\n", runtime)
	switch runtime {
	case "firecracker":
		installFirecracker()
	case "runc":
		installRunc()
	case "youki":
		installYouki()
	default:
		fmt.Println("No valid runtime selected.")
		os.Exit(1)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "hakolatecli",
		Short: "hakolatecli is a CLI tool for setting up container runtimes and nerdctl",
	}

	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install container runtime and nerdctl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hakolate setup: First, we will install nerdctl...")
			installNerdctl()

			m := runtimeModel{
				choices: []string{"firecracker", "runc", "youki"},
			}

			p := tea.NewProgram(m)
			model, err := p.Run()
			if err != nil {
				fmt.Printf("Could not start the selection program: %s\n", err)
				os.Exit(1)
			}

			selectedRuntime := model.(runtimeModel).selected
			installRuntime(selectedRuntime)
		},
	}

	rootCmd.AddCommand(installCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
